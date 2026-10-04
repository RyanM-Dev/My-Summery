# AI Prompt — grpc-babal

> **Charts and diagrams:** Use fenced `mermaid` blocks for charts, flows, architecture, hierarchies, and relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` as appropriate. Use clear labels and Obsidian-compatible syntax. Keep runnable code, commands, literal output, payloads, and calculations in their original code formats.


> **Fundamental prompt (required first):** [[AI Prompt/00 - General Study Guide Prompt]]  
> The General prompt is the **foundation** for every study guide — structure, Q&A format, diagrams, source handling, and quality bar. It **cannot be used alone**. Always copy it first, then append this project prompt and [[Go/AI Prompt - Go]].

> **Purpose:** Project-specific layer for *gRPC Microservices in Go* (Babal) with *Microservices with Go, 2nd Edition* (Shuiskov) as comparison, and **internet research** as the third source.
>
> **Primary book:** Babal drives chapter structure, terminology, e-commerce architecture (Fig 1.1), and companion-repo patterns.
> **Comparison book:** Shuiskov adds microservice theory, Go idioms, sync/async communication, and alternative approaches.
> **Third source:** Internet (official docs, RFCs, grpc-go releases, Buf, OpenTelemetry) — not a fourth book.

---

## Source directories

| What | Directory | Contents |
|------|-----------|----------|
| **This prompt** | `Go/gRPC Microservices in Go/` | `AI Prompt - grpc-babal.md` (this file) |
| **Study guides & chapter notes (output)** | `Go/gRPC Microservices in Go/` | `*.md` study guides; raw notes as `{N}.OLD-{suffix} - {Title}.md` |
| **Book PDFs** | Repo root (`Book/`, sibling of `Obsidian/`) | `Hüseyin_Babal_gRPC_Microservices_in_Go_Manning_Publications_2023.pdf`, `Microservices with Go 2nd Edition Alexander_Shuiskov.pdf` |
| **Related Go notes** | `Go/gRPC/` | Proto/codegen cheat sheets |
| **Companion code (supplementary)** | GitHub | `github.com/huseyinbabal/microservices`, `github.com/huseyinbabal/microservices-proto` |

---

## Book roles

| Role | Book | Path | Use in study guides |
|------|------|------|---------------------|
| **Primary (Ref 1)** | *gRPC Microservices in Go* (Manning, Hüseyin Babal, 2023) | `Hüseyin_Babal_gRPC_Microservices_in_Go_Manning_Publications_2023.pdf` | Structure, gRPC mechanics, hexagonal layout, Order/Payment flow, K8s/CI/CD/observability narrative |
| **Comparison (Ref 2)** | *Microservices with Go, 2nd Edition* (Alexander Shuiskov) | `Microservices with Go 2nd Edition Alexander_Shuiskov.pdf` | Microservice foundations, Go project structure, sync comm (Ch.5), serialization (Ch.4), patterns catalog |
| **Third (Ref 3)** | Internet / official docs | Web search enabled | {{CURRENT_YEAR}} deprecations (`grpc.Dial` → `grpc.NewClient`), Buf, OTel, health probes, industry consensus |

**Companion codebase (supplementary — cite when relevant, not Ref 3):** `https://github.com/huseyinbabal/microservices` (Order, Payment, `microservices-proto` module)

**Set in general prompt:** `{{PRIMARY_REFERENCE}}` = `1` (Babal leads; Shuiskov enriches; internet adds {{CURRENT_YEAR}} column)

**Output folder:** `Go/gRPC Microservices in Go/`

**Exemplar study guide:** [[Go/gRPC Microservices in Go/1. Introduction to Go gRPC microservices - Study Guide]] — match its tone, comparison tables, and Obsidian callout format.

---

## Babal chapter map (use for part titles)

