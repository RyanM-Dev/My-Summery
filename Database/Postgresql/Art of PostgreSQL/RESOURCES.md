# The Art of PostgreSQL Course Resources

## Knowledge

- [The Art of PostgreSQL — official contents](https://theartofpostgresql.com/book/contents/)
  Primary map for the 2026 edition's part and chapter sequence. Use it to align lessons with the book; the local vault does not contain the book text.
- [The Art of PostgreSQL — companion datasets and queries](https://github.com/dimitri/TheArtOfPostgreSQL)
  Author-maintained datasets and book query examples. Use these when a lesson needs to connect SQL techniques to the book's real datasets.
- [PostgreSQL 18 documentation](https://www.postgresql.org/docs/18/)
  Version-pinned reference for SQL behavior, PostgreSQL features, and server operations. Check the target server major version whenever a lesson discusses newer syntax.
- [PostgreSQL tutorial: SQL language](https://www.postgresql.org/docs/18/tutorial-sql.html)
  Official hands-on introduction to relations, queries, joins, aggregates, updates, and deletes. Use for foundational lessons and runnable practice.
- [PostgreSQL: queries](https://www.postgresql.org/docs/18/queries.html)
  Authoritative reference for SELECT, joins, grouping, set operations, sorting, limits, CTEs, and window functions.
- [PostgreSQL: data definition and constraints](https://www.postgresql.org/docs/18/ddl.html)
  Official guide to schemas, data types, constraints, privileges, and row security. Use when teaching durable data rules and schema design.
- [PostgreSQL: data types](https://www.postgresql.org/docs/18/datatype.html)
  Version-pinned semantics for built-in types, ranges, arrays, JSON, and custom types.
- [PostgreSQL: indexes](https://www.postgresql.org/docs/18/indexes.html)
  Primary source for B-tree, Hash, GiST, SP-GiST, GIN, BRIN, multicolumn, expression, partial, and covering indexes.
- [PostgreSQL: EXPLAIN](https://www.postgresql.org/docs/18/sql-explain.html)
  Reference for plan output and `EXPLAIN ANALYZE` caveats; use when lessons teach performance diagnosis.
- [PostgreSQL: concurrency control](https://www.postgresql.org/docs/18/mvcc.html)
  Primary reference for MVCC, isolation levels, locks, deadlocks, and serialization failures.
- [PostgreSQL: constraints](https://www.postgresql.org/docs/18/ddl-constraints.html)
  Practical rules for check, unique, primary-key, foreign-key, exclusion, and not-null constraints.
- [PostgreSQL: `pg_stat_statements`](https://www.postgresql.org/docs/18/pgstatstatements.html)
  Official operational guide to aggregate planning and execution statistics for SQL statements.
- [pgvector](https://github.com/pgvector/pgvector)
  Maintainer documentation for vector types, distance operators, exact and approximate indexes, and version-compatible installation.

## Wisdom (Communities)

- [PostgreSQL mailing lists](https://www.postgresql.org/list/)
  Public technical discussion with PostgreSQL contributors and users. Use for nuanced behavior questions after checking the versioned documentation and preparing a minimal example.

## Gaps

- The local workspace has chapter titles and concise chapter notes, but not the purchased book text or full chapter examples. Lessons will teach the listed topics independently and link to the official source material; they will not claim to summarize unseen passages.
- Extension-specific lessons should add the extension maintainer's documentation before teaching installation or version-sensitive behavior.
