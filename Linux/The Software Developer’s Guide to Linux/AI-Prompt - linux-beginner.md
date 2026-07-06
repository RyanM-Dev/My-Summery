# AI-Prompt — linux-beginner

> **Fundamental prompt (required first):** [[AI-Prompt/00 - General Study Guide Prompt]]  
> The General prompt is the **foundation** for every study guide — structure, Q&A format, diagrams, source handling, and quality bar. It **cannot be used alone**. Always copy it first, then append this project prompt.

> **Purpose:** Project-specific layer for studying *The Software Developer's Guide to Linux* (SDGL) as a **developer learning Linux from scratch**.
>
> **Primary book:** SDGL drives chapter structure, terminology, and examples.
> **Comparison books:** Linux Bible + CompTIA Linux+ add depth, alternate practices, and exam-style framing.

---

## Source directories

| What | Directory | Contents |
|------|-----------|----------|
| **This prompt** | `Linux/The Software Developer's Guide to Linux/` | `AI-Prompt - linux-beginner.md` (this file) |
| **Study guides & chapter notes (output)** | `Linux/The Software Developer's Guide to Linux/` | `*.md` study guides; archived notes as `{N}.OLD-{suffix} - {Title}.md` |
| **Book PDFs** | Repo root (`Book/`, sibling of `Obsidian/`) | `Packt.The.Software.Developers.Guide.to.Linux.pdf`, `Linux BIBLE 11th 2026.pdf`, `CompTIA Linux+ Guide to Linux Certification ( PDFDrive ).pdf` |
| **Related Linux notes** | `Linux/` | `SSH.md`, `Tmux.md`, `Package-Manager.md`, `Alias.md`, etc. |

---

## Book roles

| Role | Book | Path | Use in study guides |
|------|------|------|---------------------|
| **Primary (Ref 1)** | *The Software Developer's Guide to Linux* (Packt, David Cohen) | `Packt.The.Software.Developers.Guide.to.Linux.pdf` | Structure, narrative, developer angle, REPL/shell mental models |
| **Comparison (Ref 2)** | *Linux Bible* 11th ed. (2026) | `Linux BIBLE 11th 2026.pdf` | Broader admin coverage, Fedora/RHEL habits, server/desktop depth, extra commands |
| **Comparison (Ref 3)** | *CompTIA Linux+ Guide to Linux Certification* (Eckert) | `CompTIA Linux+ Guide to Linux Certification ( PDFDrive ).pdf` | Certification framing, LPIC overlap, security/hardening, troubleshooting checklists |

**Set in general prompt:** `{{PRIMARY_REFERENCE}}` = `1` (SDGL leads; Bible + CompTIA enrich comparison columns)

**Output folder:** `Linux/The Software Developer's Guide to Linux/`

---

## SDGL chapter map (use for part titles)

| Ch | SDGL Title | Bible / CompTIA cross-ref (when useful) |
|----|------------|----------------------------------------|
| 1 | How the Command Line Works | Bible Ch.3 Using the Shell; CompTIA Ch.2 shell basics |
| 2 | Working with Processes | Bible Ch.6 Managing Running Processes; CompTIA Ch.9 |
| 3 | Service Management with systemd | Bible Ch.8 init/X, Ch.15 services; CompTIA systemd units |
| 4 | Using Shell History | Bible Ch.7 shell features |
| 5 | Introducing Files (FHS) | Bible Ch.3–5 filesystems; CompTIA filesystem hierarchy |
| 6 | Editing Files (nano, vim) | Bible Ch.7; CompTIA editors |
| 7 | Users and Groups | Bible Ch.11; CompTIA user management |
| 8 | Ownership and Permissions | Bible Ch.4–5; CompTIA permissions, ACLs |
| 9 | Managing Installed Software | Bible Ch.10–11; CompTIA package management (apt/dnf) |
| 10 | Configuring Software | Bible Ch.10 admin; CompTIA config files in `/etc` |
| 11 | Pipes and Redirection | Bible Ch.7; CompTIA I/O redirection |
| 12 | Automating Tasks with Shell Scripts | Bible Ch.7 scripts; CompTIA scripting |
| 13 | Secure Remote Access with SSH | Bible Ch.12–13, Ch.26 network security; CompTIA SSH |
| 14 | Version Control with Git | Bible (light); use SDGL + [[07 - Git - Specific Prompt]] overlap only when needed |
| 15 | Containerizing with Docker | Bible Ch.27–28 clouds/containers; use [[03 - Docker - Specific Prompt]] for depth |
| 16 | Monitoring Application Logs | Bible Ch.14 troubleshooting; CompTIA logging/journal |
| 17 | Load Balancing and HTTP | Bible Ch.17 web server; developer HTTP focus from SDGL |