| Ch | Babal Title | Shuiskov cross-ref (when useful) | Repo touchpoint |
|----|-------------|----------------------------------|-----------------|
| 1 | Introduction to Go gRPC microservices | Ch.1 Introduction to Microservices | Architecture only (Fig 1.1) |
| 2 | gRPC meets microservices | Ch.1 monolith vs MS; scale cube; Ch.6 sagas preview | — |
| 3 | Getting up and running with gRPC and Golang | Ch.4 Serialization; Ch.5 gRPC intro | `microservices-proto` module |
| 4 | Microservice project setup | Ch.2 Go structure; hexagonal patterns | `order/`, `payment/` layout |
| 5 | Interservice communication | Ch.5 Synchronous Communication | Order → Payment gRPC stub |
| 6 | Handling errors and timeouts | Ch.2 context; error handling | Client interceptors |
| 7 | Storing service data | Ch.7+ persistence topics | MySQL adapters |
| 8 | Testing microservices | Ch.2+ testing patterns | `bufconn`, mocks |
| 9 | Observability | Ch.12 telemetry | OTel on Payment |
| 10+ | Deployment, security, advanced topics | Ch.8 K8s; Ch.10+ patterns | `deployment/` manifests |

When Shuiskov has no matching chapter, set `{{REFERENCE_2}}` to the closest section and note partial overlap in the Sources table.

---

## Ch.1 Introduction — concepts to preserve (from vault notes)

Mirror these themes when generating Ch.1 (or when any chapter revisits foundations):

**Babal Ch.1.1 — Why gRPC**
- Go suitability: fast compile, standalone binaries, K8s-native
- gRPC: Protobuf + HTTP/2, stub codegen, idempotency, TLS, four streaming modes
- **Stub definition:** generated client proxy implementing remote methods locally

**Babal Ch.1.2–1.3 — REST vs RPC**
- Comparison table: payload, protocol, latency, browsers, streaming
- gRPC for internal backend-to-backend; REST (or gRPC-Web/gateway) for public/browser APIs
- Pitfall: proto maintenance overhead for 1–2 service startups

**Babal Ch.1.4 — Production-grade ecosystem (Fig 1.1)**
- Five business-capability services: Product, Cart, Checkout, Payment, Shipping
- K8s runtime, CI/CD (proto codegen, tests, images), observability triad (metrics/logs/traces)
- API gateway for auth, rate limiting, public access
- Fast-fail on missing env vars (e.g. `DATA_SOURCE_URL`)

**Shuiskov Ch.1 — Microservice theory**
- Definition: services by business capability
- Monolith pain: build time, coupled deploys, blast radius, uniform scaling
- When NOT to split early; design for failure; embrace automation

**Three-way comparison columns (required in opening):** Babal | Shuiskov | {{CURRENT_YEAR}} internet

---

## Suggested general-prompt inputs (copy & adjust per chapter)

| Field | Example (Ch.1) |
|-------|----------------|
| `{{TOPIC_DOMAIN}}` | Go gRPC microservices |
| `{{CHAPTER_NUM}}` | 1 |
| `{{CHAPTER_TITLE}}` | Introduction to Go gRPC Microservices |
| `{{OUTPUT_FILENAME}}` | `1. Introduction to Go gRPC microservices - Study Guide.md` |
| `{{OUTPUT_FOLDER}}` | `Go/gRPC Microservices in Go/` |
| `{{REFERENCE_1}}` | gRPC Microservices in Go (Babal) — Ch.1 |
| `{{REFERENCE_1_PATH}}` | `Hüseyin_Babal_gRPC_Microservices_in_Go_Manning_Publications_2023.pdf` |
| `{{REFERENCE_2}}` | Microservices with Go, 2nd Edition (Shuiskov) — Ch.1 |
| `{{REFERENCE_2_PATH}}` | `Microservices with Go 2nd Edition Alexander_Shuiskov.pdf` |
| `{{REFERENCE_3}}` | Internet: gRPC Go docs, Buf, grpc-go releases, OpenTelemetry |
| `{{REFERENCE_3_PATH}}` | (web search — no local path) |
| `{{PRIMARY_REFERENCE}}` | `1` |
| `{{USE_INTERNET_RESEARCH}}` | yes |

**Ch.5 example:**

| Field | Value |
|-------|-------|
| `{{CHAPTER_NUM}}` | 5 |
| `{{CHAPTER_TITLE}}` | Interservice Communication |
| `{{OUTPUT_FILENAME}}` | `5. Interservice communication - Study Guide.md` |
| `{{REFERENCE_1}}` | Babal — Ch.5 Interservice communication |
| `{{REFERENCE_2}}` | Shuiskov — Ch.5 Synchronous Communication |

