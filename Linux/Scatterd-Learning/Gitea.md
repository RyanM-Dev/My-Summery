# Gitea + Gitea Actions Runner — Production Setup Guide

This document covers the full setup, configuration, and troubleshooting reference for our self-hosted Gitea instance with Gitea Actions (CI/CD) enabled, based on the official Gitea documentation and lessons learned during initial setup.

---

## 1. Architecture Overview

```
                        ┌─────────────────────────┐
                        │  Nginx Proxy Manager     │
                        │  (public 80/443)         │
                        └───────────┬─────────────┘
                                    │  gitea_proxy network
                    ┌───────────────┼────────────────┐
                    │                                │
          ┌─────────▼─────────┐          ┌───────────▼──────────┐
          │   gitea container  │          │   runner container    │
          │  (git.my-rm.com)   │◄────────►│   (runner-01)          │
          └─────────┬──────────┘  DNS via │                        │
                    │             shared  │  spawns job containers │
          gitea_internal network  network │  on gitea_proxy too    │
                    │                     └────────────────────────┘
          ┌─────────▼──────────┐
          │  db (postgres)      │
          └─────────────────────┘
```

**Two Docker networks matter here:**

|Network|Purpose|Type|
|---|---|---|
|`gitea_internal`|Private network between Gitea app and its Postgres DB. Not exposed to the runner or proxy.|internal, defined per-stack|
|`gitea_proxy`|Shared network so NPM can reach Gitea's HTTP port, and so the runner (and its job containers) can reach Gitea by hostname.|`external: true`, created once, shared across stacks|

**Golden rule:** any container that needs to reach Gitea by name (`http://gitea:3000`) — including ephemeral job containers spawned by the runner — must be on `gitea_proxy`.

---

## 2. Gitea Server — `docker-compose.yml`

```yaml
services:
  db:
    image: postgres:16-alpine
    restart: always
    environment:
      POSTGRES_USER: gitea
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: gitea
    volumes:
      - ./db:/var/lib/postgresql/data
    networks:
      - gitea_internal
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U gitea"]
      interval: 10s
      timeout: 5s
      retries: 5
    cap_drop:
      - ALL
    cap_add:
      - CHOWN
      - DAC_OVERRIDE
      - SETGID
      - SETUID
      - NET_BIND_SERVICE

  gitea:
    image: gitea/gitea:1.27.3-rootless
    restart: always
    depends_on:
      db:
        condition: service_healthy
    environment:
      USER_UID: 1000
      USER_GID: 1000
      GITEA_CUSTOM: /etc/gitea

      GITEA__database__DB_TYPE: postgres
      GITEA__database__HOST: db:5432
      GITEA__database__NAME: gitea
      GITEA__database__USER: gitea
      GITEA__database__PASSWD: ${DB_PASSWORD}

      GITEA__server__DOMAIN: git.my-rm.com
      GITEA__server__ROOT_URL: https://git.my-rm.com
      GITEA__server__SSH_DOMAIN: git.my-rm.com
      GITEA__server__SSH_PORT: 2222
      GITEA__server__LFS_START_SERVER: "true"
      GITEA__server__HTTP_PORT: 3000

      GITEA__service__DISABLE_REGISTRATION: "true"
      GITEA__actions__ENABLED: "true"
      GITEA__actions__DEFAULT_ACTIONS_URL: https://gitea.com

    volumes:
      - ./data:/var/lib/gitea
      - ./config:/etc/gitea
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
    ports:
      - "2222:2222"   # SSH only – HTTP goes through NPM
    networks:
      - gitea_internal
      - gitea_proxy
    cap_drop:
      - ALL
    cap_add:
      - CHOWN
      - DAC_OVERRIDE
      - SETGID
      - SETUID
      - NET_BIND_SERVICE

networks:
  gitea_internal:
    # Private network for Gitea + DB
  gitea_proxy:
    external: true
```

### Notes on this file

- `GITEA__actions__ENABLED: "true"` is what turns on Actions instance-wide — required before any runner can register jobs.
- HTTP (port 3000) is **not** published to the host — only reachable internally via `gitea_proxy`, and exposed publicly through Nginx Proxy Manager, which proxies `https://git.my-rm.com` → `gitea:3000` over that same network.
- SSH (port 2222) is published directly since NPM doesn't proxy raw TCP/SSH traffic the same way as HTTP.
- Least-privilege `cap_drop: ALL` + explicit `cap_add` is good practice — the container only gets the specific Linux capabilities it actually needs.

### Creating the shared network (one-time)

```bash
docker network create gitea_proxy
```

Do this once, before bringing up either stack. Both `docker-compose.yml` files reference it as `external: true` so Compose won't try to recreate it.

---

## 3. Gitea Actions Runner — `docker-compose.yml`

