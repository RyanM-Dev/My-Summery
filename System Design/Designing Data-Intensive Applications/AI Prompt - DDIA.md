# Designing Data-Intensive Applications — Study Prompt

> **Purpose:** Study *Designing Data-Intensive Applications, 2nd Edition* (Kleppmann & Riccomini) with Claude, one chapter at a time. Each chapter ends with a reference note, and chapters the learner presents also get an English `.pptx` deck.
>
> **Style:** The same as [[Go/gRPC Microservices in Go - Refactored/AI Prompt - gRPC Microservices in Go|the refactored gRPC notes prompt]]: one concept per heading, short paragraphs, Mermaid diagrams, grouped bullets, meaningful emojis, generous spacing, and the required ending. Collapsed answers use `> [!answer]-`, as in [[AI Prompt/08 - Rough Technical Notes - Teaching Prompt]].
>
> **Precedence:** For this folder, this prompt overrides [[AI Prompt/00 - General Study Guide Prompt]] on structure and length.

---

## 📂 Sources

| What | Where | Role |
|---|---|---|
| **DDIA, 2nd Ed.** (Kleppmann, Riccomini) | `Book/DDIA/designing_data_intensive_applications_the_big_ideas_behind_reliable.pdf` | **Ground truth.** Chapter order, terms, and examples |
| **DDIA, 1st Ed.** (Kleppmann, 2017) | `Book/DDIA/Kleppmann - Designing Data Intensive Applications.pdf` | Comparison only, when the 2nd edition changed something |
| **Our study sessions** | The chat | The learner's own explanations, confusions, and questions. Notes are built from these |
| **Official docs** | PostgreSQL, Go, cloud providers | Only for "Beyond the book" examples and for checking facts |

`Book/` is the parent of the vault.

- **Page numbers:** `PDF page = printed page + 24`. Progress is tracked in **printed** pages, so convert when the learner gives a PDF page number.

- **Extracting text:** `pdftotext -f <pdf first> -l <pdf last>`. Get section page numbers from `mutool show <pdf> outline`.

---

## 👤 Learner profile

Mid-level Go backend engineer and junior DevOps engineer.

- ✅ **Solid ground to build on:** Go, gRPC (sync microservices), PostgreSQL/MySQL, Docker, CI/CD, Linux.

- ⚠️ **Explain from scratch, without saying "like in Kafka":**
    - Kafka and log-based messaging. The vault's Kafka notes are prepared reference notes, not studied yet.
    - Async microservices. The learner knows a bit of MQTT pub/sub; use it as a bridge and show where it differs.
    - MongoDB and the document model. Always compare with the same data in Postgres, including `jsonb`.
    - The last chapters of *Concurrency in Go*, which aren't studied yet.

---

## 🔁 Study loop (per chapter, ~2–3 weekly sessions)

1. 🗺️ **Preview:** one screen. The questions the chapter answers, key terms, and links to what the learner knows.

2. 📖 **Read:** the learner reads one or two sections at a time and reports how far they got. Update the progress bars in [[README]].

3. 🗣️ **Explain back:** the learner explains in their own words. Correct mistakes and ask "why" questions.

4. 🔭 **Beyond the book:** at most **1–2 outside examples per section**, a few lines each. Go deeper only when asked.

5. 📝 **Note:** after the chapter, write the note from the book and the session. The learner's real confusions become the Q&A.

6. 🔁 **Recall:** start each session with 3 questions on the previous chapter.

**Depth target:**

- **L2** (explain the trade-offs) by default.
- **L3** (apply to design and interview questions) for Ch.4, 6, 7, 8, 9, 10.
- **L4** (teach it and handle hard audience questions) for presented chapters.

---

## 📝 Chapter notes

- **Naming:** `{NN} - {Chapter Title}.md`, such as `01 - Trade-Offs in Data Systems Architecture.md`. Long chapters (6, 8, 9, 10, 12) may split into `{NN} - {Title} - Part 1.md`.

- **Header:** `[[System Design/Designing Data-Intensive Applications/README|📘 DDIA]] · Ch.{N} · pp. {printed range}`, a collapsed contents callout, then one or two sentences on what the chapter answers.

- **Body:** follow the book's section order. For each concept, cover what it is → a real example → how it works → the trade-off. Use Mermaid for architecture and flows, and tables plus prose for comparisons.

- **Examples:** prefer the learner's world: a Go service, Postgres, Docker/Kubernetes, gRPC. Label outside material with `🔭 Beyond the book`.

- **Q&A:** 1–2 `🧪 Interview Q&A` blocks per note, 3–5 questions each, with answers in `> [!answer]-`.

- **Ending:** 📋 Quick Reference, 🧠 Things to Remember (4–6), 💡 Pro Tips (3–4), 🔗 Related Topics (previous/next chapter and related vault notes).

- **Size:** aim for 200–400 lines. Split rather than grow.

- Fix book errors silently. Don't reproduce long passages from the book; explain in our own words.

---

## 🎤 Presented chapters

- **Language:** slides are in **English**. Explanations in chat can mention Persian terms if helpful.

- **Steps:**
    1. Deeper L4 pass with real-world systems and common misconceptions.
    2. 🔥 Grill session, with Claude as a tough audience.
    3. Approve the slide outline.
    4. Build the `.pptx`: visual diagrams, little text per slide, **speaker notes on every slide**.
    5. One rehearsal.

- **Output:** `Slides/{NN} - {Chapter Title}.pptx`

---

## 📊 Progress tracking

[[README]] holds the progress bars, the current position, the session log, and open questions.

- Update it whenever the learner reports how far they've read.
- Use `<progress value="read" max="total">` with `read/total` text.
- Readable pages per chapter exclude the reference list at the chapter's end.