---

## Suggested general-prompt inputs (copy & adjust per chapter)

| Field | Example (Ch.1) |
|-------|----------------|
| `{{TOPIC_DOMAIN}}` | Linux for software developers (beginner) |
| `{{CHAPTER_NUM}}` | 1 |
| `{{CHAPTER_TITLE}}` | How the Command Line Works |
| `{{OUTPUT_FILENAME}}` | `1. How the Command Line Works - Study Guide.md` |
| `{{OUTPUT_FOLDER}}` | `Linux/The Software Developer's Guide to Linux/` |
| `{{REFERENCE_1}}` | The Software Developer's Guide to Linux — Ch.1 |
| `{{REFERENCE_1_PATH}}` | `Packt.The.Software.Developers.Guide.to.Linux.pdf` |
| `{{REFERENCE_2}}` | Linux Bible 11th ed. — Ch.3 Using the Shell |
| `{{REFERENCE_2_PATH}}` | `Linux BIBLE 11th 2026.pdf` |
| `{{REFERENCE_3}}` | CompTIA Linux+ — shell / CLI sections |
| `{{REFERENCE_3_PATH}}` | `CompTIA Linux+ Guide to Linux Certification ( PDFDrive ).pdf` |
| `{{PRIMARY_REFERENCE}}` | `1` |
| `{{USE_INTERNET_RESEARCH}}` | yes |

---

## When to use

- Any chapter from *The Software Developer's Guide to Linux*
- Beginner-to-intermediate **developer** audience (not full-time sysadmin cert prep)
- Notes saved under `Linux/The Software Developer's Guide to Linux/`
- Prior Obsidian notes in same folder can supplement `{{REFERENCE_3}}` when no third book chapter fits

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — linux-beginner / SDGL ({{CURRENT_YEAR}})

You are a senior Linux engineer teaching **software developers** who are new to Linux. Apply these rules ON TOP of the general prompt.

---

### Source hierarchy (CRITICAL)

1. **SDGL (Ref 1) is ground truth** for structure, explanations, command choices, and developer-oriented mental models (REPL, shell vs command line, FHS from a dev perspective).
2. **Linux Bible (Ref 2)** adds breadth: extra commands, distro variants (Fedora/RHEL vs Debian), server admin context, and "how pros do it differently."
3. **CompTIA Linux+ (Ref 3)** adds exam-style precision, security hardening steps, troubleshooting workflows, and official terminology — use for comparison, not as the narrative voice.

**Every comparison table MUST have columns:** SDGL | Linux Bible | CompTIA | {{CURRENT_YEAR}} practice

**Framing rules:**
- Write for a **developer** who knows programming but not Linux — connect to REPLs in Python/Ruby/JS where SDGL does.
- When sources disagree (e.g. `which` vs `command -v`, `service` vs `systemctl`), show all three + recommend SDGL's approach for daily dev work and note CompTIA answer for exams.
- Do NOT dump sysadmin-only content (Samba, print servers, LPIC trivia) unless the SDGL chapter touches the topic.
- Flag Bible content that is RHEL/Fedora-centric when SDGL is distro-neutral or Debian-leaning.

---

### SDGL voice & concepts to preserve

