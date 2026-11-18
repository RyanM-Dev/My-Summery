# Linux Handbook — Maintenance and Note Prompt

> **Purpose:** Keep **one unified Linux reference** in `Linux/Linux Handbook/`. It merges every Linux concept the learner has studied, from any source (SDGL, the Nana course, Server Administration, the books), into numbered notes ordered from beginner to advanced.
>
> **Style:** Follow [[DevOps/NaNa/AI Prompt - NANA DevOps|AI Prompt - NANA DevOps]] in full: direct explanations, examples, Mermaid diagrams, grouped bullets, meaningful emojis, generous spacing, collapsed Q&A, and the required ending.
>
> **Precedence:** For this folder, this prompt (structure and maintenance) and the Nana prompt (style) override [[AI Prompt/00 - General Study Guide Prompt]] and [[Linux/AI Prompt - Linux]].

---

## 🎯 What the Handbook Is

- 📘 **The one place to learn and review Linux.** Course and book notes stay where they are as records of each source. The Handbook is the merged, cleaned-up version.

- ✅ **Only studied concepts get content.** A topic the learner hasn't studied yet stays in the index as 📖 *planned*, with no note and no content.

- 🧭 **Beginner → advanced.** The order is fixed in [[Linux/Linux Handbook/0. Linux Handbook|the index]]. Every note builds only on the notes before it.

---

## 🗂️ Fixed Curriculum

Numbers never change. A new topic goes into the existing note it belongs to, as a new section, or into its reserved number. Never renumber notes.

| # | Note | Part |
|---|---|---|
| 1 | The Shell and How Commands Run | 🐚 Shell basics |
| 2 | Files and Directories | 🐚 Shell basics |
| 3 | Editing Files with Vim | 🐚 Shell basics |
| 4 | Streams, Pipes, and Redirection | 🐚 Shell basics |
| 5 | Text Processing | 🐚 Shell basics |
| 6 | Environment Variables and Shell Startup | 🐚 Shell basics |
| 7 | Shell Scripting | 🐚 Shell basics |
| 8 | Users, Groups, and sudo | 👥 Access |
| 9 | Ownership and Permissions | 👥 Access |
| 10 | Processes and Signals | ⚙️ Running software |
| 11 | systemd, Services, and Timers | ⚙️ Running software |
| 12 | Software and Packages | ⚙️ Running software |
| 13 | Logs | ⚙️ Running software |
| 14 | IP Addressing and Routing | 🌐 Network and remote access |
| 15 | Ports, Sockets, and DNS | 🌐 Network and remote access |
| 16 | SSH Keys and Hardening | 🌐 Network and remote access |
| 17 | Firewalls and Network Exposure | 🌐 Network and remote access |
| 18 | HTTP and Load Balancing | 🌐 Network and remote access |
| 19 | Disks, Filesystems, and Mounts | 🖥️ Server operations |
| 20 | Performance and Resources | 🖥️ Server operations |
| 21 | Containers Under the Hood | 🖥️ Server operations |
| 22 | Security Hardening | 🖥️ Server operations |
| 23 | Auditing a Server | 🖥️ Server operations |

Where a concept goes when it's unclear:

- 🔤 `cut`, `sort`, `uniq`, `awk`, `sed`, `xargs`, regex → **5**. Streams, pipes and `grep` stay in **4**.
- 📜 Shell history, aliases, startup files, config layers (SDGL Ch.10) → **6**.
- 📜 `journalctl` basics stay in **11** until **13 Logs** exists; then move the log-reading parts to 13 and link back.
- 🔗 Links, inodes, `find`, `rsync` → **2**. Mounts, `df`, `du`, LVM → **19**.
- 🧱 cgroups and namespaces → **21**, with a short pointer from **10**.

---

## 📝 Note Template

**Filename:** `{N}. {Title}.md`, using the titles in the table above.

```markdown
# {emoji} {Title}

[[Linux/Linux Handbook/0. Linux Handbook|🐧 Linux Handbook]] · Part {n}: {part name} · ⬅️ [[{N-1}. …]] · ➡️ [[{N+1}. …]]

{One or two sentences: what this note answers.}

> [!toc]- 📑 Contents
> ...

...merged sections, beginner → advanced...

### 🧪 Interview Q&A
## 🔨 Hands-On Practice
### 📋 Quick Reference
### 🧠 Things to Remember
### 💡 Pro Tips
### 🔗 Related Topics
## 📚 References
```

