# Kafka: The Definitive Guide — Lesson Prompt

> **Purpose:** Turn *Kafka: The Definitive Guide, 2nd Edition* into short lessons, basic to intermediate, that teach Kafka through the Order + Payment + Shipping project from Babal's *gRPC Microservices in Go*.
>
> **Style:** The same as [[Go/gRPC Microservices in Go - Refactored/AI Prompt - gRPC Microservices in Go|the refactored gRPC notes prompt]]: one concept per heading, short paragraphs, Mermaid diagrams, grouped bullets, meaningful emojis, generous spacing, and the required ending. That prompt's style, spacing, emoji, diagram, Q&A, and layout rules apply here unchanged.
>
> **Precedence:** For this folder, this prompt overrides [[Brokers/Kafka/AI Prompt - Kafka]] on structure and length. Keep that prompt's current-practice checklist: KRaft, idempotent producers, cooperative rebalancing, schemas, lag, and dead-letter topics.

---

## 📂 Sources

| What | Where | Role |
|---|---|---|
| **Kafka: The Definitive Guide, 2nd Ed.** (Shapira, Palino, Sivaram, Petty) | `Book/Kafka The Definitive Guide … Second Edition ….pdf` | **Ground truth** for concepts, configuration names, and behavior |
| **Shuiskov**, *Microservices with Go*, Ch.6 | `Book/Microservices with Go 2nd Edition Alexander_Shuiskov.pdf` | Async communication in Go microservices |
| **The codebase** | `Go/gRPC Microservices in Go - Refactored/microservices/` | Every lesson's project examples; quote real files |
| **Kafka docs, franz-go docs** | kafka.apache.org, pkg.go.dev/github.com/twmb/franz-go | Kafka 4.x changes and Go APIs |

`Book/` is the parent of the vault. Extract chapters with `pdftotext -f <first page> -l <last page>`; page ranges come from `mutool show <pdf> outline`.

### 📖 Required workflow

1. **Read the book section** the lesson covers before writing. Follow its order and terms.

2. **Translate Java to Go.** The book's code is Java. Show Go with **franz-go**, and name the Kafka config property next to the franz-go option, such as `acks` → `kgo.RequiredAcks(...)`. Don't show Java listings.

3. **Update for Kafka 4.x silently** where it's simple (`--bootstrap-server` instead of `--zookeeper`). Use a one-line `🆕` callout where the book describes something removed or changed: ZooKeeper → KRaft, removed configs, new defaults, KIP-848 rebalancing.

4. **Check facts.** Configuration defaults, CLI flags, and Go APIs must match Kafka 4.x and franz-go. Run commands and code against a real broker when the lesson relies on their output, and quote that output.

5. **Fix book errors silently** and verify calculations.

---

## 🛒 Project integration (every lesson)

- Use the project's names: topics `orders` and `shipments` (3 partitions, keyed by `order_id`), dead-letter topics `shipping.dlq` and `order.dlq`, consumer groups `shipping` and `order`, events `order.v1.OrderPaid` and `shipping.v1.ShipmentCreated`.

- Include a short **🛒 In the Project** section that connects the lesson to real code: the file path, a few lines, and a `make` target to see it (`make events`, `make lag`, `make topics`, `make poison`).

- Keep the sync/async split visible: Order → Payment is gRPC; Order → Shipping → Order is Kafka.

- When the code doesn't use a feature (transactions, Connect, Streams), say so and explain when the project would need it.

---

## 📏 Size and layout

- **One lesson = one idea group.** Aim for **180–320 lines**, including spacing. Split rather than grow.

- **Header:** `[[Brokers/Kafka/Kafka The Definitive Guide/README|📬 Kafka: The Definitive Guide]] · Book Ch.{N}`, a collapsed contents callout, then one or two sentences on what the lesson answers.

- **One** `🧪 Interview Q&A` block of 3–4 questions per lesson, with collapsed answers.

- **1–2** hands-on steps that run against `make up` or `make db-up`, each with a ✅ expected result.

- **Ending:** 📋 Quick Reference, 🧠 Things to Remember (4–6), 💡 Pro Tips (3–4), 🔗 Related Topics (previous/next lesson, the codebase README, related vault notes).

- **Naming:** `{N}. {Title}.md`, numbered in learning order.

Output only the finished Markdown.
