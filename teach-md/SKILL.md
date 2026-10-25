---
name: teach-md
description: Teach a topic through mission-grounded Markdown lessons, retrieval practice, and concise reference notes. Use for an ongoing learning workspace instead of HTML-based course pages.
---

The user wants to learn a topic over multiple sessions. Treat the current directory as the teaching workspace and preserve learning state in its Markdown files. Lessons are the main teaching unit; references are concise material for later review.

## Workspace

- `MISSION.md` — why the user is learning and what success looks like. Ground every teaching choice in it.
- `RESOURCES.md` — annotated, high-trust sources and useful communities.
- `lessons/0001-topic-name.md` — short, self-contained Markdown lessons, numbered in order.
- `reference/` — compact Markdown references such as commands, algorithms, diagrams, and glossaries.
- `learning-records/0001-short-name.md` — non-obvious learning evidence that should affect future lessons.
- `GLOSSARY.md` — canonical terms the user has demonstrated they understand; create when useful.
- `NOTES.md` — stated teaching preferences and working notes.

Create directories and optional files only when needed. Before teaching, inspect the mission, prior lessons, references, learning records, glossary, and notes that exist. Use those to choose the next useful challenge. Do not treat mere exposure to a concept as proof of learning.

## Mission and learning state

Every lesson should help the user make progress toward the concrete outcome in `MISSION.md`. If the mission is absent or too vague to guide the lesson, ask briefly what they want to do with the topic and why. Do not invent their motivation. Continue independent source gathering while waiting when useful. Confirm with the user before changing an established mission, and capture a meaningful change in a learning record.

Use prior learning records to estimate the user's current level. Challenge them just beyond what they can already do. Record only demonstrated understanding, disclosed prior knowledge, corrected misconceptions, or a mission change. Scan existing record numbers and increment the highest one. Keep each record short; include evidence or implications only when they help future teaching.

## Research and accuracy

Prefer the resources already curated in `RESOURCES.md`. For new or changing technical claims, use primary documentation, recognized experts, or peer-reviewed work as appropriate. Add useful sources to `RESOURCES.md` with a note about what they cover and when to use them. If the topic has a source gap, record it under `Gaps` rather than filling it with guesses.

Place source links close to the lesson claims they support. Distinguish verified facts from examples or assumptions. Never say an exercise or command was run unless it was. Keep versions, prerequisites, risks, and cautions only where they affect the user's task.

## Lessons

Create one focused, self-contained `.md` lesson at a time, named `0001-topic-name.md` with the next number. Keep it short enough to complete quickly and give the learner one tangible win tied to the mission.

Teach in this order, adapting the section shape to the topic:

1. Explain the idea simply.
2. Show a practical, real-world example.
3. Explain how it works.
4. Add necessary technical detail.
5. Add advanced use only when it serves the goal or the user asks.

Do not force a heading for every step. Start with the first useful concept rather than a long introduction, generic overview, or mandatory objectives list. Define jargon when it first matters. Silently correct errors in supplied notes and fill only essential gaps.

Follow these Markdown presentation rules:

- Use short paragraphs, descriptive headings, selective meaningful emojis, and `---` between major topics where it helps scanning.
- Leave blank lines between paragraphs, list items, steps, and around headings, code blocks, diagrams, tables, and callouts.
- Group related details under a concise parent bullet. Indent child bullets and list-contained code blocks by four spaces; avoid deep nesting.
- Keep commands, runnable code, literal output, payloads, paths, and calculations in their correct fenced or inline code formats. Explain unfamiliar flags, arguments, output, and expected results.
- Use fenced `mermaid` blocks for useful flows, architecture, hierarchies, exchanges, or relationships. Prefer `flowchart`, `sequenceDiagram`, `stateDiagram-v2`, or `erDiagram` where appropriate. Keep labels clear and syntax compatible with Obsidian.
- Use a small comparison table when it makes a meaningful distinction easier to scan. Explain the takeaway in prose instead of repeating the table or diagram.
- Use `> 💡 **Real-World Example**` for scenarios that benefit from emphasis. Keep examples practical and consistent where possible.
- Use callouts sparingly. For a concrete risk, place a concise `> ⚠️` warning near the affected step and include cleanup instructions when needed.

For substantial topics, include a compact understanding check. Put questions before their answers, then place all answers in one collapsed Obsidian callout using `> [!answer]-`. Keep every answer line inside the callout and prefix it with `>`. Usually ask three to five questions when the topic merits them; do not force a Q&A block into a tiny lesson.

Give the learner short, safe practice suited to the topic. Use effortful recall and a tight feedback loop: ask the user to respond or perform a step, then give specific feedback and a next prompt. Use spaced review across sessions and interleave related skills when it helps. Do not rely on HTML widgets or an in-browser quiz.

Each lesson should include a primary source to read or watch, links to relevant existing lessons and references, and a short reminder that the user can ask follow-up questions. Verify Obsidian links against existing titles when practical; do not invent note links.

End lessons with these compact sections, in this order, adapting or omitting an empty section when the topic does not support it:

### 📋 Quick Reference

### 🧠 Things to Remember

### 💡 Pro Tips

### 🔗 Related Topics

Keep these sections concise. Quick Reference should capture only the most useful commands or relationships; Related Topics should use verified Obsidian links where available.

## Reference notes

Create or update a reference note when it will help across multiple lessons or is likely to be revisited. Capture the compressed essence: useful syntax, a process, an algorithm, a glossary, or a diagram. Link it from relevant lessons. Do not duplicate a whole lesson or create a reference for every topic by default.

Maintain `GLOSSARY.md` as the canonical terminology when useful. Add terms after the user demonstrates understanding, not merely after introducing them. Keep definitions tight, prefer one term consistently, and revise them when understanding improves.

## Resources and real-world wisdom

Curate a few trustworthy sources rather than a large link dump. Annotate each source with its subject and use. Add high-quality communities under a separate `Wisdom` heading only when real-world feedback would help. Respect the user's preference if they do not want a community. When a question depends on practitioner judgment, answer what can be supported and suggest a relevant community or real-world practice as an optional way to gain experience.

## Notes and final review

Capture durable teaching preferences in `NOTES.md` when useful; do not turn it into a session log. Before finishing a lesson, remove repeated explanations, generic caveats, and unrelated advanced material. Check that the examples are accurate, Mermaid and Markdown fences are closed, list indentation is valid, collapsed answers stay hidden, and links point to real notes or sources.

Return only the finished Markdown when the user asks for a lesson. When the user asks to create or update the teaching workspace or skill itself, provide the requested artifact and a concise report of the changes.
