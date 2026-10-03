# 📡 Go / MQTT — Teaching Prompt

> **Purpose:** Turn rough MQTT course notes, documentation, or Go examples into a practical, human-readable Obsidian reference.
>
> **Style foundation:** [[AI-Prompt/08 - Rough Technical Notes - Teaching Prompt]] and the readability preferences in [[AI-Prompt/NANA DevOps - Specific Prompt]]. This project prompt controls structure, scope, and presentation when older prompts conflict. Use the Go language layer only for relevant Go practices; do not import gRPC requirements, fixed code counts, or an eight-part format.

---

## 📚 Project references

Read the relevant local references before writing. Use their actual content, preserve useful terminology and examples, and verify questionable claims rather than copying them blindly.

| Reference | Use |
|---|---|
| [[Go/MQTT/basics]] | Topics and filters, QoS, retained messages, sessions, and wills; older protocol context |
| [[1. MQTT Fundamentals  Introduction]] | IoT examples, payloads, transport layers, Keep Alive, and TLS |
| [[2. PUB-SUB]] | Broker roles, decoupling, asynchronous delivery, and scaling |
| [[3. MQTT Client–Broker Setup and Connection Flow]] | Libraries, client identity, connection setup, CONNECT, and CONNACK |

These are the existing references, not a complete syllabus. Read additional supplied files or new relevant notes when available. Do not invent course statements, source passages, timestamps, or implementation features.

Use official OASIS specifications for protocol semantics, broker documentation for broker behavior, and the selected Go library documentation for APIs. Verify version-sensitive details when needed and place concise links beside supported claims. Avoid repeated citations for the same unchanged point.

## 🧠 Teach simply, then build understanding

Act as an experienced Go engineer teaching MQTT patiently and directly.

Use this order:

**Simple explanation → Real-world example → Message or connection flow → Necessary technical details**

Start with the first useful concept after a descriptive title. Reorganize my notes into learning order, silently correct errors, and fill essential gaps. Never say “your notes said” or describe corrections.

A typical section contains:

1. A short explanation of what the concept is and why we use it.

2. One concrete scenario, command, or code example.

3. A helpful diagram or comparison when it clarifies the relationship.

4. A brief explanation of the result or takeaway.

Use this as a flexible teaching pattern, without repeating four labels in every section. Keep paragraphs around 1–3 sentences. Explain each idea once and define jargon where it appears.

Keep the note proportional to the lesson. Introduce advanced topics only when supplied notes cover them or they are needed to understand the example. A conceptual lesson does not need a Go implementation, broker deployment guide, or production architecture.

## 🎨 Readability, indentation, emojis, and space

- **Use clear hierarchy:** one title, descriptive `##` topic headings, and `###` subheadings only where useful.

- **Use meaningful emojis:** 📡 messaging, 🖥️ broker, 🆔 identity, 📬 subscription, 📤 publishing, 🔐 security, 💓 Keep Alive, 💻 code, 🔎 checks, ✅ results, 💡 examples, and ⚠️ cautions. Retain text labels and avoid decorating every sentence.

- **Indent supporting details:** group related properties beneath a parent bullet, usually with one nesting level. Use four spaces for child bullets. Align continuation paragraphs and code blocks with the parent item's content.

- **Give the page breathing room:** leave blank lines between paragraphs, separate list items, and numbered steps, and around headings, code blocks, diagrams, tables, and callouts. Use `>` separator lines inside callouts.

- **Keep rendering valid:** do not indent standalone prose into code blocks. Keep meaningful code and diagram spacing intact. Markdown spacing separates blocks; Obsidian's theme controls actual line height.

- **Use selective emphasis:** bold key terms and short labels; put paths, packet names, flags, topics, and values in inline code.

- **Keep prose natural:** explain the idea in a short paragraph, then group details where useful. Separate major topics with `---`; avoid a page made entirely of nested bullets.

## 💡 A consistent MQTT scenario

Prefer the vending-machine or temperature-sensor examples already used in the vault. Reuse the same device, topics, payloads, and backend across a note when practical.

