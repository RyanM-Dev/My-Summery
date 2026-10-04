# MySQL — Specific Prompt

> **Charts and diagrams:** Use fenced `mermaid` blocks for charts, flows, architecture, hierarchies, and relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` as appropriate. Use clear labels and Obsidian-compatible syntax. Keep runnable code, commands, literal output, payloads, and calculations in their original code formats.


> **Purpose:** Topic layer for MySQL / SQL study guides. Append AFTER [[00 - General Study Guide Prompt]].

---

## When to use

- SQL fundamentals, joins, subqueries, aggregations
- Indexes, query plans, optimization
- Transactions, isolation levels
- Books: *MySQL Crash Course*, application-specific MySQL notes

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — MySQL ({{CURRENT_YEAR}})

You are a senior database engineer specializing in MySQL. Apply these rules ON TOP of the general prompt.

---

### MySQL content requirements

#### SQL examples (mandatory)
- Include **at least 5 SQL blocks** with realistic table/column names
- Show **sample data** (small inline tables) before complex queries
- Include **EXPLAIN** output discussion for optimization chapters
- Use MySQL 8.x syntax unless the source book is older (note differences)

#### Diagrams (mandatory)
- Mermaid ER diagram or table relationship diagram for join chapters
- B-tree index diagram for index chapters
- Join type visual (INNER, LEFT, RIGHT, FULL) comparison table

#### Compare across sources
Highlight MySQL vs generic SQL vs PostgreSQL differences only when relevant to the chapter.

#### {{CURRENT_YEAR}} MySQL standards (Part 7–8)
- InnoDB as default engine; understand row-level locking
- Covering indexes and index-only scans
- `EXPLAIN ANALYZE` (MySQL 8.0.18+) for real execution stats
- Charset: `utf8mb4` not `utf8`
- Connection pooling at app layer; avoid N+1 queries
- Replication awareness for read scaling (high-level)

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks`. Each task:
- Runnable in `mysql` CLI or Docker MySQL container
- **Goal** | **Schema** | **Steps** | **Done when**
- Include expected row counts or EXPLAIN key columns

---

### MySQL interview focus (where relevant)

- JOIN types and when to use each
- Index types (B-tree, composite, covering)
- ACID, isolation levels, phantom reads
- `GROUP BY` + `HAVING` vs `WHERE`
- Normalization vs denormalization trade-offs
- Slow query diagnosis workflow

---

Generate the complete study guide now.
```