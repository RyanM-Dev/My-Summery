# Chapter 17 — Understanding Nulls

> [🧑‍🏫 Interactive lesson](lessons/0016-understanding-nulls.html)

> [!info] Source and scope
> Individual learning reference for *The Art of PostgreSQL*, following the publisher's 2026 chapter numbering. The local vault contains no copy of the book text; this note teaches the chapter's listed subject using the public contents and general PostgreSQL knowledge.

**Book part:** Part IV — SQL Toolbox  
**Official contents:** [The Art of PostgreSQL](https://theartofpostgresql.com/book/contents/)

## 📖 Chapter Summary

`NULL` represents missing or unknown information, so ordinary comparisons can yield `UNKNOWN` rather than true or false. This affects filters, joins, constraints, aggregates, and application decoding. Use `IS NULL` and `IS NOT NULL`; use `COALESCE` only when substituting a value is semantically correct.

## 🧭 Key Ideas

- **Problem:** Identify the application or data-management need before choosing a SQL construct or PostgreSQL feature.
- **Mechanics:** Follow how inputs become rows, how PostgreSQL applies operators/constraints, and what result or state change the database guarantees.
- **Trade-off:** Consider correctness, query shape, write cost, operational burden, and the server version together.

## 🌐 Practical Application

When applying this chapter, write down one real query or invariant from your application. State the expected rows or state transition, include boundary cases, and check the result against representative data. For performance-related behavior, inspect the plan and measure; for integrity behavior, try invalid and concurrent writes.

## 💻 Worked Example

```sql
SELECT id
FROM customers
WHERE email IS NULL;
```

Use `IS NULL` for missing values. `email = NULL` evaluates to unknown, so it does not select the row.

## 🧪 Check Your Understanding

**Q1:** What practical problem does this chapter help solve?

**Q2:** What assumption would make an example from this chapter incorrect or unsafe in production?

**Q3:** What evidence would you collect to decide whether this technique is working?

> [!answer]- 📋 Answers
> **A1:** `NULL` represents missing or unknown information, so ordinary comparisons can yield `UNKNOWN` rather than true or false. This affects filters, joins, constraints, aggregates, and application decoding. Use `IS NULL` and `IS NOT NULL`; use `COALESCE` only when substituting a value is semantically correct.
>
> **A2:** The answer depends on the data shape, workload, transaction boundaries, PostgreSQL version, and invariants. Make the assumption explicit and test the edge case that could violate it.
>
> **A3:** Check correctness with representative and boundary-case data. For performance, inspect `EXPLAIN (ANALYZE, BUFFERS)` and workload measurements; for integrity, verify PostgreSQL accepts valid cases and rejects invalid ones.

## 🔨 Hands-On Practice

**Goal:** Apply one idea from this chapter to a small, disposable schema or the book's companion dataset.

**Steps:**

1. State the expected result or invariant before writing SQL.
2. Implement the simplest query or schema rule that expresses it.
3. Try normal, boundary, and invalid inputs.
4. Record what PostgreSQL guarantees and what the application must still handle.

**Done when:** You can explain the result and its limitations without relying on an unexplained ORM behavior or a single happy-path example.

## 🧠 Things to Remember

- Start with the requirement and expected result, then choose the SQL feature.
- Make ordering, null behavior, cardinality, and transaction assumptions explicit where relevant.
- Validate the behavior with realistic data and your deployed PostgreSQL version.

## 🔗 Related Chapters

- [[16 - Group By, Having, With, Union All]]
- [[18 - Understanding Window Functions]]
- [[Art of PostgreSQL - All Chapters Reference]]

*Chapter 17 learning reference; concise companion, not a reproduction of the book.*