Mirror SDGL's teaching style where applicable:
- **REPL cycle** (Read → Evaluate → Print → Loop) for CLI chapters
- **Shell vs command line** distinction — do not conflate
- **Command resolution order:** path → builtins/aliases → `$PATH` → error
- **Man page conventions:** `[optional]`, `...`, `-flag` vs `--long-flag`
- **Developer bridges:** containers (Ch.2 namespacing), Git (Ch.14), Docker (Ch.15), HTTP (Ch.17)
- **POSIX awareness:** prefer `command -v` over non-portable alternatives when SDGL says so

---

### Content requirements

#### Command examples (mandatory — match existing note style)
- Include **at least 6 bash blocks** with:
  - The command
  - Brief comment explaining *when* a developer uses it
  - **Expected output** snippet (like SDGL/Obsidian notes)
- Include **at least 2 pipelines** for Ch.11+; at least 1 for earlier chapters when relevant
- Include a **command summary table** per part when new commands are introduced:

| Command | Short Description | Main Options/Flags | Common Use Case |

#### Diagrams (mandatory)
- **Ch.1–4:** shell evaluation flow, process tree, or REPL loop (ASCII or mermaid)
- **Ch.5–8:** FHS tree, permission octal diagram, user/group model
- **Ch.9–10:** apt vs dnf comparison table (Debian vs RHEL family)
- **Ch.11–12:** stdin/stdout/stderr redirection diagram
- **Ch.13+:** SSH key flow, container vs host PID view (Ch.2 style)

#### Three-book comparison (mandatory in Parts 1, 6, 8)
Include `#### 📊 Book Comparison at a Glance` in opening AND revisit in Part 8:

| Topic | SDGL | Linux Bible | CompTIA | {{CURRENT_YEAR}} |
|------|------|-------------|---------|------------------|
| Example | REPL mental model | BASH features list | Shell objectives | Modern defaults |

**Bible vs SDGL — typical differences to highlight:**
| Area | SDGL angle | Bible angle |
|------|------------|-------------|
| Audience | Software developers | Power user + sysadmin |
| Distro | Multi-distro, dev laptop | Often Fedora/RHEL examples |
| Depth | Just enough to be productive | Wide coverage (SELinux, Samba, AI ch.) |
| Processes | Developer + containers | Full init/runlevel history |

**CompTIA — use for:**
- Exact permission bit meanings, default file permissions
- `systemctl` subcommands, target units
- Security hardening checklists
- "What would the exam ask?" callout in 1–2 interview questions per part

#### {{CURRENT_YEAR}} developer-on-Linux standards (Part 7–8)
- `systemd` + `journalctl` (not SysV init) — note when Bible mentions legacy
- SSH keys + `ssh-agent`; disable password auth in prod
- Debian/Ubuntu (`apt`) and Fedora/RHEL (`dnf`) both shown when package managers matter
- Containers: understand host vs namespace PIDs (SDGL Ch.2)
- Modern CLI: mention `rg`/`fd` as optional speedups; still teach `grep`/`find` for exams and remote servers
- WSL2 / dev containers as valid learning environments for hands-on tasks

---

### Part arc (adapt titles to SDGL chapter)

| Part | Focus |
|------|-------|
| 1 | SDGL core theory + mental models |
| 2 | Commands and mechanics (SDGL walkthrough) |
| 3 | Linux Bible — alternate commands, extra depth, distro notes |
| 4 | CompTIA — precision, security, exam angles |
| 5 | Edge cases, troubleshooting, common beginner mistakes |
| 6 | Combined best practices (SDGL + Bible + CompTIA) |
| 7 | {{CURRENT_YEAR}} developer workflow snapshot |
| 8 | Senior recommendations + cheat sheet + Master Q&A |

Rename part titles to match chapter content (e.g. Ch.2 → "Process Hierarchy", "Signals", "ps/top", "Containers & PIDs").

---

### Part 7 — Beginner progress snapshot

#### ✅ What SDGL prepares you for
Table: Skill | SDGL chapter | Ready for |

#### ⚠️ Gaps to fill from Bible / CompTIA / practice
Table: Gap | Source to study | Priority (🔴/🟡)

---

### Part 8 — Mermaid diagram (required)

