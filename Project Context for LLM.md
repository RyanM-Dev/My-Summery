# Obsidian Vault — Project Context for LLMs

> **Purpose:** Give a future LLM enough context to continue work in this Obsidian vault without repeating discovery or confusing vault material with separate software projects.
>
> **Updated:** 2026-10-23
>
> **Coverage:** This is a project map and a summary of relevant prior Codex work, not a transcript or an exhaustive inventory of every note. Chat history is summarized from accessible Codex threads and saved project memory; consult the original chat when exact wording, command output, or full evidence matters.

## Project identity

This repository is Ryan’s personal Obsidian learning and technical-reference vault. It contains study notes, course material, project-specific prompts, reference guides, and navigation/index notes across software engineering and DevOps topics. It is a content project: the Markdown files and Obsidian links are the main artifacts, and each technical topic has its own source boundaries and conventions.

The project should not be treated as one unified course. The top-level areas include:

- `AI Prompt/` — shared general and teaching prompts.
- `Agentic Development/` — agent development material.
- `Brokers/` — broker and messaging notes, including MQTT and Kafka.
- `Database/` — PostgreSQL, MySQL, and MongoDB material.
- `DevOps/NaNa/` — Nana DevOps course notes and its course prompt.
- `Git/` — Git reference notes.
- `Go/` — Go language, concurrency, testing, HTTP, gRPC, and book-based learning material.
- `Image/`, `LinkedIn/`, `Linux/`, and `Search Engines/` — additional reference and learning areas.
- `Resume/` — a separate résumé workspace inside the vault.

The vault’s folder and filename conventions matter to Obsidian navigation. Use spaces between words in folder and note names. Preserve technical spellings such as `go-cmp`, `.OLD-` archive markers, and ` - ` between lesson numbers and titles. After any rename, update Obsidian links and saved application references.

## Instructions and working rules

Read and follow `/home/ryan/.codex/RTK.md` for shell work: prefix every shell command with `rtk`. The repository instructions are in `AGENTS.md`. The user may also provide updated instructions in the current chat; current user instructions take precedence over older prompt conventions.

### Prompt discovery and precedence

Before drafting or substantially editing a study note:

1. Inspect the target directory and each parent up to the vault root for `AI Prompt - {topic or project}.md`.
2. Read the shared general and teaching prompts in `AI Prompt/`.
3. Read all applicable topic, course, and project prompts. The most specific applicable prompt wins when rules conflict.
4. Inspect the target note and related index/reference notes before changing structure or replacing content.

Prompts belong beside the material they govern. Keep only shared prompts in `AI Prompt/`; do not create duplicate or redirect-only copies of local prompts.

### Teaching and writing conventions

- Teach in this order: understanding, real-world example, mechanics, technical details, advanced use.
- Explain concepts in direct, concise prose; silently correct errors and explain calculations step by step.
- Explain commands by purpose, syntax, flags/arguments, examples, and expected result.
- Use practical examples throughout, and use Mermaid diagrams for workflows, architecture, relationships, and exchanges. Keep diagrams Obsidian-compatible.
- Include understanding-focused questions with collapsed answers, and safe hands-on exercises where appropriate.
- End study notes with **Quick Reference**, **Things to Remember**, **Pro Tips**, and **Related Topics** using Obsidian links, subject to the more specific prompt.
- Preserve runnable code, commands, literal output, payloads, and calculations in their original code formats.
- Never claim an example or runtime check was executed unless it actually was. Separate structural checks from builds, tests, live services, hardware, or production validation.

### Topic-specific rules

- **Nana DevOps:** Follow `DevOps/NaNa/AI Prompt - NANA DevOps.md`. Use concise, direct, example-led teaching; generous whitespace; meaningful emojis; grouped and correctly indented Markdown; and useful Mermaid diagrams. The Nexus Repository note is a prose/style reference; the Blob note is a diagram reference. Keep exercises compact and caveats proportional.
- **MQTT:** Follow `Brokers/MQTT/AI Prompt - mqtt.md`; for Go examples also follow `Go/AI Prompt - Go.md`. Consult the relevant vault references, distinguish MQTT protocol versions from client-library versions, and keep code proportional to the lesson.
- **Database notes:** Database material belongs under `Database/` (currently PostgreSQL, MySQL, MongoDB). PostgreSQL chapter references are organized under `Database/Postgresql/Art of PostgreSQL/`.
- **Vault organization:** Search existing material and preserve the established structure; do not create new project prompts unless a project-specific prompt is actually needed.

