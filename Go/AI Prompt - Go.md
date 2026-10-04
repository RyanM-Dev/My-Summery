# Go — Specific Prompt

> **Purpose:** Language layer for Go study guides (fundamentals, idioms, gRPC, HTTP, testing, microservices, concurrency, mistakes). Append AFTER [[AI Prompt/00 - General Study Guide Prompt]] — the fundamental prompt.
>
> **Babal gRPC track?** Also append [[Go/gRPC Microservices in Go/AI Prompt - grpc babal]] (project prompt in source folder).
> **Learning Go + Mistakes track?** Also append [[Go/Learning Go/AI Prompt - learning go]] (project prompt).

---

## When to use

- Go language fundamentals, idioms, project structure — pair with [[Go/Learning Go/AI Prompt - learning go]] for the Learning Go (Bodner) + 100 Mistakes track
- gRPC / microservices — pair with [[Go/gRPC Microservices in Go/AI Prompt - grpc babal]] for the Babal book track
- HTTP APIs (`net/http`, Gin, Echo)
- Testing (testify, testcontainers, dockertest)
- Concurrency (goroutines, channels, `sync`)

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — Go ({{CURRENT_YEAR}})

You are a senior Go engineer. Apply these rules ON TOP of the general prompt.

---

### Go content requirements

#### Code blocks (mandatory where implementation is discussed)
- Include **at least 3 Go code blocks** across the full document
- Include **at least 1 bash/shell block** (`go test`, `go mod`, `buf generate`, `grpcurl`, `docker compose`)
- Include **at least 1 YAML or protobuf block** when chapter touches config or API contracts
- Label blocks with repo-relative path: `microservices/order/internal/adapters/payment/payment.go`
- Show **book version** vs **repo version** vs **{{CURRENT_YEAR}} recommended** when they differ

Example pattern:
```go
// Babal (book)
conn, err := grpc.Dial(url, opts...)

// Repo (current)
conn, err := grpc.NewClient(url, opts...)

// {{CURRENT_YEAR}} senior standard
conn, err := grpc.NewClient(url,
    grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
    grpc.WithTransportCredentials(...),
)
```

#### Directory trees (mandatory — minimum 2)
Show ASCII trees for relevant packages when discussing project layout.

#### Compare implementations across sources
For every major pattern, use a comparison table across references + {{CURRENT_YEAR}}.

Highlight **different methods and approaches** — not just "they are the same." Explain trade-offs.

#### {{CURRENT_YEAR}} Go standards (apply in Part 7–8)
- `grpc.NewClient` (not `grpc.Dial`)
- `context.Context` end-to-end propagation
- OpenTelemetry (`otelgrpc` handlers)
- Buf for proto lint/breaking
- Idempotency keys on non-idempotent RPCs
- `google.rpc` error details + correct `codes.*` mapping
- Graceful shutdown: `GracefulStop()`, `conn.Close()`
- gRPC health service for K8s probes
- Version-pinned modules (not `@latest`)
- Test: `bufconn`, testcontainers, mock ports/interfaces

---

### Part 7 — Repo snapshot (when `{{REFERENCE_3}}` is a codebase)

#### ✅ Ahead of the books
Table: Area | Status

#### ⚠️ Gaps vs {{CURRENT_YEAR}} senior standards
Table: Gap | Risk | Priority (🔴 High / 🟡 Medium)

Verify every row against actual files. No invented features.

---

### Part 8 — Mermaid diagram (required)

Diagram must reflect the chapter's main flow (request path, service calls, or architecture).

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks` before Further Reading. Each task:
- **Goal** | **Files** | **Steps** | **Done when**
- Reflect {{CURRENT_YEAR}} best practices
- Be concrete and verifiable

---

### Go interview focus (where relevant)

- Error handling (`error` values, wrapping, `%w`)
- Interfaces and composition
- Concurrency patterns (worker pools, context cancellation)
- gRPC: unary vs streaming, stubs, status codes, idempotency
- Hexagonal architecture: ports and adapters
- Testing pyramid: unit mocks vs integration with testcontainers

---

Generate the complete study guide now.
```