```yaml
services:
  runner:
    image: gitea/runner:3.3.0
    restart: always
    environment:
      GITEA_INSTANCE_URL: http://gitea:3000
      GITEA_RUNNER_REGISTRATION_TOKEN: ${RUNNER_TOKEN}
      GITEA_RUNNER_NAME: runner-01
      CONFIG_FILE: /config.yaml
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./data:/data
      - ./config.yaml:/config.yaml
    networks:
      - gitea_proxy
    cap_drop:
      - ALL
    cap_add:
      - NET_ADMIN
      - NET_RAW
      - CHOWN
      - DAC_OVERRIDE
      - SETGID
      - SETUID

networks:
  gitea_proxy:
    external: true
```

### Why each piece matters

- `/var/run/docker.sock` mounted → lets the runner spawn job containers using the **host's** Docker daemon (simplest mode; the alternative is Docker-in-Docker with `--privileged`, which is more isolated but more complex — not used here).
- `./data:/data` → persists the runner's registration file (`.runner`) so it survives container recreation. **Without this, the runner would re-register as a brand new runner every time the container is recreated.**
- `./config.yaml:/config.yaml` + `CONFIG_FILE: /config.yaml` → tells the runner to load our custom config instead of pure defaults. This is required to override `container.network` (see below) — without it, job containers get an isolated per-job network and **cannot reach Gitea by hostname.**
- `gitea_proxy` network on the runner's own container is necessary so the runner daemon itself can reach `http://gitea:3000` to register and poll for jobs — but this does **not** automatically apply to the ephemeral job containers it spawns. That's a separate setting (`container.network` in config.yaml).

---

## 4. Generating and Configuring `config.yaml` (the part that trips everyone up)

### The one command that actually works

The runner's Docker image entrypoint is `run.sh`, which **always** launches the full daemon/registration flow regardless of what command you pass it — unless you explicitly clear the entrypoint. This is the single most common point of confusion.

**Correct command** (per official Gitea documentation):

```bash
docker run --rm --entrypoint="" gitea/runner:3.3.0 gitea-runner generate-config > config.yaml
```

Key details:

- `--entrypoint=""` clears the baked-in entrypoint so Docker runs your command directly instead of `run.sh`.
- The binary inside this image is called **`gitea-runner`** (not `act_runner` — that name belongs to the older upstream project some docs and blog posts still reference).
- Do **not** omit `--entrypoint=""` — without it, the container ignores your command entirely, tries to register/run the daemon, and fails with a misleading "missing token" error loop, or (worse) captures log noise into your redirected file instead of real YAML.

### Verify before trusting it

```bash
cat config.yaml
```

You should see a large, fully-commented YAML file with sections like `log:`, `runner:`, `cache:`, `container:`, `host:`. If instead you see log lines like `level=info msg=...`, the entrypoint wasn't cleared and the file is garbage — delete it and redo the command above.

### Required edits to the generated file

**1. Registration file path** — must point at the persisted `/data` volume:

```yaml
runner:
  file: /data/.runner
```

