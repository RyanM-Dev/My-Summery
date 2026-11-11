# 🔎 Raw Evidence — Develop Server Audit 2026-11-11

The develop server: Ubuntu 26.04.1 LTS, 2 vCPU, 3.8 GiB RAM, no swap, 58 GB disk.

All commands were **read-only**, run as root over SSH. Secrets are redacted.

> Note: these numbers are a snapshot from 2026-11-11, around 10:00–12:10 UTC. Re-run the commands to see the current state.

---

## 👤 Accounts and SSH

| Check | Command | Result |
|---|---|---|
| Login shells | `awk -F: '$7 !~ /nologin\|false/' /etc/passwd` | `root`, `ubuntu` (1000), `ryan` (1001) |
| sudo group | `getent group sudo` | `ubuntu,ryan` |
| NOPASSWD | `grep -r NOPASSWD /etc/sudoers*` | `root`, **`ubuntu ALL=(ALL) NOPASSWD:ALL`** |
| Password status | `passwd -S <user>` | root **P** (usable), ubuntu **L** (locked), ryan **P** |
| Root keys | `/root/.ssh/authorized_keys` (600) | 2 RSA keys: yours, and **`support@provider`** |
| ryan keys | `/home/ryan/.ssh/authorized_keys` | mode **664** (group-writable). A key login as ryan was tested and **works** (`Accepted publickey for ryan`): OpenSSH tolerates group-write for a private user group |
| sshd effective | `sshd -T` | `permitrootlogin yes`, **`passwordauthentication yes`**, `x11forwarding yes`, `maxauthtries 6`, `logingracetime 120` |
| sshd drop-ins | `/etc/ssh/sshd_config.d/` | `50-cloud-init.conf: PasswordAuthentication yes` is read **before** `60-cloudimg-settings.conf: … no`; first value wins |
| Failed logins 24 h | `journalctl -u ssh --since "24 hours ago" \| grep -cE "Failed password\|Invalid user"` | **22,941** |
| Attackers 24 h | distinct source IPs | **135**; top IP 9,175 attempts |
| Root password attempts 24 h | `grep -c "Failed password for root"` | **8,780** |
| Usernames tried | `Invalid user` | `user` 3,496, `admin` 245, `debian` 125, `test` 59, `deploy` 29 |
| Successful logins 7 d | `grep Accepted` | **36 × `publickey root`**; `ryan` never logs in |
| Brute-force protection | `systemctl is-active fail2ban crowdsec` | inactive / not installed |
| Journal volume | lines by unit, 24 h | **ssh.service 75,857** (docker 8,528) |
| AppArmor | `aa-status` | loaded, 105 profiles enforcing |

## 🌐 Network exposure

| Port | Service | Notes |
|---|---|---|
| 22 | sshd | password auth enabled (see above) |
| 80 / 443 | Nginx Proxy Manager | `http://example.net` → 200 with **no redirect**; no `Strict-Transport-Security` header seen on `git.`/`npm.` |
| **81** | NPM admin UI | `curl http://SERVER_IP:81` from the internet → **200** |
| — | `npm.example.com` | NPM admin **also** published over HTTPS → 200 |
| 2222 | Gitea SSH | expected |
| **1884** | old MQTT broker, **plaintext** | internet scanners connect (the `LZR` scanner) |
| 8884 | old MQTT broker, TLS | |
| **8099** | `crypto-interop-spike-viewer` | from the internet → **200**, no auth |
| `shop.example.net` | NPM → `172.20.0.1` (test backend UI) | **HTTP only**, no access list |
| Firewall | `ufw status` | **inactive**; Docker-published ports bypass ufw anyway |
| `iptables -S DOCKER-USER` | | only `-N DOCKER-USER`: the chain exists and is **empty** |

## 🐳 Docker

- **Daemon:** Docker 29.8, Compose v5.5.1. **No `/etc/docker/daemon.json`**, so: `live-restore=false`, default json-file logging with **no size limit**, no userns-remap. Security options: apparmor, seccomp builtin, cgroupns.

- **docker.sock:** `srw-rw---- root:docker`. The `docker` group is empty. **Mounted into `gitea-runner-runner-1`.**

- **Runner config:** `/opt/gitea-runner/config.yaml` (`container.docker_host` and `privileged` are commented out, so the defaults apply):
    - job `options: --volume /opt/robot-market-front:/opt/robot-market-front`
    - `valid_volumes: [/opt/test-mqtt-5/backend, /opt/robot-market-front]`
    - `network: gitea_proxy`

| Container | User | cap_drop | RO rootfs | Mem limit | Restart | Health | Log limit |
|---|---|---|---|---|---|---|---|
| mqtt5-test-backend ×2 | 10001 | ALL | ✅ | — | unless-stopped | healthy | 10m×3 |
| vending-machine-game | **root** | — | ❌ | — | unless-stopped | none | **none** |
| crypto-interop-spike ×3 | **root** | — | ❌ | — | **no** | broker healthy | **none** |
| robot-market-mqtt5-test-broker | 1883 | ALL | ✅ | 256 MiB | unless-stopped | **unhealthy** | 5m×3 |
| gitea-runner | root | ALL | ❌ | — | always | none | **none** |
| gitea | 1000 | ALL | ❌ | — | always | none | **none** |
| gitea-db, npm, npm-db | root | ALL (+ some adds) | ❌ | — | always | db healthy | **none** |

