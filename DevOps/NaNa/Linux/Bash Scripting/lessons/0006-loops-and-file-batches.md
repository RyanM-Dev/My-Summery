# 06 · Loop Over a Batch of Files

[[DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path|🐚 Bash Scripting Learning Path]]

A `for` loop runs once for each item in a list. A filename glob such as `*.log` gives Bash the matching paths, and quoting the loop variable keeps a path with spaces together when it is passed to a command.

This script counts lines in log files in a practice folder:

```bash
#!/usr/bin/env bash

log_dir="$HOME/bash-labs/lesson-6"

for log_file in "$log_dir"/*.log; do
    [[ -f "$log_file" ]] || continue
    wc -l -- "$log_file"
done
```

The directory portion is quoted, but `*.log` stays unquoted so Bash can expand it. `"$log_file"` preserves the full path as one argument. `--` tells `wc` that the next item is a filename, even if it begins with a dash.

By default, if no path matches, Bash leaves the pattern unchanged. The `[[ -f ... ]]` guard then fails and `continue` skips that non-file value. The Bash manual describes [looping constructs](https://www.gnu.org/software/bash/manual/html_node/Looping-Constructs.html).

### 🧪 Check Your Understanding

**Q1:** Why should `"$log_file"` be quoted?

**Q2:** Why is `*.log` unquoted?

**Q3:** What happens when there are no matches in the default Bash configuration?

> [!answer]- 📋 Answers
>
> **A1:** Quoting keeps a filename with spaces as one argument.
>
> **A2:** Bash needs to expand the glob into matching paths.
>
> **A3:** The pattern remains literal text; the `-f` guard fails and `continue` skips it.

### 🔨 Try It

1. Create two small practice logs. These commands write only in the named practice folder:

    ```bash
    mkdir -p "$HOME/bash-labs/lesson-6"
    printf 'WARN starting\nERROR 503\n' > "$HOME/bash-labs/lesson-6/api one.log"
    printf 'Started\nReady\nListening\n' > "$HOME/bash-labs/lesson-6/worker.log"
    ```

2. Save the script above as `count-logs.sh`, then run it:

    ```bash
    bash count-logs.sh
    ```

✅ It should report `2` lines for `api one.log` and `3` for `worker.log`. Notice that the space in the first filename does not split it into two files.

### 📋 Quick Reference

| Need | Syntax |
|---|---|
| Loop over list | `for item in ...; do ...; done` |
| Expand matching paths | `"$dir"/*.log` |
| Skip current item | `continue` |
| Count lines | `wc -l -- "$file"` |

### 🧠 Things to Remember

- Leave the wildcard unquoted when you want Bash to expand it.
- Quote the resulting path when passing it to a command.
- Guard against the unmatched-pattern case before processing.

### 💡 Pro Tips

- Practice batch commands with read-only actions such as `wc` before using commands that modify files.
- Keep test files in a dedicated directory so the loop’s scope is easy to see.

### 🔗 Related Topics

- Previous: [[0005-branching-with-if-and-case|Choose an action with `if` and `case`]]
- [[DevOps/NaNa/Linux/1. Linux Pipes, `grep`, `less`, and Output Redirection|Pipes and output redirection]]
- [Decisions and loops quick reference](../reference/05-decisions-and-loops.html)
- [GNU Bash: Looping Constructs](https://www.gnu.org/software/bash/manual/html_node/Looping-Constructs.html)

**Primary reading:** [GNU Bash: Looping Constructs](https://www.gnu.org/software/bash/manual/html_node/Looping-Constructs.html).

Ask me about any confusing output. For practice, explain why the glob is only partly quoted and what protects the loop when there are no matching files.
