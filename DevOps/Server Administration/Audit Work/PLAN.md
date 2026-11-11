# 🗺️ Develop Server Audit — Work Plan

Working folder for the 2026-11-11 audit of the develop/test server `SERVER_IP`.
Finished learning notes live one level up, in [[DevOps/Server Administration/0. Server Administration Learning Path|Server Administration]].

## 🎯 Goal

Ryan is new to DevOps and server administration. This audit should:

1. Find the **security** and **performance/reliability** mistakes on the server, with evidence.

2. Explain each finding: why it matters, how to fix it safely, and how to verify the fix.

3. Turn the findings into a **learning roadmap**: the concepts Ryan should learn next, at his current level (Nana DevOps course: Linux basics, Git, Docker, Jenkins, artifact management).

4. Reference the books in `/home/ryan/Media/Code/Book` by chapter, plus existing vault notes.

## 🚧 Rules for this work

- **Read-only on the server.** Run inspection commands only. Change nothing, restart nothing, install nothing. Ryan applies every fix himself.

- **Never print secrets.** Redact `.env` values, passwords, tokens, private keys and the contents of public keys (show only type and comment).

- **Every finding needs evidence:** the command that was run and what it showed.

- **Follow vault conventions:** read `AGENTS.md` and `DevOps/NaNa/AI Prompt - NANA DevOps.md`. That means concise teaching, Mermaid diagrams, emojis, generous spacing, `> [!answer]-` Q&A blocks, and Obsidian links.

## 📋 Steps

| # | Step | Output |
|---|---|---|
| 1 | Survey the vault structure, its conventions and the Linux books | conventions noted in PROGRESS |
| 2 | Security audit: accounts, SSH, sudo, network exposure, firewall, Docker, CI runner, TLS, secrets, patching | `RAW EVIDENCE.md` |
| 3 | Performance and reliability audit: CPU, memory/swap/OOM, disk, logs, containers, limits, backups, time sync | `RAW EVIDENCE.md` |
| 4 | Map book chapters (Linux Bible 11th, Software Developer's Guide to Linux, CompTIA Linux+) to each topic | references in the notes |
| 5 | Write the audit report: findings ranked by severity, then a fix order | `1. Develop Server Audit 2026-11-11.md` |
| 6 | Write concept notes, one per topic the findings point to | `2.` … `N.` notes |
| 7 | Write the learning path / roadmap index and link it from the Nana table of contents | `0. Server Administration Learning Path.md` |
| 8 | Critical self-review: re-verify claims against the evidence, check links, check book page references, check for scope creep and missing topics | `REVIEW.md`, then fixes |

## ✅ Definition of done

- Every finding in the report has evidence, a severity, a fix and a verification step.

- Every wikilink resolves to an existing note.

- Every book reference points to a real chapter (checked against the PDF's table of contents).

- `REVIEW.md` lists what was checked, what was wrong and fixed, and what remains uncertain.
