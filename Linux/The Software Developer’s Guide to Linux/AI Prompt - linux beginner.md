# SDGL — Chapter Reference Notes Prompt

> **Charts and diagrams:** Use fenced `mermaid` blocks for charts, flows, architecture, hierarchies, and relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` as appropriate. Use clear labels and Obsidian-compatible syntax. Keep runnable code, commands, literal output, payloads, and calculations in their original code formats.


> **Purpose:** Turn a chapter of *The Software Developer's Guide to Linux* (SDGL) and our study session into a clear, practical Obsidian reference note.
>
> **Based on:** [[DevOps/NaNa/AI Prompt - NANA DevOps|AI Prompt - NANA DevOps]]. Follow it in full: teaching order, section pattern, visual readability, concise writing, commands, Q&A, hands-on practice, required ending, and final editing pass. This file only adds what is specific to SDGL.
>
> **Precedence:** For this folder, this prompt overrides [[AI Prompt/00 - General Study Guide Prompt]] and [[Linux/AI Prompt - Linux]]. The study process (loop, progress, README) lives in [[Linux/The Software Developer’s Guide to Linux/AI Prompt - SDGL|AI Prompt - SDGL]].

---

## 🎨 Style references

- **Structure and voice:** [[DevOps/NaNa/Artifact Management/Nexus/2. Nexus Repository Manager — Raw Repositories, Users Roles, and REST API]]. Direct explanations, practical examples, command breakdowns, clear flows.

- **Diagrams:** [[DevOps/NaNa/Artifact Management/Nexus/3. Blob]]. Useful Mermaid relationships with short surrounding prose.

- **Not a style reference:** the older `1. … - Study Guide.md` and `2. … - Study Guide.md` notes in this folder. They follow the old 8-part, three-book format. Reuse their facts, not their structure.

---

## 📚 Sources

| Role | Book | Path (under `Book/`) | Use |
|---|---|---|---|
| **Primary** | *The Software Developer's Guide to Linux* (Packt, Cohen & Sturm) | `SDGL/Packt.The.Software.Developers.Guide.to.Linux.pdf` | Chapter order, terms, commands, and examples. Ground truth |
| **Enrichment** | *Linux Bible*, 11th Ed. (2026) | `Linux BIBLE 11th 2026.pdf` | An extra command or distro difference, when it fills a real gap |
| **Enrichment** | *CompTIA Linux+ and LPIC-1*, 6th Ed. (Eckert, 2024) | `CompTIA Linux+ and LPIC-1 Guide to Linux Certification 6th ed - Eckert.pdf` | Exact meanings (permission bits, signals, systemd targets), scripting exercises |

- **Page numbers:** SDGL `PDF = printed + 31`. CompTIA 6th `PDF = printed + 16`. The Linux Bible's offset changes by chapter (+19 to +36), so find its PDF page by the section heading.

- **SDGL leads.** Bible and CompTIA content goes inline, beside the concept it supports, marked `🔭 Beyond the book` when it isn't in SDGL. No mandatory three-book comparison tables, no Bible/CompTIA "parts".

- **Disagreements:** when the books differ (`which` vs `command -v`, `service` vs `systemctl`), give the modern recommendation in one sentence and a small table only if the difference matters.

- Don't add sysadmin-only topics (Samba, printing, SELinux policy, LPIC trivia) unless the chapter covers them.

---

## 👤 Reader

A Go backend developer and junior DevOps engineer who already uses the terminal daily but wants to understand *why* things work. Skip "what is a terminal" basics; explain the mechanism under familiar commands.

Use the learner's world for examples: a Go service as the process, a Docker container as the namespaced process, a CI job as a shell script, an Ubuntu server reached over SSH.

---

## 🐧 SDGL ideas to preserve

- 🔁 **REPL cycle** (Read → Evaluate → Print → Loop) as the CLI mental model.

- 🐚 **Shell vs. command line**: don't conflate them.

- 🔎 **Command resolution**: path → builtins/aliases → `$PATH` → "command not found".

- 📖 **Man page conventions**: `[optional]`, `...`, `-f` vs. `--flag`.

- ⚙️ **Process attributes**: PID, PPID, EUID/EGID, environment, working directory.

- 🐳 **Developer bridges**: container PIDs (Ch.2), Git (Ch.14), Docker (Ch.15), HTTP (Ch.17).

- 🧩 **POSIX awareness**: prefer portable forms (`command -v`) where SDGL does.

---

## 🗺️ Chapter map

| Ch | SDGL title | Pages | Useful diagram | Enrichment, if needed |
|---|---|---|---|---|
| 1 | How the Command Line Works | 1–20 | REPL / shell evaluation flow | Bible Ch.3; CompTIA Ch.2 basic usage (p. 51) |
| 2 | Working with Processes | 21–36 | Parent/child tree, signal flow, host vs. container PIDs | Bible Ch.6; CompTIA Ch.9 (p. 361) |
| 3 | Service Management with systemd | 37–44 | Unit states, `systemctl` lifecycle | Bible Ch.15; CompTIA Ch.8 systemd (p. 338) |
| 4 | Using Shell History | 45–49 | History file flow | Bible Ch.3 |
| 5 | Introducing Files | 51–69 | FHS tree, symlink vs. hard link (inode) | Bible Ch.4; CompTIA Ch.4 FHS and linking (pp. 116, 125) |
| 6 | Editing Files on the Command Line | 71–80 | Vim modes state diagram | [[Linux/Scatterd Learning/Vim]] |
| 7 | Users and Groups | 81–91 | User/group model, `sudo` flow | Bible Ch.11; CompTIA Ch.10 users (p. 408) |
| 8 | Ownership and Permissions | 93–102 | Permission bits → octal | CompTIA Ch.4 permissions, ACLs (pp. 129–144) |
| 9 | Managing Installed Software | 103–113 | Package manager flow; `apt` vs. `dnf` table | Bible Ch.10; CompTIA Ch.11 RPM and DPM (pp. 459–471) |
| 10 | Configuring Software | 115–125 | Config precedence (`/etc`, env, flags) | CompTIA Ch.7 environment files (p. 279) |
| 11 | Pipes and Redirection | 127–149 | stdin/stdout/stderr data flow | Bible Ch.3; CompTIA Ch.7 redirection and pipes (pp. 259–272) |
| 12 | Automating Tasks with Shell Scripts | 151–166 | Script execution flow, exit codes | Bible Ch.7; CompTIA Ch.7 shell scripts (pp. 280–299) |
| 13 | Secure Remote Access with SSH | 167–184 | SSH key auth sequence | [[Linux/Scatterd Learning/SSH]]; CompTIA Ch.12 SSH (p. 511) |
| 14 | Version Control with Git | 185–201 | Working tree → index → repo | Vault Git notes |
| 15 | Containerizing Applications with Docker | 203–221 | Image → container, namespaces | [[DevOps/NaNa/Docker/1. Container]]; Bible Ch.27 namespaces; CompTIA Ch.13 containers (p. 554) |
| 16 | Monitoring Application Logs | 223–233 | Log path: app → journald/file → reader | Bible Ch.13 logging; CompTIA Ch.10 logs (pp. 399–408) |
| 17 | Load Balancing and HTTP | 235–251 | Client → LB → backends sequence | Bible Ch.17 |

CompTIA pages above are printed pages of the 6th edition.

Diagrams are suggestions: draw one only when it makes the concept easier, as in the Nana prompt.

---

## 💻 Linux-specific rules

- **Commands:** say what we want to do, show the command, explain the flags that matter, and say what to expect. Show sample output only when reading it teaches something, and label it illustrative.

- **Distro:** default to Ubuntu/Debian. Show the Fedora/RHEL form only when it differs and matters (package managers, firewall).

- **Modern defaults:** `systemd` and `journalctl`, SSH keys with `ssh-agent`, `ip` over `ifconfig`. Mention `rg`/`fd` once as optional; still teach `grep`/`find`.

- **Safety:** practice runs in a VM, container, or a scratch directory such as `~/sdgl-lab`. Add a `> ⚠️` line beside `rm -r`, `kill -9`, `chmod -R`, `sudo` edits to `/etc`, and SSH config changes, with how to undo.

- **Q&A focus:** "command not found", permission denied, zombie processes, SIGTERM vs. SIGKILL, soft vs. hard links, `su` vs. `sudo`, `>` vs. `>>` vs. `2>&1`. One "the exam would ask" question per note is fine when CompTIA adds something useful.

---

## 📝 Note conventions

- **Filename:** `{N}. {SDGL Chapter Title}.md`. Archive an older note as `{N}.OLD-{Title}.md` before replacing it.

- **Header:** `[[Linux/The Software Developer’s Guide to Linux/README|🐧 SDGL]] · Ch.{N} · pp. {printed range}`, then one or two sentences on what the chapter answers.

- **Collapsed answers:** `> [!answer]- 📋 Answers`, as in the Nana prompt.

- **Related Topics:** previous/next chapter plus existing notes such as [[Linux/Scatterd Learning/SSH]], [[Linux/Scatterd Learning/Tmux]], [[Linux/Scatterd Learning/Package Manager]], and the `Linux/Abedini/` command notes.

- **Size:** scale to the chapter, usually 150–350 lines. Ch.4 should stay small.

Output only the finished Markdown, without a preamble or a report of changes.

---

## Input

**Book:** The Software Developer's Guide to Linux

**Chapter:**
`<N. CHAPTER TITLE (printed pages)>`

**Session notes / learner's explanations and confusions:**

```text
<PASTE NOTES OR SUMMARY OF THE STUDY SESSION HERE>
```