### Reliability and evidence

- Inspect the exact checkout, target note, and relevant local instructions before editing.
- Prefer exact source evidence over assumptions. Distinguish teaching examples from observed deployments and conceptual recommendations from implemented behavior.
- Do not modify unrelated files. Check Git status before and after changes when working in the repository.
- Validation should match the claim: Markdown whitespace/link checks do not prove that code compiles or a service works. Record what was and was not run.
- For résumé material, keep professional work, independent practice, and study distinct. Use factual source material and retain `[NEEDS CONFIRMATION: ...]` instead of inferring ownership or skill level.

## Vault map and key navigation files

| Area | What it contains | Entry points / local guidance |
|---|---|---|
| Nana DevOps | Linux, Git, Docker, CI/CD/Jenkins, artifact management, Nexus and Gitea | `DevOps/NaNa/0. Nana DevOps Table of Contents.md`; `DevOps/NaNa/Project Context for LLM.md`; `DevOps/NaNa/AI Prompt - NANA DevOps.md` |
| Go | Learning Go, Concurrency in Go, gRPC, HTTP, testing, scattered references | `Go/AI Prompt - Go.md`; `Go/Learning Go/README.md`; `Go/Concurrency in Go/README.md` |
| MQTT | MQTT 3.1.1 reference course and broker notes | `Brokers/MQTT/AI Prompt - mqtt.md`; `Brokers/MQTT/HiveMQ 3.1.1/0. MQTT Table of Contents.md` |
| Databases | PostgreSQL, MySQL, MongoDB references and book notes | Topic prompts live inside each database folder; PostgreSQL Art of PostgreSQL notes are under `Database/Postgresql/Art of PostgreSQL/` |
| Other topics | Linux, Git, Meilisearch, Kafka, LinkedIn and other references | Look for local prompts in the target folder and parents |
| Resume | Résumé-specific work and source artifacts | Treat as its own workspace; do not fold résumé claims into general vault skill summaries |

The Nana subsection has its own chat history and technical concepts in [[DevOps/NaNa/Project Context for LLM]]. This vault-level document summarizes the other major workstreams and the shared rules.

## Prior chat history and outcomes

The entries below summarize work that is relevant to the vault. Dates are approximate when inferred from chat records. Codex titles are quoted so they can be searched in the task list.

### Nana DevOps and Jenkins (2026-10-01 to 2026-10-21)

This is the most active vault workstream. Prior chats established a concise, example-led teaching style and built or reorganized notes across Docker, Compose, Nexus, Gitea, Git, and Jenkins. See [[DevOps/NaNa/Project Context for LLM]] for the detailed Nana note map and history.

Key concepts carried forward:

- Nexus repositories provide artifact access; blob stores hold artifact content. Cleanup policies identify eligible components, scheduled cleanup processes them, and blob-store compaction reclaims disk space after soft deletion.
- Gitea packages have an owner; linking a package to a repository is a separate action. Docker image references use lowercase owner/image path segments; tagging does not rebuild an image, and pushing does not deploy it.
- Git NANA was organized as a progressive hub plus ten numbered lessons. A release workflow goes from feature branches through `develop`, a release branch, stabilization, merge to `main`, and an annotated tag on the exact released commit. Carry release fixes back to `develop`.
- Jenkins job processes have their own environment. A Go binary installed at `/usr/local/go/bin` can exist while jobs still report `go: not found`; use job PATH configuration, Jenkins `PATH+GO`, or a repeatable image environment, then verify with `command -v go` and `go version`.
- With `/var/run/docker.sock` mounted, Jenkins Docker commands use the host Docker Engine. Registry DNS, published connector ports, TLS/CA trust, and insecure-registry configuration must work for that daemon. Diagnose the exact layer: tag, image-reference syntax, DNS, port, protocol, then authentication.
- Docker Registry image references contain a registry host/port and image path/tag; they do not contain `http://` or a Nexus UI path such as `/repository/...`. Raw repositories use HTTP file operations, not the Docker Registry push protocol.
- The saved Jenkins sizing discussion separates tiny documented minimums from practical starting capacity and notes that build-heavy workloads belong on agents. Exact sizing depends on workload.

