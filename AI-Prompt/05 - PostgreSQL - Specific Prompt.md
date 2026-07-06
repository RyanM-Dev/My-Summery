# PostgreSQL — Specific Prompt

> **Purpose:** Topic layer for PostgreSQL study guides. Append AFTER [[00 - General Study Guide Prompt]].

---

## When to use

- SQL, indexing strategy, query planning
- PostgreSQL-specific features (JSONB, CTEs, window functions)
- Books: *The Art of PostgreSQL*, application-specific PostgreSQL notes

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — PostgreSQL ({{CURRENT_YEAR}})

You are a senior PostgreSQL engineer. Apply these rules ON TOP of the general prompt.

---

### PostgreSQL content requirements

#### SQL examples (mandatory)
- Include **at least 5 SQL blocks** with realistic schemas
- Use `EXPLAIN (ANALYZE, BUFFERS)` for performance chapters
- Show PostgreSQL-specific syntax (CTEs, `RETURNING`, `JSONB`, arrays) when the chapter covers them
- Note version-specific features (PG 14–17) with version callouts

#### Diagrams (mandatory)
- Index type comparison (B-tree, GIN, GiST, BRIN) — table + use-case column
- Query planner flow (sequential scan vs index scan) — ASCII or mermaid
- MVCC / vacuum concept diagram for advanced chapters

#### Compare across sources
Cross-reference book recommendations with PostgreSQL official docs and `pg_stat_*` views.

#### {{CURRENT_YEAR}} PostgreSQL standards (Part 7–8)
- Index strategy: partial indexes, expression indexes, covering indexes via `INCLUDE`
- Connection pooling: PgBouncer for app tiers
- Monitor: `pg_stat_statements`, autovacuum health, bloat
- Prefer `GENERATED` columns / proper constraints over app-only validation
- Logical replication for upgrades; physical for DR (high-level)

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks`. Each task:
- Runnable via `psql` or Docker `postgres` image
- **Goal** | **Schema** | **Steps** | **Done when**
- Show EXPLAIN output diff before/after for optimization tasks

---

### PostgreSQL interview focus (where relevant)

- MVCC and vacuum/autovacuum
- Index selection (B-tree vs GIN for JSONB/full-text)
- Join algorithms (nested loop, hash, merge)
- Isolation levels and anomalies
- `EXPLAIN` reading: cost, rows, buffers, actual time
- Partitioning strategies (range, list, hash)

---

Generate the complete study guide now.
```