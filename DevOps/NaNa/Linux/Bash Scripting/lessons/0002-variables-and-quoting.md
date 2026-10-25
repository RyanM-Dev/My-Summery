# 02 · Pass Bash Values Safely

[[DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path|🐚 Bash Scripting Learning Path]]

Variables let one script work with different application names, paths, or release settings. Quoting an expanded value keeps spaces and wildcard characters inside one argument.

```bash
app_name='billing api'
artifact_dir='/srv/build output'

printf 'App: %s\nArtifact directory: %s\n' "$app_name" "$artifact_dir"
```

There must be no spaces around `=` in an assignment. Use `$name` to expand a value and `${name}` when braces make the variable boundary clearer:

```bash
region='eu-west'
printf '%s\n' "${region}-staging"
```

Double quotes allow expansion while keeping the result as one argument. Single quotes preserve their contents literally. Unquoted expansions can be split into multiple words and can expand wildcard patterns, which may accidentally change the arguments a command receives. [GNU Bash: Quoting](https://www.gnu.org/software/bash/manual/html_node/Quoting.html) · [ShellCheck SC2086](https://www.shellcheck.net/wiki/SC2086)

> 💡 **Real-World Example**
>
> If an artifact directory is `/srv/build output`, `"$artifact_dir"` passes one path to a command. Without quotes, the command may receive `/srv/build` and `output` as two arguments.

## Configuration From the Environment

Environment variables are useful for configuration supplied by a CI job or a parent process. A script can require an important value and choose defaults for optional ones:

```bash
app_name=${APP_NAME:?Set APP_NAME before running this script}
deploy_env=${DEPLOY_ENV:-staging}

printf 'App: %s\nEnvironment: %s\n' "$app_name" "$deploy_env"
```

- `${NAME:-default}` uses `default` when `NAME` is unset or empty.

- `${NAME:?message}` reports an error and stops expansion when the value is unset or empty.

Set an environment value for one command with `NAME=value command`. Export it from the current shell when child processes should inherit it. An ordinary, unexported shell variable stays in that shell. [GNU Bash: Shell Parameter Expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)

```bash
APP_NAME='billing api' DEPLOY_ENV=test bash release-report.sh

export APP_NAME='billing api'
bash release-report.sh
```

### 🧪 Check Your Understanding

**Q1:** Why write `"$artifact_dir"` when passing it to a command?

**Q2:** Which value does `${DEPLOY_ENV:-staging}` produce when `DEPLOY_ENV` is empty?

**Q3:** What must happen for a child process to inherit a shell variable?

> [!answer]- 📋 Answers
>
> **A1:** Quoting preserves the expanded value as one argument and prevents word splitting and wildcard expansion.
>
> **A2:** `staging`.
>
> **A3:** The variable must be exported, or supplied as an environment assignment for that command.

### 🔨 Try It

Save the report example as `release-report.sh`. Predict the output, then run it:

```bash
APP_NAME='billing api' DEPLOY_ENV=test bash release-report.sh
```

✅ The application name should remain `billing api` on one line as one value. Try leaving out `APP_NAME`; the required-value expansion should show its message.

### 📋 Quick Reference

| Need | Syntax |
|---|---|
| Assign | `name=value` |
| Expand safely as an argument | `"$name"` |
| Optional default | `${NAME:-default}` |
| Required value | `${NAME:?message}` |
| Set for one command | `NAME=value command` |
| Pass to child processes | `export NAME=value` |

### 🧠 Things to Remember

- Assignments have no spaces around `=`.
- Quote variable expansions used as command arguments.
- Child processes inherit exported environment values.

### 💡 Pro Tips

- Keep secret values out of printed reports.
- Use a clear default for optional settings and a direct error for required settings.

### 🔗 Related Topics

- Previous: [[0001-first-script|Run your first Bash script]]
- Next: [[0003-script-inputs|Pass inputs to a Bash script]]
- [[DevOps/NaNa/Linux/3. Environment Variables|Linux environment variables]]
- [Variables and quoting quick reference](../reference/02-variables-and-quoting.html)

**Primary reading:** [GNU Bash: Quoting](https://www.gnu.org/software/bash/manual/html_node/Quoting.html) and [Shell Parameter Expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html).

Ask me about any confusing output. For practice, tell me what changed when you passed the app name with a space.
