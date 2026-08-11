# AI-Prompt — learning-go

> **Fundamental prompt (required first):** [[AI-Prompt/00 - General Study Guide Prompt]]  
> The General prompt is the **foundation** for every study guide — structure, Q&A format, diagrams, source handling, and quality bar. It **cannot be used alone**. Always copy it first, then append this project prompt **and** [[AI-Prompt/01 - Go - Specific Prompt]].

> **Purpose:** Project-specific layer for *Learning Go, Second Edition* (Jon Bodner, 2024) as the **primary idiomatic guide**, with *100 Go Mistakes and How to Avoid Them* (Teiva Harsanyi, 2022) as the essential **comparison for pitfalls and anti-patterns**, and *Go in Practice, Second Edition* (Kozyra, Butcher, Farina, 2025) as the **practice-patterns supplement** for real-world recipes.
>
> **Primary book:** Learning Go (Bodner) drives chapter structure, philosophy explanations, correct idioms, "why Go works this way", and practical examples.
> **Comparison book:** 100 Go Mistakes (Harsanyi) supplies categorized concrete mistakes, the "what not to do", memory/concurrency gotchas, testing traps, and production war stories.
> **Practice supplement:** Go in Practice (2nd ed.) adds hands-on patterns — CLI/config, HTTP servers, testing/benchmarking, gRPC, cloud/microservices, reflection/codegen — as a fourth comparison column, not a replacement for Bodner's narrative.
> **Third source:** Internet / official Go docs, Go blog, release notes (Go 1.22–2026+), Effective Go updates, and current community consensus — not a fourth book.

---

## Source directories

| What | Directory | Contents |
|------|-----------|----------|
| **This prompt** | `Go/Learning-Go/` | `AI-Prompt - learning-go.md` (this file) |
| **Study guides & chapter notes (output)** | `Go/Learning-Go/` | `{N}. {Title} - Study Guide.md`; raw notes or old versions as `{N}.OLD-{suffix} - {Title}.md` |
| **Book PDFs** | Repo root (`Book/`, sibling of `Obsidian/`) | `Learning Go 2024 494pages.pdf`, `100-Go-Mistakes-and-How-to-Avoid-Them-(Teiva-Harsanyi).pdf`, `Go in Practice, Second Edition.pdf` |
| **Related Go notes** | `Go/` (various subfolders) | `Concurrency-in-Go/`, `HTTP/`, `gRPC/`, `Testing/`, existing partial notes under `Learning-Go/` |
| **Companion / examples** | (book examples + stdlib) | Bodner examples are inline; cross-reference Go by Example, source of `testing`, `slices`, `maps`, `context`, etc. |

---

## Book roles

