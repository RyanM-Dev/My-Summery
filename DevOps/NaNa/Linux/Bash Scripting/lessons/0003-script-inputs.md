# 03 · Pass Inputs to a Bash Script

[[DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path|🐚 Bash Scripting Learning Path]]

Command-line arguments make a script reusable: the script stays the same while each invocation supplies the release name or file to process. Quote a multiword value when calling the script, then quote parameters when using them inside it.

```bash
bash show-inputs.sh 'billing api' 'release bundle.tar.gz'
```

The arguments after the script name become **positional parameters**:

| Parameter | Meaning |
|---|---|
| `$0` | Script name or path used to invoke it |
| `$1` | First argument |
| `$#` | Number of arguments |
| `"$@"` | Each original argument kept as its own value |

Save this as `show-inputs.sh`:

```bash
#!/usr/bin/env bash

printf 'Script: %s\n' "$0"
printf 'Argument count: %s\n' "$#"
printf 'First value: <%s>\n' "${1-}"
printf 'Received value: <%s>\n' "$@"
```

In `"$@"`, every argument remains separate, including one that contains spaces. Avoid using unquoted `$@`, which can split arguments again. The Bash manual documents [positional parameters](https://www.gnu.org/software/bash/manual/html_node/Positional-Parameters.html) and [special parameters](https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html).

Arguments and environment variables are two different input channels. A CI job may use environment values for configuration and an argument for the particular file to process. For human prompts, `read -r` reads a line without treating backslashes specially; automation should generally provide inputs without waiting for a person at a terminal.

### 🧪 Check Your Understanding

**Q1:** What does `$#` contain?

**Q2:** Why is `"$@"` useful when forwarding arguments to another command?

**Q3:** In `bash show-inputs.sh 'billing api'`, how many arguments are passed?

> [!answer]- 📋 Answers
>
> **A1:** The number of arguments passed to the script.
>
> **A2:** It preserves each original argument as a separate value, including values with spaces.
>
> **A3:** One argument: `billing api`.

### 🔨 Try It

1. Save the listing as `show-inputs.sh`.

2. Run it with two values containing spaces:

    ```bash
    bash show-inputs.sh 'billing api' 'release bundle.tar.gz'
    ```

✅ The count should be `2`, and the output should show two received values, each on its own line. Then try invoking it with one quoted argument.

### 📋 Quick Reference

| Need | Syntax |
|---|---|
| First argument | `$1` |
| Argument count | `$#` |
| Preserve all argument boundaries | `"$@"` |
| Read an interactive line literally | `read -r value` |

### 🧠 Things to Remember

- Quote multiword values at the call site and when using them in the script.
- `"$@"` preserves argument boundaries; `$#` only counts them.

### 💡 Pro Tips

- Validate the expected argument count before performing an operation.
- Keep automated scripts non-interactive so CI jobs do not wait for input.

### 🔗 Related Topics

- Previous: [[0002-variables-and-quoting|Variables and quoting]]
- Next: [[0004-exit-status-and-tests|Check status and prerequisites]]
- [Inputs and status quick reference](../reference/03-inputs-and-status.html)
- [GNU Bash: Positional Parameters](https://www.gnu.org/software/bash/manual/html_node/Positional-Parameters.html)

**Primary reading:** [GNU Bash: Positional Parameters](https://www.gnu.org/software/bash/manual/html_node/Positional-Parameters.html) and [Special Parameters](https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html).

Ask me about any confusing output. For practice, reply with your predicted argument count before running the second invocation.