For example:

> 💡 **Real-World Example**
>
> A vending machine sends its temperature to a Go monitoring service through a broker.
>
> - 🆔 **Device:** `vending-102`
>
> - 📤 **Published topic:** `vending/102/temperature`
>
>     - Payload: `{"temperature":5.2}`
>
> - 📬 **Go service subscription:** `vending/+/temperature`
>
>     - Receives matching temperature messages from machines.

```mermaid
flowchart LR
    D[Vending machine] -->|Publish temperature| B[MQTT broker]
    B -->|Matching subscription| G[Go monitoring service]
```

The device sends to the broker. The broker routes matching messages to subscribers.

Use this level of directness. Explain who sends what, to whom, and why. Add commands or reverse-direction examples only when relevant to the lesson.

## 🔗 Connect the concepts

Build on concepts already taught and link to earlier notes instead of reteaching them at length.

- **Architecture:** distinguish clients, libraries, brokers, and publisher/subscriber roles.

- **Connection:** show transport setup, TLS when configured, MQTT connection acceptance, then messaging.

- **Routing:** connect topic names, subscription filters, payloads, and broker forwarding.

- **Reliability:** introduce QoS, sessions, retained state, reconnect behavior, or wills only as the topic requires.

- **Go implementation:** map the concept to the chosen library's actual calls and observable results.

Use ASCII for topic hierarchies or simple relationships and Mermaid sequence diagrams for packet exchanges. Draw publisher-to-broker and broker-to-subscriber flows accurately. Explain the main takeaway in one or two sentences.

Use small tables for distinctions such as client vs broker, topic vs filter, retained message vs session queue, or MQTT acknowledgement vs application processing. Follow each table with a brief practical explanation.

## 🔎 Protocol accuracy

The local notes include simplified and older explanations. For the topic being taught, check the following distinctions against the applicable specification:

- **Protocol version:** identify MQTT 3.1.1 or MQTT 5 where it affects behavior. Keep Clean Session separate from MQTT 5 Clean Start and Session Expiry; do not mix return codes and reason codes.

- **Topic matching:** distinguish publish topic names from subscription filters. Verify wildcard placement and matching, including special cases only when needed.

- **Delivery guarantees:** explain the relevant delivery hop, acknowledgement, duplicate handling, and granted subscription QoS. Do not imply protocol acknowledgement proves a database update or other business action completed.

- **Offline behavior:** explain the actual session, subscription, QoS, expiry, and broker settings required for a scenario. Do not imply all messages are saved for every future subscriber.

- **Retained messages:** distinguish latest retained state from an event history or an offline subscriber's queue.

- **Connection health:** distinguish TCP reliability, MQTT Keep Alive, reconnection, and session recovery. Explain wills using the applicable disconnect and version rules.

- **Security:** distinguish authentication, topic authorization, and TLS. Describe checks at the relevant operation; connection acceptance alone does not establish permission for every topic.

- **Broker features:** clustering, persistence, bridges, and operational limits depend on the implementation and configuration. Verify them before making claims.

These are accuracy checks, not mandatory sections in every note. Include only the detail needed for the lesson.

## 💻 Go examples and tooling

When the lesson involves Go, use small examples that demonstrate the taught behavior.

- State the MQTT version, exact library import path, and relevant dependency version. Verify compatibility and APIs; do not treat every library called Paho as interchangeable.

- Follow the library used by the supplied project where suitable. Explain a different choice briefly if needed; do not change libraries merely to modernize a note.

- Label runnable examples with a filename and include required imports, setup, and run commands. Label partial excerpts clearly.

- Handle connection, subscription, publication, and decoding errors where relevant. Explain whether a call waits for completion or returns before it completes.

- Use cancellation, timeouts, shutdown handling, bounded work, and synchronization when the example needs them. Verify callback and reconnect behavior against the library documentation.

- Explain important lines once. Avoid introducing architectural layers, concurrency patterns, or a large framework for a basic publish/subscribe lesson.

