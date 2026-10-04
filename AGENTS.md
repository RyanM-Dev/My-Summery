# Workspace Instructions

Read and follow `/home/ryan/.codex/RTK.md`.

## Technical notes and study guides

Before producing study notes, read the shared general and teaching prompts in `AI Prompt/`. Topic, course, and project prompts live in the directory they govern, named `AI Prompt - {topic or project}.md`. Always look in the target note's directory first, then walk through its parent directories up to the vault root to find applicable prompts. Read both the local project prompt and relevant parent topic or course prompts; the most specific prompt takes precedence when instructions conflict. Do not assume project prompts are stored centrally in `AI Prompt/`.

When creating or moving a project prompt, keep it beside that project's notes and update links to its actual location. Keep only shared prompts in `AI Prompt/`; avoid duplicate prompt copies or redirect-only prompt files. If a directory has no local prompt, use applicable parent and shared prompts rather than inventing project-specific instructions.

For rough technical notes from videos or courses, read and follow `AI Prompt/08 - Rough Technical Notes - Teaching Prompt.md`. This is the user’s saved instruction, added on 2026-10-01. For this workflow, its specific requirements take precedence over conflicting older prompt conventions (including using `> [!answer]-` for collapsed answers and adapting structure to the topic).

Teach in this order: understanding, real-world example, mechanics, technical details, advanced use. Silently correct mistakes, connect related concepts, explain calculations step by step, and explain command purpose, syntax, and examples. Include practical scenarios throughout, useful comparisons with prose, understanding-focused Q&A with collapsed answers, and safe hands-on exercises. End with Quick Reference, Things to Remember, Pro Tips, and Related Topics using Obsidian links. Return only finished Markdown when generating notes.

For Nana DevOps course notes, follow `DevOps/NaNa/AI Prompt - NANA DevOps.md` as the course-specific override. The user prefers the direct, example-led pattern of the Nexus Repository note and the useful diagrams of the Blob note, with shorter prose, fewer caveats, and compact exercises. Improve visual readability with meaningful emojis, whitespace, and indented supporting details under related bullets or numbered steps; keep prose concise and use valid Obsidian Markdown indentation. These preferences override conflicting verbosity or structural requirements in the older prompts.

Use generous vertical spacing in Nana notes: blank lines between paragraphs, separate list items, and steps, and around headings and visual blocks. Use `>` separator lines inside callouts and preserve correct list indentation.

For MQTT notes, follow `Brokers/MQTT/AI Prompt - mqtt.md`. For Go implementation examples, also read the relevant guidance in `Go/AI Prompt - Go.md`. Apply the same concise explanations, helpful diagrams, meaningful emojis, indentation, and generous spacing preferred for Nana notes. Read the relevant MQTT vault references, distinguish protocol and library versions, and keep Go code proportional to the lesson. The MQTT project prompt overrides conflicting requirements in older general and topic prompts.

## Vault organization

Group database notes under `DB/` (currently `Postgresql/`, `Mysql/`, and `MongoDB/`), search engine notes under `Search Engines/` (currently `Meilisearch/`), and broker or messaging notes under `Brokers/` (currently `MQTT/` and `Kafka/`). Place new projects in the appropriate category and keep their prompts in their own directories. MQTT notes belong under `Brokers/MQTT/`, including Go examples.

## Vault naming convention

Use spaces between words in note filenames and folder names. Keep ` - ` as a separator between lesson numbers and titles or other distinct title parts. Preserve technical names such as `go-cmp`, hidden configuration names, and `.OLD-` archive markers. When renaming files or folders, update affected Obsidian links, prompt paths, and saved application references.
