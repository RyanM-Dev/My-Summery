# Learning Go + 100 Go Mistakes — Study Guides

This folder contains structured study guides derived from:

- **Primary:** *Learning Go, Second Edition* — Jon Bodner (2024)
- **Comparison:** *100 Go Mistakes and How to Avoid Them* — Teiva Harsanyi (2022)

## How these guides are created

Prompt stack (in order):

1. [[AI Prompt/00 - General Study Guide Prompt]]
2. [[Go/AI Prompt - Go]]
3. [[Go/Learning Go/AI Prompt - learning go]]

See the project prompt for chapter mapping, inputs examples, and regenerate instructions.

## File naming

- Study guides: `{N}. {Bodner Chapter Title} - Study Guide.md`
- Old versions (before regeneration): `{N}.OLD-... - {Title}.md`

## Current progress

- [x] 1. Setting Up Your Go Environment - Study Guide.md
- [x] 2. Predeclared Types and Declarations - Study Guide.md
- [x] 3. Composite Types - Study Guide.md
- [ ] 4. Blocks, Shadows, and Control Structures
- [ ] ... (continue with the chapter map in the project prompt)
- [ ] ...
- [ ] 12. Concurrency in Go (cross-ref existing Concurrency-in-Go/ notes)
- [ ] 15. Writing Tests (existing raw notes under 15-Writing-Tests/ can feed the guide)

## Existing raw / partial notes

These were created before the standardized study guide format:

- `Embeded File.md`
- `http net.md`
- `15-Writing-Tests/` (multiple files)

They can be used as source material when generating the corresponding full study guides.

## Regenerating

Always follow the regenerate workflow in the project prompt. Rename the existing guide to `.OLD-guide` first.