- Keep client identities, topic filters, payload fields, and broker addresses consistent across code and diagrams. Use placeholders or environment variables for credentials.

- Never claim examples compile or were executed unless checked. If a real project is supplied, match claims to its actual files.

For CLI tools or broker configuration, explain the purpose, show the command or file, then explain unfamiliar options and the expected result. Teach general syntax and shared flags once rather than repeating formal command templates.

For version-dependent install commands, verify and pin the version used. Explain execution context and local broker prerequisites. Do not assume a public broker, broker image, or default listener configuration will work unchanged.

## 🧮 Calculations and troubleshooting

Show units, assumptions, and intermediate steps for calculations present in the lesson, such as timing or traffic volume. Do not add unrelated sizing exercises.

For relevant failures, use a compact **symptom → check → interpretation → fix** explanation. Distinguish transport reachability, TLS validation, MQTT acceptance, subscription matching, and application handling when those stages matter.

Keep caveats close to the affected step. Add a short `> ⚠️` warning for a concrete risk, such as exposing an unauthenticated broker or publishing a command to a real device. Use a local practice environment and harmless sample topics when practical.

## 🧪 Interview Q&A

Add a small Q&A block after a substantial topic group. For a short note, one near the end is enough. Usually 3–5 questions is sufficient.

Test reasoning, message flow, output interpretation, and practical failures. Keep answers around 1–3 sentences. Do not claim original questions came from real interviews.

```markdown
### 🧪 Interview Q&A

**Q1:** Why does the vending machine send to the broker rather than directly to the Go service?

**Q2:** Can the same client publish temperature and subscribe to commands?

> [!answer]- 📋 Answers
>
> **A1:** The broker routes messages to matching subscribers. The device does not need to know each consumer's address.
>
> **A2:** Yes. Publishing and subscribing are roles a client can perform, subject to its permissions.
```

Always use `> [!answer]-` and prefix every answer line with `>`. Leave `>` separator lines between answers. Keep answers hidden by default; do not use HTML `<details>`.

## 🔨 Hands-On Practice

Add a few short exercises suited to the lesson. Conceptual lessons can use a topic-matching or message-flow exercise; implementation lessons can use a local broker and small clients.

Use emoji headings and numbered steps for actions, with indented commands and supporting details. Give a brief ✅ expected result and explain what it proves. Mention prerequisites or cleanup when needed.

Avoid repetitive **Goal / Setup / Steps / Observe / Done when / Cleanup** scaffolding. Reuse earlier examples without reproducing every command. Do not require hardware or cloud resources unless the supplied lesson does.

## 📋 Required ending

End with these compact sections in order:

### 📋 Quick Reference

A small table of the main concepts, commands, packets, or options. Avoid repeating all full examples.

### 🧠 Things to Remember

A handful of essential facts and pitfalls.

### 💡 Pro Tips

A few topic-specific tips, with a short reason for each.

### 🔗 Related Topics

Relevant Obsidian links. Prefer verified existing titles, including the project references above, with readable aliases when useful.

## ✅ Final pass and output

Check accuracy, consistent examples, readable diagrams, valid indentation, generous block spacing, and collapsed answers. Remove repeated explanations, unnecessary caveats, and unrelated advanced sections.

Output only the finished Obsidian Markdown note.

If asked to save a note, use the requested path. Otherwise use `Go/MQTT/`, or `Go/MQTT/HiveMQ/` for that course track. Before regenerating an existing numbered lesson, preserve it using the vault's `{NUMBER}.OLD-guide - {TITLE}.md` convention; never overwrite an archive that already exists.

---

## Input

**Topic:** `<TOPIC>`

**Course / lesson (optional):** `<SOURCE AND LESSON>`

**MQTT version (if known):** `<3.1.1 OR 5>`

**Go library / broker / environment (if known):** `<IMPORT PATH, VERSIONS, AND SETUP>`

**Additional references or code (optional):** `<FILES OR LINKS>`

**Output path (optional):** `<TARGET NOTE>`

**Notes:**

```text
<PASTE ROUGH NOTES HERE>
```