- **Unhealthy broker:**
    - The health log shows `OCI runtime exec failed: exec failed: unable to start container process: procReady not received`, every 30 s.
    - A manual `docker exec <c> true` fails the same way.
    - The cgroup shows `pids.current=61`, `pids.max=64`, `pids.events: max 3878909`, yet `cgroup.procs` has **1** task: mosquitto, 1 thread, 10 fds, 4.3 MB RSS.
    - The broker still listens and answers TLS on :8884. Its last log line is `2026-11-10 18:51:13 UTC`.

- **Disk:** images 3.3 GB; **build cache 6.2 GB (6.0 GB reclaimable)**.

- **Biggest container logs:**
    - `crypto-interop-spike-broker-1`: **200 MB** (`log_type all`, plus subscribe/unsubscribe, no rotation)
    - `gitea`: 84 MB
    - `gitea-db`: 9.7 MB

## 🩹 Patching

- **Uptime 72 days.** Running kernel **7.0.0-30**; installed **7.0.0-43**.

- `/var/run/reboot-required.pkgs` lists kernels -31, -34, -38, -42, -43 and **libc6**.

- `unattended-upgrades` is on (security and updates origins). Last run 2026-11-11 06:09, "All upgrades installed". `Automatic-Reboot` is commented out (false). 55 packages are upgradable.

- Time: UTC, NTP synchronized.

## ⚡ Performance and capacity

- **Memory:** 3,910 MB total; used 1,909; available ~2,000; **no swap**.
    - `Committed_AS` **6.08 GB** vs `CommitLimit` 2.0 GB (heuristic overcommit, `vm.overcommit_memory=0`). Processes have reserved about 1.5× physical RAM.
    - 0 OOM kills in 30 days.

- **Pressure (PSI):**
    - CPU `some avg300=2.8%`.
    - **IO `full avg300=6.0%`**: for an idle box, something is waiting on disk 6% of the time.
    - Memory `full avg300≈1%`.
    - Load average 0.39. vmstat shows CPU ~95% idle and 0 wa.

- **Top memory:**
    - **VS Code Server `node` as root: 491 MB** (`/root/.vscode-server`, **4.5 GB on disk**).
    - gitea 185–230 MB, npm (node) 130 MB, dockerd 126 MB, mariadb 94–103 MB.

- **Disk:** `/` 22/58 GB (38%), inodes 6%.
    - **Journal 3.6 GB**, no `SystemMaxUse` set. `/var/log` 3.9 GB.
    - Most journal lines come from sshd, because of the brute force.

- **sysctl:** `ip_forward=1` (Docker), `tcp_syncookies=1`, `rp_filter=2`, `swappiness=60`, `somaxconn=4096`, `dmesg_restrict=1`.

## 💾 Backups

- `/backups/gitea-bak/`: one **manual** Gitea dump plus a DB dump from **2026-09-10** (87 MB), on the **same disk**. Nothing newer.

- No cron or systemd timer runs `pg_dump`, `gitea dump`, restic, borg or rsync.

- Not backed up at all: the NPM MariaDB (proxy hosts and certificates), `/opt/*` compose files and `.env`, and the MQTT data.

## 🔐 Secret files

- `.env` files are mode **644 (world-readable)**: `/opt/gitea/.env`, `/opt/gitea-runner/.env` (holds the runner registration token), `/opt/nginx-proxy-manager/.env`, `/opt/robot-market-front/.env(.previous)`, `/opt/test-mqtt-5/{deploy,backend}/.env(.previous)`, plus a backup copy.

- Good: the MQTT secrets directories are `750 root:<service gid>` and the client JSON files are `640`.

## ✅ Things that are already right

- AppArmor is enforcing, and Docker's seccomp and AppArmor profiles are active.
- Unattended security upgrades are installed and running daily.
- NTP is synchronized, and the clock is in UTC.
- Gitea: `DISABLE_REGISTRATION=true`.
- Most stacks drop all capabilities. The MQTT stack runs non-root, with a read-only rootfs, pinned digests and limits.
- Databases are not published on the host; they sit on internal Docker networks.
- Root login is already key-only in practice: all 36 accepted logins used public keys.

## 🔁 Verification pass (review, 2026-11-11 ~12:20 UTC)

- `ssh -o BatchMode=yes ryan@host true` → success. The claim "664 makes sshd refuse ryan's key" was **wrong** and has been corrected.
- `/etc/sudoers.d/90-cloud-init-users` exists and holds the `ubuntu` NOPASSWD line.
- NPM DB env var **names**: `MYSQL_ROOT_PASSWORD`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DATABASE`. Values were not read.
- Gitea image `gitea/gitea:1.27.3-rootless` runs as `git` (uid 1000), config at `/etc/gitea/app.ini`.
- `http://shop.example.net/` `<title>` = `MQTT test viewer`.
