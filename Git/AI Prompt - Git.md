# Git — Specific Prompt

> **Charts and diagrams:** Use fenced `mermaid` blocks for charts, flows, architecture, hierarchies, and relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` as appropriate. Use clear labels and Obsidian-compatible syntax. Keep runnable code, commands, literal output, payloads, and calculations in their original code formats.


> **Purpose:** Topic layer for Git / version control study guides. Append AFTER [[00 - General Study Guide Prompt]].

---

## When to use

- Git fundamentals: commits, branches, merges, rebases
- Workflows: feature branches, stash, cherry-pick, reset
- Remote operations: push, pull, fetch, PR workflow
- Obsidian notes in `Git/` folder

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — Git ({{CURRENT_YEAR}})

You are a senior engineer and Git educator. Apply these rules ON TOP of the general prompt.

---

### Git content requirements

#### Command examples (mandatory)
- Include **at least 6 git command blocks** with flags explained inline
- Show **before/after** state (branch diagram or `git log --oneline` snippet)
- Include **at least 1 workflow diagram** (feature branch → PR → merge)
- Warn on destructive commands (`reset --hard`, `push --force`) with safer alternatives

#### Diagrams (mandatory)
- Commit DAG (Mermaid) for merge vs rebase chapters
- Working tree / staging / repository three-tree model
- Comparison table: merge vs rebase vs squash merge

#### {{CURRENT_YEAR}} Git standards (Part 7–8)
- Trunk-based or short-lived feature branches for most teams
- `git switch` / `git restore` over legacy `checkout` for clarity
- Signed commits where security matters
- Conventional commits for changelog automation (optional, note trade-offs)
- Never force-push shared branches; use `--force-with-lease` if unavoidable
- PR review + CI gates before merge

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks`. Each task:
- Runnable in a throwaway repo (`git init` sandbox)
- **Goal** | **Commands** | **Steps** | **Done when**
- Show `git log --graph --oneline` expected output

---

### Git interview focus (where relevant)

- Three-tree model (working, staging, repo)
- Merge vs rebase trade-offs
- `git reset` modes (soft, mixed, hard)
- Stash use cases and pitfalls
- Detached HEAD state
- Resolving merge conflicts workflow

---

Generate the complete study guide now.
```