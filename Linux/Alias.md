


# 🧰 Linux Aliases: Simple and Clear Tutorial

A **quick guide** to creating and managing command shortcuts (aliases) in Linux (Ubuntu-based systems).

---

## 🔹 What is an alias?

An **alias** is a shortcut command that replaces a longer command. It's useful for saving time when you run repetitive commands.

---

## 🔹 Creating a Temporary Alias

This alias only works during the **current terminal session**.

```bash
alias shortcut_name='your_command_here'
````

**Example:**

```bash
alias myserver='ssh user@your.server.ip -p 22'
```

You can now just type `myserver` instead of the full SSH command.

> ⚠️ Temporary aliases are lost after closing the terminal.

---

## 🔹 Making an Alias Permanent

To make aliases **persistent**, add them to your shell config file.

### ✅ For Bash (default on Ubuntu):

1. Open the `.bashrc` file:
    
    ```bash
    nano ~/.bashrc
    ```
    
2. Add your alias at the bottom:
    
    ```bash
    alias myserver='ssh user@your.server.ip -p 22'
    ```
    
3. Save and exit (`CTRL+O`, then `CTRL+X`)
    
4. Apply changes:
    
    ```bash
    source ~/.bashrc
    ```
    

### ✅ For Zsh (Oh My Zsh, etc.):

1. Open `.zshrc`:
    
    ```bash
    nano ~/.zshrc
    ```
    
2. Add the alias as above, then apply it:
    
    ```bash
    source ~/.zshrc
    ```
    

---

## 🔹 Checking and Removing Aliases

- List all current aliases:
    
    ```bash
    alias
    ```
    
- Remove an alias for the session:
    
    ```bash
    unalias myserver
    ```
    

---

## 🔹 Example Aliases for Developers

```bash
alias gs='git status'
alias gc='git commit -m'
alias ll='ls -la'
alias up='sudo apt update && sudo apt upgrade'
alias dcu='docker-compose up'
alias dcd='docker-compose down'
```

---

## 🔹 MongoDB Aliases (From MongoDB Tutorial)

```bash
alias mongostart='sudo systemctl start mongod'
alias mongostop='sudo systemctl stop mongod'
alias mongorestart='sudo systemctl restart mongod'
alias mongolog='sudo journalctl -u mongod -f'
alias mongoshell='mongosh'
alias mongobackup='mongodump --out ~/mongo_backup'
alias mongorestore='mongorestore ~/mongo_backup'
```

---

## ✅ Tips

- You can create a separate alias file (e.g. `~/.bash_aliases`) and source it in your `.bashrc`.
    
- Use `functions` instead of aliases when arguments are needed.
    

---

## 📝 Example: Custom Alias File

1. Create and edit a new file:
    
    ```bash
    nano ~/.bash_aliases
    ```
    
2. Add your aliases.
    
3. In your `.bashrc`, ensure this line exists:
    
    ```bash
    if [ -f ~/.bash_aliases ]; then
        . ~/.bash_aliases
    fi
    ```
    
4. Apply changes:
    
    ```bash
    source ~/.bashrc
    ```
    

---

Happy automating! 🚀  
Let me know if you want a version tailored for Git, Docker, PostgreSQL, or Python development!