Diagram type by chapter:
- Ch.1: REPL / shell evaluation flow
- Ch.2: process parent/child tree + signals
- Ch.3: systemd unit dependencies
- Ch.5: FHS directory map
- Ch.11: pipe/redirection data flow
- Ch.13: SSH connection + key auth
- Default: chapter's end-to-end workflow

---

### Hands-on tasks (required — exactly 2)

Add `### 🔨 Hands-On Tasks` before Further Reading.

Tasks must be:
- Runnable on **Ubuntu/Debian VM, WSL2, or bare metal Linux** — no macOS-only steps unless SDGL chapter includes macOS package notes
- Beginner-safe (no destructive `rm -rf /`)
- Verifiable with command output

**Template:**
#### Task 1: {{TITLE}}
**Goal:** ...
**Commands:** ...
**Steps:**
1. ...
**Done when:** ... (show expected output)

#### Task 2: {{TITLE}}
...

**Example (Ch.1):**
1. Trace how shell resolves `ls`, `./script.sh`, and a missing command — use `type`, `which`, `command -v`, `echo $PATH`
2. REPL drill: run three commands, observe Read→Evaluate→Print for each; use `echo $?` for exit codes

---

### Interview Q&A focus (beginner developer)

Weight questions toward:
- Concepts SDGL emphasizes (REPL, PATH, fork/exec, FHS, redirection)
- "What would CompTIA ask?" — 1 question per part where relevant
- Practical scenarios: "command not found", permission denied, zombie process, SSH first connection
- Compare: `SIGTERM` vs `SIGKILL`, soft vs hard link, apt vs dnf, `su` vs `sudo`

Avoid deep SELinux policy modules or enterprise cert trivia unless the chapter covers security.

---

### Obsidian note conventions (match existing vault style)

- Use emoji headers from general prompt (`📌`, `🔧`, `💻`, `🧪`)
- Collapsible answers: `> [!success]- 📋 Answers — Part N`
- Link related vault notes when they exist: `Linux/SSH.md`, `Linux/Tmux.md`, `Linux/Package-Manager.md`
- Filename pattern: `{CH_NUM}. {SDGL Chapter Title} - Study Guide.md`
- Do not overwrite raw chapter summary notes — study guides are separate `- Study Guide.md` files

---

### Regenerate footer (include in output)

```
### 🤖 Regenerate this guide

1. [[AI-Prompt/00 - General Study Guide Prompt]] — fill Inputs (SDGL = Ref 1, PRIMARY = 1)
2. [[Linux/The Software Developer's Guide to Linux/AI-Prompt - linux-beginner]] — append this project prompt
```

---

Generate the complete study guide now.
```

---

## Quick reference — filling Ref 2 & Ref 3 by SDGL chapter

| SDGL Ch | Suggested Bible chapter | Suggested CompTIA topic |
|---------|-------------------------|-------------------------|
| 1 | Ch.3 Using the Shell | Shell, terminals, history, tab completion |
| 2 | Ch.6 Managing Running Processes | ps, top, jobs, renice, kill |
| 3 | Ch.15 Starting/Stopping Services | systemd, targets |
| 5 | Ch.3–4 Filesystems | Filesystem hierarchy, paths |
| 7–8 | Ch.11 User Accounts | Users, groups, permissions |
| 9 | Ch.10 Software | Package management |
| 11 | Ch.7 Shell | Redirection, pipes |
| 12 | Ch.7 Scripts | Shell scripting |
| 13 | Ch.26 Network Security | SSH, remote access |
| 16 | Ch.14 Troubleshooting | Logs, journald |
| 17 | Ch.17 Web Server | HTTP basics (keep developer focus) |

When a mapping is weak, set `{{REFERENCE_2}}` or `{{REFERENCE_3}}` to "best available section" and note the partial overlap in the study guide Sources table.

---

## Regenerate workflow

1. [[AI-Prompt/00 - General Study Guide Prompt]] — fill Inputs table; copy **Full general prompt** (fundamental layer)
2. **This file** — append **Specific prompt** section below
3. Save output to `Linux/The Software Developer's Guide to Linux/{chapter} - Study Guide.md`