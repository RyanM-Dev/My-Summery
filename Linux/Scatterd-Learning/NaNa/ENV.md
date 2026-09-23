export CAPITAL_SNAKECASE=capitalsnakecase  to set env for current shell
printenv : to show all the enviroments varibale
edit in home dir, .bashrc or .zshrc and add export CAPITAL_SNAKECASE=capitalsnakecase to permanently set the ENV
source .bashrc or .zshrc : to loead new env into the currnet shell
you can append custom dir to $PATH in .bashrc or .zshrc and have scripts in those dir that can run the globally after thier path is appended to global $PATH

# 🌱 Environment Variables & Shell Configuration

Tags: #shell #bash #zsh #linux #environment-variables

---

## 1️⃣ `export` — Setting an Environment Variable

### 🔍 Explanation

An **environment variable** is a named value stored in the shell's memory that any process launched from that shell can read. `export` takes a normal shell variable and _promotes_ it so that child processes (scripts, programs, subshells) can see it too.

|Command|Scope|Visible to child processes?|Persists after shell closes?|
|---|---|---|---|
|`VAR=value`|Current shell only|❌ No|❌ No|
|`export VAR=value`|Current shell + children|✅ Yes|❌ No (session-only)|
|Added to `.bashrc`/`.zshrc` + `export`|Every new shell session|✅ Yes|✅ Yes|

The naming convention `CAPITAL_SNAKE_CASE` isn't enforced by the shell — it's a **convention** to visually distinguish env vars from regular lowercase script variables and to avoid clashing with built-in shell variables.

### 📋 Structure

```bash
export VARIABLE_NAME=value
```

- No spaces around `=`
- Quote values that contain spaces: `export NAME="value with spaces"`

### 💡 Examples

```bash
# Set a simple variable for this session only
export API_KEY=abc123xyz

# Set a variable with spaces (needs quotes)
export GREETING="Hello, World"

# Reference another variable while building a new one
export PROJECT_HOME="$HOME/projects/myapp"

# Temporarily override for a single command only (not exported globally)
API_KEY=temporary_key ./run_script.sh
```

### 🚀 Advanced Use Cases

```bash
# Export a variable computed from a command
export BUILD_DATE=$(date +%Y-%m-%d)

# Export multiple related config values for a dev environment
export NODE_ENV=development
export DEBUG=app:*
export PORT=3000

# Conditionally export based on OS
if [[ "$OSTYPE" == "darwin"* ]]; then
  export EDITOR=vim
else
  export EDITOR=nano
fi
```

### ⚠️ Important Notes

> ⚠️ Variables set with plain `export` **die when the terminal closes**. If you want it every time you open a shell, it must go into your shell's startup file (see Section 3).

> ⚠️ `export` only affects the _current shell and its children_ — it will **not** appear in a sibling terminal tab/window that was already open before you ran it.

> ⚠️ Sensitive values (API keys, tokens) exported this way are visible to any process on the system that can read `/proc/<pid>/environ` — don't treat `export` as secure secret storage.

---

## 2️⃣ `printenv` — Viewing Environment Variables

### 🔍 Explanation

`printenv` prints the current environment — i.e., all variables that have been `export`-ed and inherited by the current shell. It's related to but distinct from a few similar commands:

|Command|Shows|Notes|
|---|---|---|
|`printenv`|Only exported (environment) variables|Simple, POSIX-friendly|
|`env`|Same as `printenv`, but can also run a command with a modified environment|More flexible, `env VAR=x command`|
|`set` (bash)|**All** shell variables — exported _and_ non-exported, plus functions|Much noisier output|
|`echo $VAR`|Value of a single specific variable|Fastest for checking one var|

### 📋 Structure

```bash
printenv               # list all environment variables
printenv VAR_NAME      # show value of one specific variable
```

### 💡 Examples

```bash
# List everything
printenv

# Check a specific variable
printenv PATH

# Same thing, more common shorthand
echo $PATH

# Search environment variables for something specific
printenv | grep -i java
```

### 🚀 Advanced Use Cases

```bash
# Compare env vars between two shells/sessions by dumping to files
printenv | sort > env_before.txt
# ...make changes...
printenv | sort > env_after.txt
diff env_before.txt env_after.txt

# Use in scripts to fail fast if a required var is missing
: "${API_KEY:?Error: API_KEY is not set}"
```

### ⚠️ Important Notes

> ⚠️ `printenv VAR_NAME` returns **nothing and exit code 1** if the variable isn't set or isn't exported — this is a common silent-failure trap in scripts. Always check with `echo $?` or use the `${VAR:?}` pattern above.

> ⚠️ A variable can exist in the shell (visible via `set`) but **not** show up in `printenv`/`env` if it was never `export`-ed.

---

## 3️⃣ `.bashrc` / `.zshrc` — Persisting Environment Variables