Relevant chat titles include **“Create Jenkins Basics Obsidian reference,” “Estimate Jenkins controller and installation resources,” “Diagnose Jenkins `go: not found` after Go installation in a container,” “Fix missing Docker image tag,”** and **“Reorganize Git NANA learning path and release workflow.”**

### Go learning material (2026-10)

Recent vault chats expanded the book-based Go references and navigation:

- **“Create Concurrency in Go references”** added the remaining references for chapters 5 and 6 and a chapter map. Topics include error propagation, cancellation, heartbeats, replicated requests, rate limiting, recovery, runtime scheduling, work stealing, `GOMAXPROCS`, and runtime diagnostics.
- **“Create Learning Go chapter reference”** added guides for chapters 10–13, 15, and 16, and updated the folder progress list. Chapter 14 is absent from the referenced second edition. Existing concurrency/testing notes are linked where useful rather than duplicated.
- The broader Go area also has study notes and references for types, interfaces, functions, pointers, generics, errors, modules, tooling, standard library, testing, HTTP, gRPC, and concurrency. Find the current canonical notes through their folder indexes and local prompts rather than relying on old filenames.

These chats checked structure, links, and filenames. Such checks did not execute all embedded examples or validate every claim against a running Go program.

### PostgreSQL and skills inventory (2026-10)

- **“Create PostgreSQL chapter references”** generated reference notes for the chapters in *The Art of PostgreSQL* and linked them into the existing book area. Preserve the book’s chapter organization and existing references when extending it.
- **“Create Obsidian skills inventory”** created a separate inventory from vault content with explicit exclusions. Treat it as a catalog of material studied or documented, not proof of professional experience or proficiency. The exact requested scope is in the original chat.

### Résumé workspace (2026-10)

The chat **“Create a new resume”** worked in the nested `Resume/` workspace. Keep résumé facts source-backed and separate evidence of employment from independent practice and study. Do not infer skill level, ownership, or production experience merely from notes or repository contents; retain confirmation placeholders when facts are missing.

### Cross-project boundaries

Some related chats in the wider Codex history concern separate repositories, especially the `mqtt-5-learning` course and `robot-market-vending-backend`. Those codebases are not part of this Obsidian vault. The vault may contain notes or learning references connected to them, but do not present repository state, runtime behavior, or implementation status from those chats as a fact about this vault. For the MQTT course, the remembered distinction is that learner HTML lives in `teaching/lessons/`, source lives in `teaching/course-source/`, and `COURSE.md` is the coverage anchor; consult that repository for current truth.

## Concepts and choices that recur across the vault

These are recurring teaching and documentation distinctions, not claims that every note already covers them:

- **Concept before procedure:** Explain the idea, ground it in a practical case, show mechanics, then introduce technical detail and advanced use.
- **Navigation is part of the content:** A new lesson should be reachable from its index/TOC. Renames require link and saved-reference updates.
- **Examples are not validation:** A code block is instructional until it is actually run. Clearly label illustrative output and unverified commands.
- **Configured versus active:** Configuration files, prompt text, or a note describing a system do not prove that a runtime service uses that configuration.
- **Protocol and implementation versions differ:** For MQTT and similar technology notes, keep protocol specification versions separate from library/broker versions.
- **Teach operational boundaries:** For DevOps, distinguish control plane from worker/agent, repository from storage, container from host daemon, authentication from authorization, and publishing from deployment.
- **Preserve provenance:** For résumé and skills artifacts, the existence of study notes is evidence of study/documentation, not automatically professional proficiency.

## Suggested startup procedure for a future LLM

1. Read this file and `AGENTS.md`; apply the user’s latest instruction if it changes scope.
2. Identify the target note and inspect the exact folder, sibling notes, parent indexes, and any applicable `AI Prompt - ...` files.
3. Check repository status and avoid disturbing unrelated changes.
4. Confirm whether the task concerns the vault itself or a separate repository referenced by a note/chat. Inspect that exact source before making code or runtime claims.
5. Make the smallest coherent update, preserve established filenames and links, then validate only what was actually requested. Report the validation boundary plainly.

## Chat-history coverage note

Codex’s thread list is a rolling, bounded view, and this document summarizes the relevant accessible records and saved project memory as of 2026-10-23. It intentionally omits unrelated personal chats and does not reproduce full transcripts. If a future task depends on an older exact exchange, search Codex by the chat title and inspect that thread before relying on this summary.
