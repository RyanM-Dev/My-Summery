# Docker — Specific Prompt

> **Charts and diagrams:** Use fenced `mermaid` blocks for charts, flows, architecture, hierarchies, and relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` as appropriate. Use clear labels and Obsidian-compatible syntax. Keep runnable code, commands, literal output, payloads, and calculations in their original code formats.


> **Purpose:** Topic layer for Docker / container study guides. Append AFTER [[00 - General Study Guide Prompt]].

---

## When to use

- Docker images, containers, Dockerfile, layers
- Docker Compose multi-service setups
- Container networking, volumes, health checks
- Go testing with dockertest / testcontainers
- Deployment prep for Kubernetes

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — Docker ({{CURRENT_YEAR}})

You are a senior platform engineer specializing in containers. Apply these rules ON TOP of the general prompt.

---

### Docker content requirements

#### Code blocks (mandatory)
- Include **at least 2 Dockerfile examples** (multi-stage when building apps)
- Include **at least 1 docker-compose.yaml** block for multi-service chapters
- Include **at least 3 bash blocks** (`docker build`, `docker run`, `docker compose up`, `docker logs`)
- Pin base image tags — warn against `latest` in production

#### Diagrams (mandatory)
- Container vs image layer stack (Mermaid)
- Compose service network diagram (mermaid flowchart)
- Volume mount vs bind mount comparison table

#### {{CURRENT_YEAR}} Docker standards (Part 7–8)
- Multi-stage builds for Go (`golang:alpine` builder → `scratch` or `distroless` runtime)
- Non-root `USER` in production images
- `HEALTHCHECK` or Compose `healthcheck` for dependency ordering
- `.dockerignore` to reduce build context
- Compose v2 (`docker compose`, not `docker-compose` hyphen)
- Prefer BuildKit (`DOCKER_BUILDKIT=1`)
- Security: scan images, minimal base images, no secrets in layers

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks`. Each task:
- **Goal** | **Files** | **Steps** | **Done when**
- Verifiable with `docker ps`, `curl`, or `docker compose logs`

---

### Docker interview focus (where relevant)

- Image layers and cache invalidation
- CMD vs ENTRYPOINT
- Bridge vs host networking
- Named volumes vs bind mounts
- Compose service dependencies and health checks
- Container isolation limits (CPU, memory)

---

Generate the complete study guide now.
```