### 🔍 Explanation

These are **startup scripts** that your shell automatically runs when a new interactive session begins. Putting `export` lines here means the variable is set up fresh every time you open a terminal — this is how you make a variable "permanent" (really: automatically re-created each session, not literally stored forever).

|File|Shell|Runs on...|
|---|---|---|
|`~/.bashrc`|Bash|Every new _interactive, non-login_ shell (most terminal tabs)|
|`~/.bash_profile` / `~/.profile`|Bash|_Login_ shells (e.g., SSH sessions, some terminal apps)|
|`~/.zshrc`|Zsh|Every new interactive Zsh shell|
|`~/.zprofile`|Zsh|Login shells (Zsh equivalent of `.bash_profile`)|

### 📋 Structure

```bash
# Open the file in an editor
nano ~/.bashrc      # or ~/.zshrc

# Add a line like this anywhere in the file
export VARIABLE_NAME=value
```

### 💡 Examples

```bash
# Inside ~/.bashrc or ~/.zshrc:

export EDITOR=vim
export LANG=en_US.UTF-8
export HISTSIZE=10000
export PYTHONDONTWRITEBYTECODE=1
```

### 🚀 Advanced Use Cases

```bash
# Load secrets from a separate, git-ignored file instead of hardcoding
# (keeps .bashrc shareable/committable without leaking keys)
if [ -f "$HOME/.env.secrets" ]; then
  source "$HOME/.env.secrets"
fi

# Set variables conditionally per machine
case "$(hostname)" in
  work-laptop) export ENV=work ;;
  home-pc)     export ENV=personal ;;
esac

# Keep config modular by splitting into separate files
for file in ~/.bashrc.d/*.sh; do
  [ -r "$file" ] && source "$file"
done
```

### ⚠️ Important Notes

> ⚠️ Editing `.bashrc`/`.zshrc` does **nothing** to already-open terminals until you either `source` the file or open a new terminal window/tab.

> ⚠️ On macOS, the default Terminal app opens **login shells**, so Bash users often need `.bash_profile` (which should then `source ~/.bashrc`) rather than `.bashrc` alone. Zsh users are less affected since `.zshrc` covers both cases by default in most setups.

> ⚠️ A typo in `.bashrc`/`.zshrc` (e.g., a broken `if` block) can break your shell startup entirely, sometimes locking you out of normal terminal use. Always test changes with `bash -n ~/.bashrc` (syntax check) or open a **new** terminal tab before closing your working one.

---

## 4️⃣ `source` — Reloading Shell Configuration

### 🔍 Explanation

`source` (or its shorthand `.`) executes a script's commands **in the current shell**, rather than in a new subshell. This is exactly what you need after editing `.bashrc`/`.zshrc`: it re-runs those `export` lines in your _existing_ terminal so you don't have to close and reopen it.

|Command|Runs in|Variables persist after script finishes?|
|---|---|---|
|`./script.sh`|New subshell|❌ No — subshell exits, changes are lost|
|`bash script.sh`|New subshell|❌ No|
|`source script.sh`|Current shell|✅ Yes|
|`. script.sh`|Current shell (POSIX shorthand for `source`)|✅ Yes|

### 📋 Structure

```bash
source ~/.bashrc
# or the POSIX-compatible shorthand:
. ~/.bashrc
```

### 💡 Examples

```bash
# After adding a new export to .bashrc
source ~/.bashrc

# Zsh equivalent
source ~/.zshrc

# Reload a specific script's variables into your current shell
source ./setup_env.sh
```

### 🚀 Advanced Use Cases

```bash
# Common pattern: activate a Python virtual environment
# (this MUST be sourced — running it normally won't affect your shell)
source venv/bin/activate

# Reload only if the file changed, inside a script
[ ~/.bashrc -nt ~/.bashrc.loaded ] && source ~/.bashrc && touch ~/.bashrc.loaded
```

### ⚠️ Important Notes

> ⚠️ This is the **#1 beginner trap**: running `./venv/bin/activate` or `./myconfig.sh` directly does **not** persist any environment changes, because it executes in a throwaway subshell. You must `source` it.

> ⚠️ If your `.bashrc` has errors, `source`-ing it will print those errors directly into your current terminal — useful for debugging, unlike opening a fresh terminal where errors might scroll by unnoticed.

---

## 5️⃣ `$PATH` — Making Scripts Globally Runnable

### 🔍 Explanation

`$PATH` is a special environment variable containing a **colon-separated list of directories**. When you type a command name (like `git` or `python`), the shell searches these directories _in order_ looking for an executable with that name. Appending your own directory to `$PATH` lets you run your own scripts from anywhere, just like built-in commands.

🔍 **Difference explained:** Running `myscript.sh` vs `./myscript.sh` vs just `myscript`:

|You type|What happens|
|---|---|
|`myscript.sh` (script not in `$PATH`, no `./`)|❌ "command not found"|
|`./myscript.sh`|✅ Runs, but only works from that exact folder|
|`myscript.sh` (script's folder **is** in `$PATH`)|✅ Runs from **any** folder, like a real command|

### 📋 Structure

```bash
# Prepend (searched first) or append (searched last) a directory to PATH
export PATH="$HOME/my-scripts:$PATH"     # prepend — recommended
export PATH="$PATH:$HOME/my-scripts"     # append
```

### 💡 Examples

```bash
# Add a personal scripts folder in ~/.bashrc or ~/.zshrc
export PATH="$HOME/scripts:$PATH"

# Add a language-specific tool directory (common example: Go)
export PATH="$PATH:$HOME/go/bin"

# Check that it worked
echo $PATH
which myscript.sh
```

### 🚀 Advanced Use Cases

```bash
# Make a script executable AND put it in PATH so it behaves like a real CLI tool
chmod +x ~/scripts/deploy.sh
export PATH="$HOME/scripts:$PATH"
# Now from anywhere:
deploy.sh

# Rename without extension so it feels like a "real" command
mv ~/scripts/deploy.sh ~/scripts/deploy
deploy   # runs from anywhere, no .sh needed

# Avoid duplicate PATH entries when re-sourcing .bashrc repeatedly
case ":$PATH:" in
  *":$HOME/scripts:"*) ;; # already there, do nothing
  *) export PATH="$HOME/scripts:$PATH" ;;
esac
```

### ⚠️ Important Notes

> ⚠️ **Order matters.** Directories earlier in `$PATH` are searched first. If two directories have a script with the same name, the earlier one "wins" — this can cause confusing version conflicts (e.g., system Python vs. a custom Python).

> ⚠️ Scripts added to a custom `$PATH` directory still need **execute permission** (`chmod +x script.sh`) or they won't run even though the shell can "see" them.

> ⚠️ Every time you re-source `.bashrc`, any `export PATH="$HOME/scripts:$PATH"` line runs again, potentially **duplicating** entries. It's harmless for functionality but clutters `echo $PATH` output — see the dedup example above.

> ⚠️ Never put `.` (the current directory) in `$PATH`, especially not at the front — it's a well-known security risk since it lets malicious scripts in a working directory silently shadow real commands.

---

## 📊 Quick Reference Table

|Command|Purpose|Persists?|Common Gotcha|
|---|---|---|---|
|`export VAR=value`|Set an environment variable for this shell + children|Session only|Lost when terminal closes|
|`printenv`|List/view environment variables|N/A (read-only)|Returns nothing for non-exported vars|
|`printenv VAR`|View one variable|N/A|Exit code 1 if unset|
|`.bashrc` / `.zshrc`|Startup config file, runs on new shell|Permanent (across sessions)|Doesn't affect already-open shells|
|`source file` (or `. file`)|Re-run a script in current shell|Session, immediately|`./file` instead of `source` won't persist changes|
|`$PATH`|List of dirs searched for commands|Permanent if exported in rc file|Needs `chmod +x`; order matters|

---

## 🎯 Pro Tips

1. **Always `source` after editing** — a config file change does nothing until you `source` it or open a new terminal.
2. **Prepend, don't just append, to `$PATH`** for personal script directories — this ensures your custom versions take priority over system defaults when names clash.
3. **Use `printenv | grep <name>`** instead of scrolling through the full list when checking for a specific variable.
4. **Keep secrets out of `.bashrc`** if you ever plan to share/version-control it — source a separate `.env.secrets` file instead and add that file to `.gitignore`.
5. **Test syntax before sourcing** with `bash -n ~/.bashrc` to catch errors before they break your shell.
6. **Use `${VAR:?message}`** in scripts to fail loudly and clearly if a required environment variable isn't set, rather than silently continuing with an empty value.
7. **Name custom scripts without `.sh`** once they're in your `$PATH` — it makes them feel like native CLI tools.

---

## 🔐 Best Practices — Environment Management

- **Separate concerns:** Keep general shell config (`aliases`, `prompt`, `history`) separate from environment variables and separate from secrets. Split into sourced files (`~/.bashrc.d/`) for maintainability.
- **Document your `$PATH` additions** with a comment above each `export PATH=...` line — future-you will forget why a folder is there.
- **Prefer project-local `.env` files + a loader** (like `direnv`) over polluting your global shell for project-specific variables — keeps your global environment clean.
- **Version control your dotfiles** (minus secrets) in a personal git repo so your setup is portable across machines.
- **Audit periodically:** run `printenv | sort` occasionally to spot stale or duplicate variables you no longer need.
- **Least privilege for secrets:** environment variables are convenient but not secure storage — for real secrets, prefer a secrets manager or at minimum a gitignored file with restrictive permissions (`chmod 600`).