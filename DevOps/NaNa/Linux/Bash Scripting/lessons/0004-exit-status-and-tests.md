# 04 · Check Status and Prerequisites

[[DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path|🐚 Bash Scripting Learning Path]]

Commands report whether they succeeded through an **exit status**. Bash uses `0` for success and a non-zero value for failure. A script can use that signal to stop before trying an operation that lacks a required file.

For example, `[[ -f "$config_file" && -r "$config_file" ]]` checks that a path is a regular file and readable. The `if` statement branches directly on that status:

```bash
#!/usr/bin/env bash

config_file=${1:-./release.env}

if [[ -f "$config_file" && -r "$config_file" ]]; then
    printf 'Config is ready: %s\n' "$config_file"
else
    printf 'Config is missing or unreadable: %s\n' "$config_file" >&2
    exit 1
fi
```

`>&2` sends the message to standard error. `exit 1` returns a failure status to the caller, such as a CI step. File tests such as `-f` and `-r` are documented in [Bash Conditional Expressions](https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html).

The most recent command status is available as `$?`, but the next command replaces it. Prefer testing a command directly with `if` or `&&`; save `$?` immediately only when you specifically need to inspect it later. [GNU Bash: Exit Status](https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html)

### 🧪 Check Your Understanding

**Q1:** What does exit status `0` mean?

**Q2:** What do `-f` and `-r` check?

**Q3:** Why is it usually better to branch directly on a command than to print something and then read `$?`?

> [!answer]- 📋 Answers
>
> **A1:** The command succeeded.
>
> **A2:** `-f` checks for a regular file; `-r` checks whether it is readable.
>
> **A3:** A later command replaces `$?`; `if command` branches on the intended command’s status immediately.

### 🔨 Try It

Save the script as `check-config.sh`, then create a harmless practice file and check both paths:

```bash
printf 'APP_NAME=demo\n' > release.env
bash check-config.sh release.env
bash check-config.sh missing.env
```

✅ The first run should print that the config is ready. The second should print an error and return a non-zero status. To inspect it manually, run `printf 'status=%s\n' "$?"` immediately after the script.

### 📋 Quick Reference

| Need | Syntax |
|---|---|
| Test regular file | `[[ -f "$path" ]]` |
| Test readability | `[[ -r "$path" ]]` |
| Branch on command result | `if command; then ...; fi` |
| Send message to stderr | `printf 'error\n' >&2` |
| Return failure | `exit 1` |

### 🧠 Things to Remember

- Zero means success; non-zero means failure.
- Test prerequisites before starting the operation that needs them.
- Read `$?` immediately if you need it.

### 💡 Pro Tips

- Use distinct non-zero statuses when different failure types matter to callers.
- A readable-file check does not validate the file’s contents; parsing is a separate step.

### 🔗 Related Topics

- Previous: [[0003-script-inputs|Pass inputs to a Bash script]]
- Next: [[0005-branching-with-if-and-case|Choose an action with `if` and `case`]]
- [Inputs and status quick reference](../reference/03-inputs-and-status.html)
- [GNU Bash: Conditional Expressions](https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html)
- [GNU Bash: Exit Status](https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html)

**Primary reading:** [GNU Bash: Exit Status](https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html) and [Conditional Expressions](https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html).

Ask me about any confusing output. For practice, tell me which check fails for `missing.env` and what status the script returns.