- ⬅️ ➡️ **Navigation** links to the previous and next **existing** notes. Skip planned numbers that have no file yet.

- 🔜 **Later topics:** at most one line per section, such as "🔜 Hard links and inodes: SDGL Ch.5, see [[Linux/Linux Handbook/0. Linux Handbook|planned]]". Never teach a topic the learner hasn't studied.

### 📚 References table

Every note ends with this table. It lists **every** source for the note's concepts, studied or not.

```markdown
| Source | Where | Status |
|---|---|---|
| 📝 Nana: [[DevOps/NaNa/Linux/4. Users and Groups\|Users and Groups]] | whole note | ✅ studied |
| 📕 SDGL Ch.7 | pp. 88–90 | 📖 planned · 🔁 mostly review: the UID model is already here |
| 📕 SDGL Ch.11 | pp. 127–149 | 📖 planned · 🔁 pipes are here · 🆕 `cut`, `awk`, `sed` → note 5 |
| 📗 CompTIA 6th Ch.10 | pp. 408–419 [PDF 424–435] | 🔭 optional |
| 📘 Linux Bible Ch.11 | pp. 247–259 [PDF 278–290] | 🔭 optional |
```

- **Status:** ✅ studied · 📖 planned in the [[Linux/The Software Developer’s Guide to Linux/README|SDGL README]] reading plan · 🔭 optional extra reading · ⏭️ skipped on purpose.

- **Every 📖 planned row carries a hint** saying what reading it will change:
    - 🔁 what the note **already covers** from another source (or "mostly review" when that's nearly all of it).
    - 🆕 the **new** material, and → **where it will go**: "here" or another note number, which may still be planned.
    - When the status becomes ✅, drop the hint: the content is in the note now.

- **Pages:** printed pages. Add the PDF page in brackets for CompTIA (`+16`) and the Linux Bible (offset changes by chapter, so look it up by heading).

---

## 🧬 Merging Rules

- 🧩 **One explanation per concept.** Take the clearest explanation from the sources, then add the details only the other sources have. Don't stack three versions of the same idea.

- 🧹 **Silently fix contradictions.** When sources disagree, give the modern answer (`command -v` over `which`, `ss` over `netstat`, `ip` over `ifconfig`).

- 🧪 **Keep the learner's real examples.** The develop server, the `payment-iran` Go service, the `ryan`, `deploy` and `app` users, and confusions from study sessions are the best examples.

- 🔒 **The vault is public.** Keep `SERVER_IP`, `example.com` and `example.net` placeholders. Never add real addresses, hostnames or secrets.

- 📏 **Size:** usually 250–450 lines, counting the blank lines the spacing rules add. If a note grows past ~600 lines, split a section into a clearly named sub-note (`{N}a. …`) rather than renumbering.

- 🔗 **Source notes keep a pointer** to the Handbook note they were merged into (a `📘 Merged into` line under their header).

---

## 🔁 Update Workflow

**When:** every time the learner finishes studying something about Linux, from any source, and its source note is written. That includes an SDGL chapter, a Nana lesson, a Server Administration note, or a CompTIA or Bible section from the reading plan.

1. 🔎 **List the new concepts** the learner studied in this session.

2. 🗂️ **Find each concept's Handbook note** with the curriculum table above.

3. ✍️ **Merge.**
    - **Note exists:** add the new material to the right section, in learning order. Update Q&A, Quick Reference and Things to Remember only if the new material changes them.
    - **Note is still planned:** create it from the template. Pull in related material the learner already studied from other notes.

4. 📚 **Update the References table:** add the source, or change its status to ✅ studied.

5. 🧭 **Update [[Linux/Linux Handbook/0. Linux Handbook|the index]]:** the note's status, its "Merged from" sources, and the **Last updated** line.

6. 🗺️ **Update the [[Linux/The Software Developer’s Guide to Linux/README|SDGL README]] map:** turn the matching 🟡 or 📖 marks into ✅.

7. 📘 **Add the pointer** in the new source note: `📘 Merged into [[Linux/Linux Handbook/{N}. {Title}|{N}. {Title}]]`.

8. ✅ **Check:** every link resolves, Mermaid blocks render, and ⬅️ ➡️ navigation is correct on the neighbors of a newly created note.

Output only finished Markdown in the notes, without a preamble or a change report.