| Role | Book | Path | Use in study guides |
|------|------|------|---------------------|
| **Primary (Ref 1)** | *Learning Go, Second Edition* (Jon Bodner, O'Reilly 2024) | `Learning Go 2024 494pages.pdf` | Chapter organization, deep "why" explanations, idiomatic patterns, realistic code, tooling chapter, philosophy of simplicity and clarity |
| **Comparison (Ref 2)** | *100 Go Mistakes and How to Avoid Them* (Teiva Harsanyi, Manning 2022) | `100-Go-Mistakes-and-How-to-Avoid-Them-(Teiva-Harsanyi).pdf` | 100 categorized mistakes, anti-patterns, memory leaks, concurrency bugs, testing errors, project organization traps, before/after thinking |
| **Third (Ref 3)** | Internet / Go official + community (2026) | (web search enabled) | Post-2022/2024 changes (loop variable semantics, `slices`/`maps` packages, `testing/slog`, `iter`, `go test` improvements, generics maturity, performance tooling, `context` best practices) |
| **Practice supplement** | *Go in Practice, Second Edition* (Kozyra, Butcher, Farina, Manning 2025) | `Go in Practice, Second Edition.pdf` | Real-world recipes: CLI flags/config, graceful HTTP shutdown, table tests/fuzzing, JWT/auth, gRPC clients, cloud/containers, microservices comms — cite in comparison tables, not as Ref 1–3 |

> **PDF note:** The Go in Practice file on disk may be named with a line break after the comma (`Go in Practice,` + newline + `Second Edition.pdf`). Treat the repo-relative path above as canonical when filling prompt inputs.

**Set in general prompt:** `{{PRIMARY_REFERENCE}}` = `1` (Bodner / Learning Go leads the narrative and structure; Harsanyi supplies the "watch out for this" column; Go in Practice enriches with production recipes; internet supplies the 2026 column and any deprecations/fixes)

**Output folder:** `Go/Learning-Go/`

**Exemplar study guide (once created):** Match tone, rich comparison tables (Bodner | Harsanyi | Go in Practice | Current), Obsidian callouts, and balance of "correct way" vs "common mistake" vs "production recipe".

---

## Bodner chapter map (primary structure)

Use Bodner chapter order and titles for study guide filenames and part organization. Pull relevant mistakes from Harsanyi and practice patterns from Go in Practice where they overlap.

| Ch | Bodner Title (2024) | Harsanyi (100 Mistakes) | Go in Practice (2nd ed.) | Notes |
|----|---------------------|-------------------------|----------------------------|-------|
| 1 | Setting Up Your Go Environment | #16 (linters), project org early | Ch.1 Getting started with Go | Tooling, modules, `go build/fmt/vet`, workspace, env vars |
| 2 | Predeclared Types and Declarations | Data types (#17–#29) | Ch.2 CLI flags, slices/maps in CLI context | Zero values, literals, `var` vs `:=`, const |
| 3 | Composite Types | #20–#29 (slices, maps) | Ch.3 Structs, interfaces, generics, JSON tags | Arrays, slices, strings/runes, maps, structs |
| 4 | Blocks, Shadows, and Control Structures | #1 shadowing, #30–#35 range | — (partial via Ch.2 enums/flags) | Blocks, shadowing, if/for/switch |
| 5 | Functions | #42–#47 (defer, receivers) | Ch.3 methods on structs, generics constraints | Closures, defer, variadic, call by value |
| 6 | Pointers | #10, #26 leaks, range pointers | — | nil, GC, when pointers help/hurt |
| 7 | Types, Methods, and Interfaces | #5–#10 (interfaces, embedding) | Ch.3 interfaces, generics, JSON encoding | Embedding, producer-side interfaces, `any` |
| 8 | Generics | #9 (when to use generics) | Ch.3 generics constraints & type sets | Inference, `comparable`, when to avoid |
| 9 | Errors | Error management (~#48–#56) | Ch.4 errors, wrapping, panics, recover | `%w`, `errors.Is/As`, panic vs error |
| 10 | Modules, Packages, and Imports | #12–#15 (org, collisions, docs) | Ch.6 project cleanliness, dependencies | `go.mod`, `internal/`, import cycles |
| 11 | Go Tooling | #16 (linters) | Ch.6 vet, coverage, benchmarking, logging | pprof, race detector, staticcheck |
| 12 | Concurrency in Go | Sections 8 & 9 (#57–#74) | Ch.5 goroutines, channels, mutex, panic on goroutines | `select`, `context`, `sync`, errgroup |
| 13 | The Standard Library | Section 10 (#75–#81) | Ch.7 files/TCP/UDP; Ch.8–11 HTTP/REST/gRPC | `net/http`, JSON, JWT, external APIs |
| 15 | Writing Tests | Section 11 (#82–#90) | Ch.6 table tests, fuzzing, `go cover`, benchmarks | `-race`, `httptest`, fuzz, no sleeping |
| 16 | Here Be Dragons: Reflect, Unsafe, and Cgo | Advanced / selected | Ch.13 reflection, tags, code generation | unsafe, reflect, cgo costs |

**Note on numbering:** Bodner has no Chapter 14 in the 2nd edition TOC. Go in Practice uses its own 13-chapter structure across four parts — map by topic, not by matching chapter numbers.

When Harsanyi or Go in Practice has no direct match, note "partial overlap" in the Sources table and cite the closest chapter.

---

## Go in Practice chapter map (supplementary — by topic)

Use when a Bodner chapter benefits from a production recipe or end-to-end pattern. Do not let Go in Practice drive study guide structure.

| Part | Ch | Go in Practice Title | Typical Bodner cross-ref |
|------|----|----------------------|--------------------------|
| 1 | 1 | Getting started with Go | Bodner Ch.1 |
| 1 | 2 | Building a command-line application | Bodner Ch.2–3, Ch.10 |
| 1 | 3 | Structs, interfaces, and generics | Bodner Ch.3, Ch.7–8 |
| 2 | 4 | Handling errors and panics | Bodner Ch.9 |
| 2 | 5 | Concurrency in Go | Bodner Ch.12 |
| 2 | 6 | Formatting, testing, debugging, and benchmarking | Bodner Ch.11, Ch.15 |
| 2 | 7 | File access and basic networking | Bodner Ch.13 |
| 3 | 8 | Building an HTTP server | Bodner Ch.13 (`net/http`) |
| 3 | 9 | HTML and email template patterns | Bodner Ch.13 |
| 3 | 10 | Sending and receiving data | Bodner Ch.13 |
| 3 | 11 | Working with external services (REST, gRPC) | Bodner Ch.13; links to `Go/gRPC/` notes |
| 4 | 12 | Cloud-ready applications and communications | Bodner Ch.10–11; microservices preview |
| 4 | 13 | Reflection, code generation, and advanced Go | Bodner Ch.16 |

---

## Suggested general-prompt inputs (copy & adjust per chapter)

| Field | Example (Ch.1) |
|-------|----------------|
| `{{TOPIC_DOMAIN}}` | Go fundamentals and idioms |
| `{{CHAPTER_NUM}}` | 1 |
| `{{CHAPTER_TITLE}}` | Setting Up Your Go Environment |
| `{{OUTPUT_FILENAME}}` | `1. Setting Up Your Go Environment - Study Guide.md` |
| `{{OUTPUT_FOLDER}}` | `Go/Learning-Go/` |
| `{{CURRENT_YEAR}}` | 2026 |
| `{{REFERENCE_1}}` | Learning Go (Bodner) — Ch.1 |
| `{{REFERENCE_1_PATH}}` | `Learning Go 2024 494pages.pdf` |
| `{{REFERENCE_2}}` | 100 Go Mistakes (Harsanyi) — relevant org/tooling mistakes (e.g. #12–#16) |
| `{{REFERENCE_2_PATH}}` | `100-Go-Mistakes-and-How-to-Avoid-Them-(Teiva-Harsanyi).pdf` |
| `{{REFERENCE_3}}` | Internet: official Go docs, release notes, Effective Go, golangci-lint, `go 1.23+` changes |
| `{{REFERENCE_3_PATH}}` | (web search — no local path) |
| `{{SUPPLEMENT_BOOK}}` | Go in Practice (2nd ed.) — Ch.1 Getting started with Go |
| `{{SUPPLEMENT_BOOK_PATH}}` | `Go in Practice, Second Edition.pdf` |
| `{{PRIMARY_REFERENCE}}` | `1` |
| `{{USE_INTERNET_RESEARCH}}` | yes |
| `{{NUM_PARTS}}` | 8 |
| `{{QUESTIONS_PER_PART}}` | 5 |

**Ch.3 (Composite Types) example:**

| Field | Value |
|-------|-------|
| `{{CHAPTER_NUM}}` | 3 |
| `{{CHAPTER_TITLE}}` | Composite Types |
| `{{OUTPUT_FILENAME}}` | `3. Composite Types - Study Guide.md` |
| `{{REFERENCE_1}}` | Learning Go (Bodner) — Ch.3 |
| `{{REFERENCE_2}}` | 100 Go Mistakes (Harsanyi) — Ch.3 Data types (#17–#29) |
| `{{SUPPLEMENT_BOOK}}` | Go in Practice (2nd ed.) — Ch.3 Structs, interfaces, generics |

**Ch.12 (Concurrency) example:** Reference the two dedicated Harsanyi sections (foundations + practice) plus Go in Practice Ch.5 (goroutines, channels, mutex).

**Ch.13 (Standard Library / HTTP) example:** Bodner Ch.13 + Harsanyi #75–#81 + Go in Practice Ch.8–11 (HTTP server, templates, REST/gRPC clients).

---

## When to use

- Any chapter from *Learning Go* (Bodner) paired with relevant mistakes from Harsanyi and practice patterns from *Go in Practice* (2nd ed.) where they overlap.
- Study guides saved under `Go/Learning-Go/` using the pattern `{N}. {Bodner Chapter Title} - Study Guide.md`
- Raw / exploratory notes or pre-study-guide versions use the vault OLD pattern: e.g. `3.OLD-notes - Composite Types.md`
- When regenerating a study guide: first rename the current one to `{N}.OLD-guide - {Title} - Study Guide.md` (see general prompt)
- **Prompt stack (always):** General → [[01 - Go - Specific Prompt]] → this file
- Great for building a complete "Learning Go + avoid the 100 mistakes" curriculum in Obsidian.

---

## Specific prompt (append after general + Go prompts)

```
## TOPIC-SPECIFIC INSTRUCTIONS — learning-go ({{CURRENT_YEAR}})

You are a senior Go engineer and technical educator. Apply these rules ON TOP of the general prompt AND the Go-specific prompt.

---

### Source hierarchy (CRITICAL)

1. **Learning Go (Bodner, Ref 1) is ground truth** for chapter structure, explanations of Go's design philosophy, correct idiomatic patterns, realistic code examples, and "why" a certain approach is preferred.
2. **100 Go Mistakes (Harsanyi, Ref 2)** is the dedicated "pitfall" source — use it to surface common developer errors, memory/concurrency leaks, testing traps, organizational mistakes, and the "before" version of code or thinking. It enriches every comparison with concrete warnings.
3. **Internet (Ref 3)** supplies the {{CURRENT_YEAR}} senior practice column: language/runtime changes since the books (Go 1.22 loop semantics, `slices`/`maps` packages, improved tooling, `slog`, generics refinements, `iter`, context propagation patterns, diagnostic tooling), official recommendations, and deprecation notes.

**Go in Practice (2nd ed.) — practice supplement (not Ref 1–3):** Use for hands-on recipes when a Bodner chapter touches HTTP, CLI, testing workflows, gRPC clients, cloud/containers, or reflection/codegen. Cite chapter and pattern name. Verify against PDF — do not invent section content.

**Every comparison table MUST have (at minimum) these columns:**
Learning Go (Bodner) | 100 Go Mistakes (Harsanyi) | Go in Practice (2nd ed.) | Current practice ({{CURRENT_YEAR}})

When the opening Sources table allows a fifth column for stdlib/community, label it **Stdlib / community** — separate from Ref 3.

**Framing rules:**
- Write for a developer who wants to write *clear, idiomatic, production-ready* Go and avoid the most common footguns.
- Bodner drives the positive story ("here is how and why to do it right").
- Harsanyi drives the cautionary story ("here is what goes wrong in practice and how to recognize it").
- Go in Practice drives the recipe story ("here is how teams ship CLI/HTTP/gRPC/cloud code in practice").
- Internet column explicitly calls out anything that has improved, changed, or been clarified since 2022/2024/2025.
- Do not invent content — extract real concepts, code patterns, and mistake descriptions from the PDFs.
- Cite only repo-relative paths or book chapter references. Never use absolute machine paths.

---

### Bodner voice & philosophy to preserve

- **Idiomatic Go = clarity and simplicity first.** Prefer readable code over clever code.
- Explain the *rationale* behind Go decisions (e.g. no exceptions, error values, no inheritance, composition via embedding/interfaces).
- Show realistic, slightly larger examples rather than tiny snippets.
- Highlight Go's approach to tooling (everything is a command you can script).
- End-of-chapter exercises are gold — turn key ones into hands-on tasks.
- Key recurring themes: zero values, explicit error handling, interfaces as contracts, call-by-value everywhere, composition not inheritance.

---

### Harsanyi — when and how to pull from Ref 2

Pull the specific numbered mistakes that map to the current Bodner chapter. Always show:

- The mistake (what developers commonly do)
- Why it is a problem (readability, correctness, performance, maintainability)
- The better approach (often aligning with what Bodner teaches)

Important mappings to emphasize:

| Bodner topic                  | Key Harsanyi mistakes to surface                              |
|-------------------------------|---------------------------------------------------------------|
| Shadowing, blocks, control    | #1 Unintended variable shadowing, #2 nested code, #30–#35 range/control |
| Types / composites            | #17–#29 (octal, overflow, floats, slice len/cap/append/copy, nil vs empty, maps, leaks) |
| Functions & methods           | #42 receiver choice, #43–#44 named returns, #45–#47 defer & side effects |
| Pointers                      | Related leaks, pointer elements in range, when pointers hurt more than help |
| Interfaces & embedding        | #5 Interface pollution, #6 producer-side interfaces, #7 returning interfaces, #8 any, #10 embedding problems |
| Generics                      | #9 When (not) to reach for generics                           |
| Errors                        | Error handling section mistakes (custom errors, wrapping, sentinel vs typed) |
| Modules / packages            | #12 project misorg, #13 utility packages, #14 name collisions, #15 missing docs |
| Tooling & testing             | #16 linters, Section 11 (#82–#90): test categorization, -race, table tests, sleeping, benchmarks |
| Concurrency                   | Sections 8 & 9 (#57–#74): channels vs mutex, races, context, goroutine leaks, select, WaitGroup, errgroup, copying sync types |
| Stdlib                        | #75–#81 (time, JSON, SQL, http client/server, resource leaks) |
| HTTP / web apps               | GiP Ch.8–10 (routing, JWT, templates, static files, forms)    |
| External APIs / gRPC          | GiP Ch.11 (REST client, gRPC)                                 |
| Cloud / microservices         | GiP Ch.12 (containers, service comms, runtime monitoring)      |
| Reflection / codegen          | GiP Ch.13 (reflection, struct tags, code generation)          |

When a mistake does not perfectly align, note "related pattern" or "see also Harsanyi Ch.X #NN" or "Go in Practice Ch.X".

### Go in Practice — when to pull from the supplement

Use for **applied patterns** where Bodner teaches mechanics and Harsanyi warns of pitfalls:

| Bodner topic | Go in Practice enrichment |
|--------------|---------------------------|
| Ch.1 tooling / setup | Ch.1 install, workspace, hello world |
| Ch.3 composites | Ch.3 struct tags, JSON encoding, generics in APIs |
| Ch.9 errors | Ch.4 custom errors, wrapping, panic/recover on goroutines |
| Ch.11 tooling | Ch.6 `go vet`, fuzzing, benchmarks, structured logging |
| Ch.12 concurrency | Ch.5 mutex vs channels, channel close patterns |
| Ch.13 stdlib (`net/http`) | Ch.8–10 HTTP servers, graceful shutdown, JWT, templates |
| Ch.15 testing | Ch.6 table-driven tests, fuzz, coverage |
| Ch.16 reflect | Ch.13 reflection, tags, codegen |

Show **Bodner (idiom)** vs **Go in Practice (recipe)** vs **Harsanyi (pitfall)** in code blocks when all three apply.

---

### Content requirements (in addition to Go prompt)

#### Code blocks (mandatory — show contrast)
For every major concept or mistake:
- **Learning Go (Bodner)** — the recommended/idiomatic version from the primary book
- **Common mistake (Harsanyi or derived)** — the anti-pattern version that causes the problem
- **Go in Practice (2nd ed.)** — production recipe or end-to-end pattern when the chapter topic overlaps (CLI, HTTP, testing, gRPC, cloud, reflection); omit this block when no GiP chapter applies
- **{{CURRENT_YEAR}} senior standard** — modern refinement (if different), using latest packages (`slices`, `maps`, `cmp`, `slog`, `iter`, etc.)

Label clearly and show file context when possible.

Example skeleton:
```go
// Learning Go (Bodner) — recommended
func (s *Store) Get(ctx context.Context, id string) (Item, error) { ... }

// Common mistake (Harsanyi #XX)
func (s *Store) Get(id string) Item { ... }  // no context, swallows errors, etc.

// Go in Practice (2nd ed.) — production recipe (e.g. Ch.8 HTTP handler pattern)
func (s *Store) GetHandler(w http.ResponseWriter, r *http.Request) { ... }

// 2026 senior standard
func (s *Store) Get(ctx context.Context, id string) (Item, error) {
    // + tracing, timeout, structured logging
}
```

#### Directory trees (when chapter discusses layout)
- Ch.1 / Ch.10: module layout, `cmd/`, `internal/`, `pkg/`
- At least two trees per relevant chapter

#### Comparison tables (mandatory)
- Opening "Key New Concepts" and "Comparison at a Glance" must use **four columns:** Bodner | Harsanyi | Go in Practice | {{CURRENT_YEAR}}.
- Use "—" or "N/A" in the Go in Practice column when no recipe applies; never omit the column.
- Revisit the same topics in Part 8 senior recommendations with the same four-column layout.

#### {{CURRENT_YEAR}} Go standards (apply especially Parts 6–8)
- Modern loop semantics (variables per iteration since Go 1.22)
- Use of `slices`, `maps`, `cmp` packages where they simplify code
- Structured logging with `slog`
- Context everywhere for cancellation/trace
- Table-driven tests + `t.Run`, fuzzing, `testing/slogtest`
- Proper use of `go test -race -shuffle=on -count=1`
- `golangci-lint`, `staticcheck`, `govulncheck`
- Generics only when they reduce real duplication without harming readability
- Error wrapping with `%w` + `errors.Is`/`errors.As`
- Resource cleanup with `defer` (correct argument evaluation)

---

### Part arc (adapt titles to the Bodner chapter)

| Part | Focus |
|------|-------|
| 1 | Bodner core theory + chapter goals + philosophy |
| 2 | Bodner correct implementation mechanics and examples |
| 3 | Harsanyi pitfalls — what goes wrong, why, and how to detect |
| 4 | Deeper mechanics (memory model, GC interaction, runtime) when relevant |
| 5 | Edge cases, errors, and operational concerns |
| 6 | Combined best practices (Bodner + Harsanyi + Go in Practice recipes + {{CURRENT_YEAR}}) |
| 7 | 2026 snapshot — what has changed since the books, modern recommendations, tooling |
| 8 | Senior synthesis + end-to-end diagram (if applicable) + Master Q&A |

Rename part titles to fit the chapter (e.g. for Ch.12: "Goroutine Lifecycle", "Channel Patterns", "Common Concurrency Bugs", "Context Deep Dive", etc.).

---

### Part 7 — Modern snapshot (2026)

#### ✅ Ahead of the books / current community & stdlib
Table: Area | Status (what Bodner/Harsanyi/Go in Practice covered well vs what is now standard)

#### ⚠️ Gaps vs {{CURRENT_YEAR}} senior standards
Table: Gap | Risk | Priority

Typical areas to evaluate (verify against current reality):
- Loop variable capture
- Slice/map initialization patterns and new packages
- Error handling ergonomics
- Testing features (fuzz, coverage profiles, etc.)
- Concurrency helpers (`errgroup`, `sync.Once`, `sync.Pool`)
- Observability (slog + tracing)
- Generics usage guidelines matured since 1.18

---

### Part 8 — Diagrams (required)

Include:
- At least one mermaid diagram (sequence, state, or flowchart) for key flows in the chapter (e.g. request lifecycle, error paths, goroutine coordination).
- ASCII or directory tree(s) for data structures or project layout.
- Tables comparing correct vs mistaken approaches.

Default for concurrency-heavy chapters: lifecycle or coordination mermaid.

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks` before Further Reading.

Each task:
- **Goal** | **Files** | **Steps** | **Done when**
- Must be concrete and verifiable (run `go test`, produce output, pass race detector, etc.)
- Should demonstrate the Bodner idiom **and** include a check or test that would have caught the Harsanyi mistake; when GiP applies, mirror a recipe step (e.g. table test, graceful shutdown, JWT middleware).
- Use {{CURRENT_YEAR}} tooling and packages.

Example skeleton:
#### Task 1: Avoid Shadowing + Write Table Test
**Goal:** ...
**Files:** ...
**Steps:**
1. ...
**Done when:** `go test -race ./...` passes cleanly and the table test covers the edge case.

---

### Interview Q&A focus (weight toward)

- Why Go chose X (Bodner philosophy)
- Concrete symptoms and fixes for the mapped Harsanyi mistakes
- 2026 practicals: loop vars, `errors.Is`, table tests + race, context propagation, when to introduce generics
- "What is the zero value of ...", "Why does this leak?", "How would you test this without sleeping?"

Part 8 must contain a Master Q&A (5–8 strong senior-level questions).

---

### Obsidian note conventions (match style of gRPC / Linux study guides)

- Use the exact emoji headers defined in the general prompt.
- Collapsible answers: `> [!success]- 📋 Answers — Part N`
- Filename exactly: `{CHAPTER_NUM}. {Bodner Title} - Study Guide.md`
- Link to related existing notes when helpful (e.g. [[Go/Concurrency-in-Go/...]], [[Go/Testing/...]])
- Footer with regenerate instructions (include this in every generated guide):

### 🤖 Regenerate this guide

1. [[AI-Prompt/00 - General Study Guide Prompt]] — fill the Inputs table (Bodner = Ref 1, Harsanyi = Ref 2, Internet = Ref 3, PRIMARY = 1; supplement = Go in Practice PDF)
2. [[AI-Prompt/01 - Go - Specific Prompt]] — append Go code/tree/2026 requirements
3. [[Go/Learning-Go/AI-Prompt - learning-go]] — append this project prompt (four-column tables: Bodner | Harsanyi | Go in Practice | {{CURRENT_YEAR}})

---

### Further Reading (required)

- Next Bodner chapter(s)
- Relevant Harsanyi mistake numbers / sections
- Relevant Go in Practice chapters from the topic map (cite part + chapter title)
- Official: [Go documentation](https://go.dev/doc/), [Effective Go](https://go.dev/doc/effective_go), [Go blog](https://go.dev/blog/)
- Tooling: golangci-lint, `go test` docs, `pprof`, `trace`
- Community: Go by Example, the Go wiki on common mistakes
- Book references for readers who own the books (Bodner, Harsanyi, Kozyra/Butcher/Farina)

---

Generate the complete study guide now.
```

---

## Quick reference — Harsanyi section mapping (high level)

| Topic area                  | Harsanyi main sections          |
|-----------------------------|---------------------------------|
| Project / code org          | 2                               |
| Data types & slices/maps    | 3                               |
| Control & range             | 4                               |
| Strings                     | 5                               |
| Functions & methods         | 6                               |
| Error management            | 7 (approx)                      |
| Concurrency foundations     | 8                               |
| Concurrency practice        | 9                               |
| Standard library            | 10                              |
| Testing                     | 11                              |
| Optimizations & GC          | 12                              |

Use specific mistake numbers (#NN) in tables and text whenever possible.

---

## Regenerate workflow

1. [[AI-Prompt/00 - General Study Guide Prompt]] — fill Inputs table (Learning Go / Bodner = Ref 1, 100 Mistakes / Harsanyi = Ref 2, Internet = Ref 3; supplement PDF = Go in Practice, Second Edition)
2. [[AI-Prompt/01 - Go - Specific Prompt]] — append language rules (code blocks, trees, etc.)
3. **This file** (`Go/Learning-Go/AI-Prompt - learning-go.md`) — append the project-specific instructions (four-column comparison: Bodner | Harsanyi | Go in Practice | 2026)
4. Output the study guide to `Go/Learning-Go/{N}. {Title} - Study Guide.md`

If an older version of the guide exists, rename it using the `.OLD-guide` pattern first.
