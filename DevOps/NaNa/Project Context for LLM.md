# Nana DevOps Project Context for LLMs

> Maintained as a working handoff for future LLM sessions. This summarizes the Nana DevOps vault, its conventions, important technical concepts, and relevant prior chat decisions. It is a context guide, not a replacement for the linked lesson notes or the original chat transcripts.

## Project purpose and boundaries

`DevOps/NaNa/` is an Obsidian learning and reference collection for practical DevOps topics. The material is organized as concise lessons, learning paths, and operational references. The course currently covers Linux, Git, Docker, Jenkins, and artifact management with Gitea and Sonatype Nexus Repository Manager.

Treat the Markdown notes as teaching material, not proof that a command was executed against a live system. A note may include validated syntax without having run a real deployment. Preserve this distinction in future edits and explanations.

## Start here

- [[0. Nana DevOps Table of Contents]] — learner-facing course navigation.
- [[AI Prompt - NANA DevOps]] — primary Nana course writing instructions.
- [[Docker/AI Prompt - Docker]] — Docker-specific additions.
- [[Artifact Management/0. Artifact Management]] — Gitea and Nexus learning paths.
- [[Git NANA/0. Git Learning Path]] — Git lessons ordered from foundations through shared-history safety.
- [[Linux/0. Linux Learning Path]] — Linux notes ordered from shell basics through server configuration.

When editing a note, inspect the target and its neighboring references first. For lesson writing, follow the Nana prompt and any more-specific local prompt. The user has requested concise, example-led prose, diagrams where they clarify relationships, and readable spacing.

## Current vault map

### Linux

- `Linux/0. Linux Learning Path.md`
- `Linux/1. Linux Pipes, grep, less, and Output Redirection.md`
- `Linux/2. Vim Cheat Sheet.md`
- `Linux/3. Environment Variables.md`
- `Linux/4. Users and Groups.md`
- `Linux/5. Ownership and Permissions.md`
- `Linux/6. Networking Fundamentals.md`
- `Linux/7. Server Configuration.md`

### Git

`Git NANA/0. Git Learning Path.md` links ten progressively ordered lessons: history/commits/HEAD; tracking and ignore rules; branches and team workflow; merge and conflict resolution; remotes and syncing; stash; rebase; release branches, versions, and tags; amend and local reset; reset, revert, and shared history.

### Docker

`Docker/1`–`8` cover containers, architecture, basic commands, real-project workflow, Compose, Dockerfiles, volumes, and best practices. `.OLD-notes` files are archived source material; consult them only when the current reference needs historical context. `Docker/AI Prompt - Docker.md` is the topic prompt.

### Jenkins

- `Ci-CD/Jenkins/1. Basics.md` — Jenkins concepts and setup.
- `Ci-CD/Jenkins/2. Docker in Jenkins.md` — Docker CLI/Engine integration, socket access, DinD alternative, images, credentials, and volumes.
- `Ci-CD/Jenkins/1. Workflow for AI Coding.md` — small companion reference.
- `Ci-CD/Jenkins/2.OLD-notes - Docker in Jenkins.md` — archived rough source.
- `Ci-CD/Jenkins/Untitled.md` — currently empty; do not assume it is a lesson.

### Artifact management

`Artifact Management/0. Artifact Management.md` introduces the two tracks:

- **Gitea:** packages and ownership; Docker registry use; CI image publishing; storage cleanup and troubleshooting.
- **Nexus:** artifacts and repositories; Raw repositories, users/roles, and REST API; blob stores; cleanup policies; pushing Docker images from Jenkins.

## Teaching and editing conventions

- Keep filenames and folders spaced and human-readable; retain technical names and the ` - ` lesson-title separator where already used.
- Preserve canonical note paths when possible. When converting rough notes, the vault often keeps an archive named `.OLD-notes - ...` and places the finished reference at the original path.
- Explain in this order: simple understanding, a real-world example, how it works, then necessary technical detail. Add advanced detail only when it helps.
- Use Mermaid for architecture, relationships, and workflows; use sequence diagrams for exchanges. Keep diagrams Obsidian-compatible.
- Use selective meaningful emojis, short paragraphs, properly indented lists, and blank lines between separate blocks.
- For substantial lessons, use a small understanding-focused Q&A with all answers together in a collapsed `> [!answer]-` callout. Do not use HTML `<details>` for answers.
- Use safe hands-on exercises when appropriate. Finish teaching references with `Quick Reference`, `Things to Remember`, `Pro Tips`, and `Related Topics` in that order.
- Update Obsidian links and navigation when moving or renaming notes. Check the table of contents before changing it.
- Treat the current workspace state as user-owned. At the time this handoff was assembled, `.obsidian/workspace.json`, the Nana table of contents, the Nexus image-publishing note, and Jenkins notes appeared as existing local changes or untracked files. Do not discard or overwrite them blindly.

