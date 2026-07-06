# Linux — Specific Prompt

> **Purpose:** Generic Linux / shell / sysadmin study guides. Append AFTER [[00 - General Study Guide Prompt]].
>
> **Studying *The Software Developer's Guide to Linux*?** Use [[Linux/The Software Developer's Guide to Linux/AI-Prompt - linux-beginner]] instead — SDGL primary, Linux Bible + CompTIA for comparison. Pair with [[AI-Prompt/00 - General Study Guide Prompt]] first.

---

## When to use

- Command line, shell scripting, file & directory operations
- Processes, signals, job control, systemd
- Package managers (apt, yum, dnf)
- SSH, networking basics, permissions
- Books: *The Software Developer's Guide to Linux*, *Linux Bible*, CompTIA Linux+

---

## Specific prompt (append after general prompt)

```
## TOPIC-SPECIFIC INSTRUCTIONS — Linux ({{CURRENT_YEAR}})

You are a senior Linux engineer and technical educator. Apply these rules ON TOP of the general prompt.

---

### Linux content requirements

#### Command examples (mandatory)
- Include **at least 5 bash/shell command blocks** with realistic flags and sample output
- Show both **short form** and **long form** flags where useful (`ls -la` vs `ls --all`)
- Include **at least 1 pipeline** example (`grep | awk | sort` or similar)
- Label context: which distro if it matters (Debian vs RHEL)

#### Diagrams (mandatory)
- Process tree or signal flow (ASCII or mermaid)
- File permission / ownership diagram when relevant
- Filesystem hierarchy tree for path-related chapters

#### Compare across sources
| Topic | Book A | Book B | {{CURRENT_YEAR}} |
Use tables for: init systems, package managers, process tools (`ps` vs `top` vs `htop`).

#### {{CURRENT_YEAR}} Linux standards (Part 7–8)
- Prefer `systemd` over legacy init where applicable
- `journalctl` for logs; structured logging awareness
- Container-aware debugging (`nsenter`, cgroups v2)
- Security: least privilege, `sudo` not root, SSH keys + `sshd` hardening
- Modern tools: `ripgrep`/`fd` as faster alternatives where relevant (note when to use classic `grep`/`find` for portability/exams)

---

### Hands-on tasks (required — exactly 1–2)

Add `### 🔨 Hands-On Tasks`. Each task:
- Runnable on a Linux VM or WSL
- **Goal** | **Commands** | **Steps** | **Done when**
- Include expected output snippet to verify success

---

### Linux interview focus (where relevant)

- File permissions (octal vs symbolic, `chmod`/`chown`)
- Process states, signals (`SIGTERM` vs `SIGKILL`)
- Redirection (`>`, `>>`, `2>&1`, pipes, `tee`)
- `find` vs `locate`, `grep` patterns
- Package management differences (Debian vs RHEL family)
- systemd units, `systemctl`, `journalctl`

---

Generate the complete study guide now.
```