---

## When to use

- Any chapter from *gRPC Microservices in Go* (Babal)
- Study guides saved under `Go/gRPC Microservices in Go/` as `{N}. {Title} - Study Guide.md`
- Raw chapter notes use vault OLD pattern: `{N}.OLD-{suffix} - {Title}.md` (e.g. `1.OLD-intro - Introduction to Go gRPC microservices.md`, `3.OLD-notes - Getting up and running with gRPC and Golang.md`) — source material; do not overwrite
- When regenerating a study guide, rename to `{N}.OLD-guide - {Title} - Study Guide.md` first (see [[AI Prompt/00 - General Study Guide Prompt]])
- **Prompt stack:** General → Go → this file (see [[#Regenerate workflow]])

---

## Specific prompt (append after general + Go prompts)

```
## TOPIC-SPECIFIC INSTRUCTIONS — grpc-babal ({{CURRENT_YEAR}})

You are a senior Go microservices engineer teaching from *gRPC Microservices in Go* (Babal). Apply these rules ON TOP of the general prompt AND the Go-specific prompt.

---

### Source hierarchy (CRITICAL)

1. **Babal (Ref 1) is ground truth** for chapter structure, gRPC implementation steps, hexagonal architecture, e-commerce reference diagram (Fig 1.1), Order/Payment narrative, and companion-repo conventions.
2. **Shuiskov (Ref 2)** adds microservice theory, Go idioms, synchronous/async communication models, load-balancing strategies, and pattern-catalog depth — use in comparison columns, not as the narrative voice.
3. **Internet (Ref 3)** supplies {{CURRENT_YEAR}} senior practice: deprecated API replacements, Buf workflows, OpenTelemetry instrumentation, gRPC health service, official doc links, and verifiable release notes.

**Supplementary (not a numbered reference):** Babal companion repo `https://github.com/huseyinbabal/microservices` — cite for "what the book implements" in Part 7 repo snapshot. Verify claims against repo structure when mentioned.

**Every comparison table MUST have columns:** Babal | Shuiskov | {{CURRENT_YEAR}} (internet)

When the opening Sources table allows a fourth column for repo, label it **Repo** — separate from Ref 3.

**Framing rules:**
- Write for a developer building production gRPC microservices in Go.
- Babal drives the story; Shuiskov explains *why* microservices and *what alternatives exist*.
- Internet column flags outdated book APIs (e.g. `grpc.Dial`, `WithInsecure()`, raw `protoc` without Buf).
- Do not invent Babal chapter content — extract from PDF/notes. If a section is preview-only in Ch.1, say so and point to the chapter where Babal details it.
- Cite paths repo-relative only — never `/home/...` or `C:\...`.

---

### Babal voice & architecture to preserve

- **Business-capability decomposition:** Product, Cart, Checkout (Order), Payment, Shipping — not org-chart services.
- **Contract-first APIs:** `.proto` in VCS; generated stubs as versioned Go modules — never copy-pasted `.pb.go`.
- **Hexagonal architecture (Ch.4+):** `ports/` (interfaces) + `adapters/` (gRPC, DB, payment client); core logic transport-agnostic.
- **Internal gRPC, external REST:** API gateway or gRPC-Web for browser/public consumers.
- **Design for failure:** idempotency, retries, circuit breakers, `context.Context` deadlines — network calls replace in-process calls.
- **Observability triad:** metrics + logs + traces; trace propagation from gateway through all stubs.
- **Stub definition (Ch.1):** generated client-side proxy implementing remote service methods locally.

---

### Shuiskov — when to pull from Ref 2

| Babal topic | Shuiskov enrichment |
|-------------|---------------------|
| Ch.1 intro | Microservice definition, monolith pros/cons, when not to split |
| Ch.2 scaling / sagas | Scale cube, saga choreography vs orchestrator |
| Ch.3 proto / codegen | Ch.4 serialization deep dive |
| Ch.4 hexagonal setup | Ch.2 Go project structure, movie MS scaffold |
| Ch.5 interservice comm | Ch.5 sync comm, four RPC types, server vs client LB |
| Ch.6+ errors, data, test | Matching Shuiskov parts 2–3 chapters |
| Observability | Ch.12 telemetry patterns |

When Shuiskov disagrees with Babal (e.g. `grpc.Dial` examples, manual stub copy), show both + recommend {{CURRENT_YEAR}} internet-backed practice.

---

### Content requirements (in addition to Go prompt)

#### Code blocks
- Show **Babal (book)** vs **Repo (companion)** vs **{{CURRENT_YEAR}} recommended** for every major gRPC pattern.
- Minimum from Go prompt still applies: 3+ Go blocks, 1+ bash (`buf generate`, `grpcurl`, `go test`), 1+ protobuf/YAML when chapter touches contracts or K8s.

#### Directory trees (mandatory — minimum 2)
- Babal hexagonal layout: `cmd/`, `internal/application/core/`, `internal/ports/`, `internal/adapters/grpc/`, `deployment/`
- Proto module layout when Ch.3+

#### Diagrams (mandatory)
- **Ch.1:** Fig 1.1 architecture Mermaid + REST vs gRPC table + checkout sequence mermaid
- **Ch.2:** Scale cube, monolith vs microservices, saga flow
- **Ch.3–4:** Proto → stub codegen pipeline; hexagonal ports/adapters
- **Ch.5+:** Order → Payment stub call sequence; load-balancing (server-side K8s vs client-side)
- Default: chapter's main request path as mermaid `sequenceDiagram`

#### Three-source comparison (mandatory in opening + Part 8)
Revisit in Part 8 senior recommendations:

| Topic | Babal | Shuiskov | {{CURRENT_YEAR}} |
|-------|-------|----------|------------------|
| Example | `grpc.Dial` in text | `grpc.Dial` in examples | `grpc.NewClient` + OTel handler |

**Typical Babal vs Shuiskov differences to highlight:**

| Area | Babal angle | Shuiskov angle |
|------|-------------|----------------|
| Narrative | Hands-on e-commerce Order/Payment | Movie microservices + pattern catalog |
| Architecture | Fig 1.1 five-service K8s diagram | Broader microservice theory first |
| Codegen | `protoc` flags, published proto module | Directory-based `protoc`, module imports |
| Communication | gRPC-first internal RPC | Sync (gRPC) vs async (events) decision framework |
| Load balancing | K8s Service (server-side) emphasis | Server-side vs client-side comparison |

#### {{CURRENT_YEAR}} standards (Part 7–8 — from internet Ref 3)
- `grpc.NewClient` (not `grpc.Dial`)
- `context.Context` end-to-end; OTel `otelgrpc` client + server handlers
- Buf: `buf lint`, `buf breaking`, `buf generate`
- Idempotency keys on non-idempotent RPCs (`Create`, `Charge`)
- `google.rpc` status details; correct `codes.*` mapping
- gRPC health service for K8s probes; graceful `GracefulStop()`
- Version-pinned proto modules and grpc-go — not `@latest`
- Testing: `bufconn`, testcontainers, port/adapter mocks

---

### Part arc (adapt titles to Babal chapter)

| Part | Focus |
|------|-------|
| 1 | Babal core theory + chapter objectives |
| 2 | Babal implementation mechanics (code, proto, adapters) |
| 3 | Shuiskov comparison — theory, alternatives, trade-offs |
| 4 | Infrastructure: K8s, CI/CD, networking, discovery |
| 5 | Errors, edge cases, operational concerns |
| 6 | Combined best practices (Babal + Shuiskov + {{CURRENT_YEAR}}) |
| 7 | Companion repo snapshot — ahead vs gaps (verify against GitHub) |
| 8 | Senior recommendations + end-to-end mermaid + Master Q&A |

Rename part titles to match chapter content (e.g. Ch.5 → "Synchronous Stubs", "Order→Payment Flow", "Load Balancing Strategies").

---

### Part 7 — Repo snapshot (companion codebase)

#### ✅ Ahead of the books
Table: Area | Status — verify against `github.com/huseyinbabal/microservices`

#### ⚠️ Gaps vs {{CURRENT_YEAR}} senior standards
Table: Gap | Risk | Priority (🔴 High / 🟡 Medium)

Common gaps to check (do not assume — verify):
- `grpc.Dial` in book text vs `grpc.NewClient` in repo
- Missing idempotency keys on `Create` RPCs
- Partial OTel instrumentation
- No Buf breaking checks in CI
- Missing gRPC health service

---

### Part 8 — Mermaid diagram (required)

Reflect the chapter's main flow:
- Ch.1: User → Gateway → Checkout → Cart/Payment/Shipping (with OTel)
- Ch.4: Driver adapter (gRPC server) → core → driven adapters (DB, payment stub)
- Ch.5: Order `PlaceOrder` → Payment `Create` stub sequence
- Default: Babal's primary use case for that chapter

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks` before Further Reading.

Tasks must:
- Align with Babal chapter progression (Ch.1: Buf hello-world; Ch.4: scaffold hexagonal folder; Ch.5: consume payment stub)
- Use {{CURRENT_YEAR}} tooling (`buf`, `grpc.NewClient`, `grpcurl`)
- Be verifiable with command output or test pass

**Template:**
#### Task N: {{TITLE}}
**Goal:** ...
**Files:** ...
**Steps:**
1. ...
**Done when:** ...

---

### Interview Q&A focus

Weight questions toward:
- Babal chapter concepts (stubs, proto compatibility, hexagonal ports, Fig 1.1 services)
- Shuiskov theory angles (when monolith wins, saga types, sync vs async)
- {{CURRENT_YEAR}} practicals: `Dial` deprecation, Buf breaking changes, idempotency, trace propagation
- Architecture: "why gRPC internal, REST external?", "what is a stub?", "how does Checkout use Payment?"

Include Master Q&A in Part 8 (5+ questions).

---

### Obsidian note conventions (match Ch.1 exemplar)

- Emoji headers from general prompt (`📖`, `📌`, `🔧`, `💻`, `📊`, `🧪`, `💡`, `🔨`, `📚`)
- Collapsible answers: `> [!success]- 📋 Answers — Part N`
- Filename: `{N}. {Babal Chapter Title} - Study Guide.md` or `{N}. {topic} - Study Guide.md` matching vault style
- Link prior vault notes when they exist (e.g. raw summaries, [[Go/gRPC/1.Generate command]])
- Footer regenerate block (include in output):

### 🤖 Regenerate this guide

1. [[AI Prompt/00 - General Study Guide Prompt]] — fill Inputs (Babal = Ref 1, Shuiskov = Ref 2, Internet = Ref 3, PRIMARY = 1)
2. [[Go/AI Prompt - Go]] — append Go code/tree requirements
3. [[Go/gRPC Microservices in Go/AI Prompt - grpc babal]] — append this project prompt

---

### Further Reading (required)

Include:
- Next Babal chapter(s)
- Matching Shuiskov chapter(s)
- Official links: [gRPC Go docs](https://grpc.io/docs/languages/go/), [Buf docs](https://buf.build/docs), [grpc-go](https://github.com/grpc/grpc-go), [microservices.io](https://microservices.io)
- Companion repo: `https://github.com/huseyinbabal/microservices`

---

Generate the complete study guide now.
```

---

## Quick reference — Shuiskov mapping by Babal chapter

| Babal Ch | Shuiskov chapter / topic |
|----------|--------------------------|
| 1 | Ch.1 Introduction to Microservices |
| 2 | Ch.1 + Ch.6 (sagas) + service discovery patterns |
| 3 | Ch.4 Serialization formats |
| 4 | Ch.2 Go structure and tooling |
| 5 | Ch.5 Synchronous Communication |
| 6 | Ch.2 context, error handling |
| 7 | Persistence / database chapters |
| 8 | Testing chapters |
| 9 | Ch.12 Observability / telemetry |
| 10+ | Ch.8+ deployment, security, advanced patterns |

When mapping is weak, note "partial overlap" in the Sources table and lean on internet Ref 3 for depth.

---

## Regenerate workflow

1. [[AI Prompt/00 - General Study Guide Prompt]] — fill Inputs table; copy **Full general prompt** (fundamental layer)
2. [[Go/AI Prompt - Go]] — append Go code/tree requirements
3. **This file** — append **Specific prompt** section below
4. Save output to `Go/gRPC Microservices in Go/{chapter} - Study Guide.md`