## Important concepts and decisions from earlier chats

### Git release workflow

The user clarified that a release explanation should cover the full route to `main`, not just tag syntax:

1. Develop work on feature branches and merge it into `develop`.
2. Create a release branch such as `release/v1.4.0` from updated `develop` and stabilize it.
3. Merge the release branch to `main`; tag the exact released commit on `main` with an annotated tag such as `v1.4.0`.
4. Carry release-only fixes back into `develop`.

The branch is a movable preparation line; the version tag identifies a specific commit. Useful checks documented in the Git lesson include `git show --no-patch --decorate v1.4.0`, `git log -1 --oneline main`, and `git ls-remote --tags origin v1.4.0`.

### Docker and Compose

- A user-defined bridge network lets containers reach each other by service/container name and container port. Host processes use the published host address and port.
- Publishing a new image does not update an already-running container; deployment is a separate action.
- Compose `depends_on` alone orders startup; a health check plus `condition: service_healthy` is needed when a dependent service must wait for readiness.
- `docker compose down` normally keeps named volumes. `docker compose down -v` removes declared volumes and can delete data.
- MongoDB initialization variables apply when the data directory is first initialized; a retained volume keeps the existing database state.
- Prior Compose examples were structurally checked with dummy credentials, but containers were not started. Do not represent those docs checks as a live Compose deployment.

### Jenkins and Docker

- Establish which Docker Engine the CLI controls before troubleshooting. With `/var/run/docker.sock` mounted into Jenkins, the CLI inside Jenkins sends requests to the host Docker Engine.
- Socket-mounted Jenkins is Docker-outside-of-Docker. Registry DNS, daemon trust/CA configuration, insecure-registry settings, and the actual push connection are governed by that host Engine. A Docker network visible inside Jenkins does not automatically give the host daemon the same DNS view.
- The Docker socket grants powerful host-level control. Do not recommend `chmod 777` as a fix. Use a deliberate access model; a separate Docker-in-Docker service with TLS is an alternative architecture with its own setup and risk.
- Install the Docker CLI in the Jenkins environment when it is needed; do not install a second daemon when intentionally using the host socket.
- Keep Jenkins home on persistent storage. Use Jenkins credentials for registry credentials and avoid printing secrets.
- A Docker image reference looks like `registry-host[:port]/owner/image:tag`; do not put `http://` or a Nexus UI path such as `/repository/...` in it.
- The host daemon must be able to resolve and reach the registry connector. Prefer a stable hostname/IP and published connector port over a temporary container IP.

### Nexus publishing and storage

- Use a Nexus Docker repository/connector with Docker login, tag, and push for container images. Raw repository REST upload is for standalone files such as binaries, not Docker manifests and layers.
- Keep Docker registry references distinct from Nexus browser/UI and Raw repository URLs.
- A Nexus repository exposes artifact access; a blob store holds file content. A repository normally selects one blob store, and multiple repositories may share a store.
- Filesystem capacity is shared by directories on the same filesystem. Nexus's service account needs read/write/traverse access to the configured store path.
- Cleanup policy criteria select eligible components; a scheduled cleanup task determines when they are processed. Cleanup can soft-delete content; a separate blob-store compaction task is what reclaims disk space. A policy schedule and the retention rule are different controls.
- A Nexus HTTP Docker connector accessed by a Docker daemon may require explicitly configured insecure-registry support. HTTPS with trusted certificates is preferable where available. In socket-mounted Jenkins, this trust/configuration belongs to the host daemon.

### Gitea packages and registry

- A package belongs to a user or organization. Linking the package to a repository is a separate action.
- Docker image names must use lowercase owner/image path segments. `docker tag` adds a registry-qualified name without rebuilding the image.
- A successful push ends with a digest, but publishing does not deploy the image.
- Use a personal access token with package-write rights for CI registry publishing; the built-in `GITEA_TOKEN` was not considered sufficient for package-registry publishing in the referenced workflow.
- Package storage's absolute location depends on Gitea's data path, storage configuration, and container mounts. Inspect the active configuration and mounts instead of assuming a host path.
- Gitea deduplicates content. Removing a package reference may not immediately free space if blobs remain referenced or cleanup has not run.