Default is `.runner` (relative to the container's working directory), which does **not** map to your persistent volume. If left as default, every container recreation loses the registration and the runner errors out trying to re-register without a valid token.

**2. Job container network** — must match the shared network:

```yaml
container:
  network: "gitea_proxy"
```

Default is `""` (empty), which makes the runner create a **new isolated Docker network per job**. Job containers on that isolated network cannot resolve `gitea` by hostname, which causes `Could not resolve host: gitea` failures during the checkout step of every workflow. Setting this explicitly to your shared network fixes it permanently.

### Validate YAML syntax before deploying

```bash
python3 -c "import yaml; yaml.safe_load(open('config.yaml'))" && echo "YAML OK"
```

### Apply and verify

```bash
docker compose up -d --force-recreate runner
docker exec gitea-runner-runner-1 cat /config.yaml | grep -A2 "^runner:"
docker exec gitea-runner-runner-1 cat /config.yaml | grep -A2 "^container:"
```

---

## 5. Registering / Verifying the Runner

- **Token**: generated from Gitea UI at **Site Administration → Actions → Runners → Create new Runner**, or via CLI: `gitea --config /etc/gitea/app.ini actions generate-runner-token`. Put this value in `RUNNER_TOKEN` inside your `.env` file next to the runner's compose file.
- **First start**: the runner registers automatically using `GITEA_RUNNER_REGISTRATION_TOKEN` from the environment, and writes `/data/.runner`. This file is the runner's persistent identity — **do not delete it** unless you intend to re-register as a new runner.
- **Verify registration**: Gitea UI → **Site Administration → Actions → Runners** (or repo-level **Settings → Actions → Runners**). Should show `runner-01` as Idle/Online with labels: `ubuntu-latest`, `ubuntu-24.04`, `ubuntu-22.04`.

---

## 6. Test Workflow

Create `.gitea/workflows/test.yml` in any repository:

```yaml
name: Runner Test

on:
  push:
    branches: [ main ]
  workflow_dispatch:   # allows manual trigger from the Actions tab

jobs:
  hello:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Print basic info
        run: |
          echo "✅ Runner is working!"
          echo "Hostname: $(hostname)"
          echo "Date: $(date)"
          whoami
          pwd
          ls -la

      - name: Test env variables
        run: |
          echo "Branch: ${{ github.ref_name }}"
          echo "Commit SHA: ${{ github.sha }}"
          echo "Actor: ${{ github.actor }}"

      - name: Simulate a build step
        run: |
          mkdir -p build
          echo "Build artifact content" > build/output.txt
          cat build/output.txt
```

Push it, or trigger manually from **repo → Actions → Runner Test → Run workflow**. The **Checkout code** step is the key indicator: if it fetches successfully, both the registration path and the `container.network` setting are correctly configured.

### Debugging a failed run

Add a secret named `ACTIONS_STEP_DEBUG` = `true` at repo or org level (Settings → Actions → Secrets) and re-run — this surfaces the actual underlying error (DNS failure, auth failure, timeout, etc.) instead of a generic exit-code message.

---

## 7. Troubleshooting Reference (issues actually hit during setup)

|Symptom|Root Cause|Fix|
|---|---|---|
|`Could not resolve host: gitea` during checkout|Job containers spawned on an isolated per-job network, not `gitea_proxy`|Set `container.network: "gitea_proxy"` in `config.yaml`|
|`.runner is missing or not a regular file` + `missing token` loop|`runner.file` path doesn't point at the persisted `/data` volume|Set `runner.file: /data/.runner`|
|`yaml: line 3: mapping values are not allowed`|`config.yaml` was corrupted by log output captured during a bad `generate-config` run|Regenerate using the correct `--entrypoint=""` command below|
|`executable file not found in $PATH` when running `generate-config`|Wrong binary name assumed (`act_runner` instead of `gitea-runner`)|Use `gitea-runner generate-config`, not `act_runner`|
|`generate-config` silently launches the daemon instead of printing config|Entrypoint (`run.sh`) not cleared|Always pass `--entrypoint=""`|
|Runner container recreated but re-registers as new runner|`/data` volume not mounted, or `.runner` file lost|Ensure `./data:/data` is mounted and persists across `docker compose down/up`|

---

## 8. Quick Command Reference

```bash
# One-time: create shared network
docker network create gitea_proxy

# Generate a clean config file
docker run --rm --entrypoint="" gitea/runner:3.3.0 gitea-runner generate-config > config.yaml

# Validate YAML
python3 -c "import yaml; yaml.safe_load(open('config.yaml'))" && echo "YAML OK"

# Bring up / recreate the runner after config changes
docker compose up -d --force-recreate runner

# Check logs
docker compose logs -f runner

# Inspect what network a running job container actually joined
docker inspect <job-container-id> | grep -A 5 '"Networks"'

# Confirm config actually loaded inside the container
docker exec gitea-runner-runner-1 cat /config.yaml | grep -A2 "^runner:"
docker exec gitea-runner-runner-1 cat /config.yaml | grep -A2 "^container:"
```

---

## 9. Backup Checklist (for this stack specifically)

Back up these paths regularly (see broader backup strategy doc for the full-server pattern):

- `gitea/./data` → Gitea's repo data, LFS objects, avatars, attachments
- `gitea/./config` → `app.ini` and any custom templates
- `gitea/./db` → Postgres data directory (or better: `pg_dump` output instead of raw files)
- `gitea-runner/./data/.runner` → runner registration identity (losing this just means re-registering, not catastrophic, but saves a step)
- `gitea-runner/config.yaml` → your tuned runner configuration (network + file path fixes)

Example dump command for the DB:

```bash
docker compose exec db pg_dump -U gitea gitea > backup_$(date +%F).sql
```

---

## 10. Key Takeaways for Future Setups

1. Always pass `--entrypoint=""` when running one-off admin commands against images that wrap a supervising script (`run.sh`, `entrypoint.sh`, etc.) — otherwise your command is silently ignored in favor of the image's default startup behavior.
2. A shared Docker network solves 90% of "container can't reach container" issues — but remember that **runner daemons and the job containers they spawn are not the same container**, and both need to be on the right network explicitly.
3. Any stateful identity (registration tokens, `.runner` files) must live on a mounted volume, or it's lost on every `--force-recreate`.
4. When something fails with a vague error, turn on `ACTIONS_STEP_DEBUG` (or the equivalent verbose/debug flag for whatever tool you're debugging) before guessing — it usually surfaces the real cause immediately.