### Jenkins Go PATH issue

In the prior Jenkins container case, `/usr/local/go/bin/go version` worked while `go` was absent from the job's `PATH`. The cause was PATH discoverability, not executable permissions. Root's shell profile does not reliably configure Jenkins non-interactive jobs. The saved UI approach is **Manage Jenkins → System → Global properties → Environment variables**, with name `PATH+GO` and value `/usr/local/go/bin`; verify using `command -v go` and `go version`.

## Relevant prior chat history

These are summaries of Nana-related chats found in the available chat history and saved project memory. They capture decisions and outcomes; consult the original Codex chat when exact command output or the complete exchange matters.

| Approx. date | Chat or workstream | Main outcome |
|---|---|---|
| 2026-10-01 | Nana prompt/style and Nexus Blob refactor | Established concise, example-led style; clarified repository versus blob store, filesystem capacity, permissions, quotas, and cleanup. Stale intermediate Blob versions were removed after the user confirmed they were obsolete. |
| 2026-10-01 | Nexus cleanup-policy reference | Explained eligibility criteria versus task schedule and soft deletion versus compaction/reclaim. |
| 2026-10-03 to 2026-10-05 | Docker references and Compose | Created Docker workflow, Compose, and command references; covered networking, ports, health checks, volumes, MongoDB initialization, and safe lifecycle commands. Syntax/structure checks did not start containers. |
| 2026-10-05 | Docker and Nexus reference workflow | Reinforced that rough source is preserved when useful, and that new references should be concise, practical, diagrammed, and linked. |
| 2026-10-08 | Artifact Management reorganization and Gitea notes | Grouped Nexus notes under `Nexus/`, added an artifact-management overview, and created four Gitea guides for packages, registry, CI publishing, and storage. Captured package ownership/linking, token and lowercase naming requirements, storage paths, deduplication, and troubleshooting. |
| 2026-10-13 | Git NANA learning path and release workflow | Split Git material into a hub plus ten ordered lessons. Clarified feature → develop → release branch → main → annotated tag, including carrying release fixes back to develop. |
| 2026-10-14 to 2026-10-15 | Jenkins sizing, Basics, and Go PATH | Added Jenkins basics and resource-sizing guidance; isolated `go: not found` to missing `/usr/local/go/bin` in job PATH and recorded Jenkins's `PATH+GO` convention. |
| 2026-10-21 | Jenkins-to-Nexus Docker publishing (`Fix missing Docker image tag`) | Worked through mismatched tags, invalid image-reference syntax, Docker daemon DNS, unpublished ports, and HTTP/HTTPS mismatch. Confirmed socket-mounted Jenkins means host Docker Engine connectivity/trust rules apply. |
| 2026-10-21 | Nexus publishing and Docker in Jenkins references | Added a Nexus guide comparing same-host and cross-host setups and a Jenkins Docker reference linking to it. Preserved the old Jenkins rough note. Documentation checks were structural; no live Docker build or Jenkins job ran. |

## Operational troubleshooting sequence

For a Jenkins image-push failure, diagnose one layer at a time:

1. Compare the built tag with the exact image reference being pushed (`docker image ls` and the build/tag commands).
2. Check image-reference syntax: registry host and optional port, then lowercase path and tag; no scheme or Nexus UI path.
3. Identify whether Jenkins uses the host socket or a separate Docker daemon.
4. From the daemon that performs the push, check DNS and reachability to the published Nexus Docker connector port.
5. Match the connector protocol: HTTPS with trusted CA, or explicitly configured HTTP insecure registry where that is the chosen setup.
6. Check registry authentication and repository permissions only after address and protocol work.

This ordering reflects the prior debugging chats; it prevents a tag, DNS, port, or protocol failure from being mislabeled as an authentication problem.

## Guidance for the next LLM session

- Read this file for orientation, then inspect the exact target note, its local prompt, the Nana prompt, and related current notes before editing.
- Verify facts likely to drift (tool syntax, supported versions, product behavior) against current official documentation when making new technical claims.
- Preserve the learner-facing distinction between a concept, a worked example, and an environment-specific setting. Do not turn an example value into a claim about the user's live deployment.
- Keep edits within the requested topic and preserve useful existing content and canonical links.
- Do not use prior chat summaries as runtime evidence. State whether checks were formatting, static syntax, local execution, or live service validation.
- The vault's user-provided `AGENTS.md` also requires the shared instructions in `/home/ryan/.codex/RTK.md`; prefix shell commands with `rtk`.
