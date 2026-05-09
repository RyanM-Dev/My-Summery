#### 🆕 New Commands

|**Command**|**Short Description**|**Main Options/Flags**|**Common Use Case**|
|:--|:--|:--|:--|
|`which`|Locate a command binary|N/A|Finding the absolute path of an executable tool.|
|`command -v`|Verify command location|`-v`|A POSIX-standard way to find a command's path.|
|`ls`|List directory contents|`-l` (long), `-h` (human)|Seeing what files and folders are in a directory.|

---

#### 📌 Concepts and Commands Explained

- The REPL (Read-Eval-Print Loop) cycle.
- Command-line syntax and argument conventions.
- The technical difference between a "shell" and the "command line."
- The search order used by the shell to evaluate commands.
- The purpose of the POSIX standard.

### 🔧 The REPL Concept

The Command-Line Interface (CLI) is a text-based environment that operates as a **REPL**, which stands for **Read-Eval-Print Loop**. This cycle consists of four distinct steps:

1. **Read:** The shell accepts the text input you type (e.g., `ls`) after you press Enter.
2. **Evaluate:** The shell looks up the requested program, finds its binary file, and tells the computer to execute it.
3. **Print:** The program emits output (text), and the shell prints that response to your terminal window.
4. **Loop:** Once the program finishes, the shell returns to the first step to accept new input.

### 🔧 Command-Line Syntax + Arguments

Commands on the Linux command line follow a standard functional form: `commandname options`.

- **Arguments as Functions:** In programming terms, the command name acts as a function and the options are arguments passed to it.
- **Documentation Conventions:** Help manuals (manpages) use brackets `[ ]` for optional items and ellipses `...` to indicate that you can pass zero or more arguments.
- **Flags:** Options are often called "flags" and start with hyphens (e.g., `-debug`). Many programs support both short versions (`-l`) and long versions (`--long`) for the same parameter.
- **Spaces and Quoting:** A space marks the end of an argument; therefore, any string containing spaces **must** be enclosed in single or double quotes to be treated as a single argument.

### 🔧 Shell vs. Command Line

While these terms are often used interchangeably, there is a technical distinction:

- **Command Line:** Describes the general concept of a text-based interface used to interact with a system or application.
- **Shell:** A specific program (such as **Bash** or **Zsh**) that implements the command-line environment and interprets the text you type.

### 🔧 How the Shell Finds Commands (Evaluation)

When you type a command and press Enter, the shell follows a specific search order to determine what to execute:

1. **Specified Path:** If you provide a full path (e.g., `/usr/bin/ls`) or a relative path (e.g., `./script.sh`), the shell uses that exact file immediately.
2. **Built-ins & Aliases:** The shell checks if the command is a "built-in" (coded into the shell itself) or a user-defined shortcut (alias).
3. **The $PATH Variable:** The shell searches through the directories listed in the `$PATH` environment variable (typically `/bin`, `/usr/bin`, and `/sbin`) in the order they appear.
4. **Error:** If no matching executable is found, the shell returns a `command not found` error.

### 💻 which and command -v

Ambiguity in command execution can lead to bugs or misunderstandings. To identify exactly which file is being run, you can use the `which <command>` utility. On systems where `which` is unavailable, you should use `command -v`, which is the **POSIX** (Portable Operating System Interface) standard for this functionality.

---

#### 📋 Examples

```
# Basic listing of files (Step 1 of REPL)
ls

# Using 'which' to find the absolute path of the 'ls' binary
which ls
# Output: /bin/ls

# Using the POSIX-compliant 'command -v' to find the 'grep' path
command -v grep
# Output: /bin/grep

# Quoting an argument that contains spaces to ensure it is
# treated as one argument
cat "my project notes.txt"

# Checking the current shell using an environment variable
echo $SHELL
```

---

#### 💡 Pro Tips & Best Practices

- **Troubleshoot with `which`:** If a command is behaving unexpectedly, run `which command-name` to ensure you aren't accidentally running a shadowed version or a broken alias.
- **PATH Order Matters:** The shell executes the _first_ match it finds in `$PATH`. If you have two versions of a tool, the directory listed earlier in the variable takes priority.
- **Always Quote Filepaths:** In scripts and commands, always wrap file paths in double quotes `"` to prevent the shell from splitting names that contain spaces.
- **Think in REPL:** Understanding that the shell is a loop helps you realize that shell scripts are essentially just text files containing the "Read" phase for many commands at once.
- **Stick to POSIX for Portability:** If you are writing a script meant to run on multiple types of Unix (like both Linux and macOS), use POSIX-standard commands like `command -v` instead of system-specific ones.

---

#### 🧪 Practice Tests

**Questions**

1. What are the four steps represented by the acronym **REPL**?
2. What is the technical difference between a "shell" and the "command line"?
3. If a command is not an alias or a built-in, which environment variable does the shell search to find the executable?
4. In command-line documentation, what do brackets `[ ]` usually indicate?
5. Why must you use quotes around a filename like `my data.csv` when passing it to a command?
6. (Practical) Write a command to find the absolute path of the `bash` executable.
7. (Scenario) You try to run a tool called `setup.sh` that is located in your current folder by typing `setup.sh`, but the shell says `command not found`. How do you tell the shell to run that specific file?

**Answers**

1. **Read**, **Evaluate**, **Print**, and **Loop**.
    
2. The **Command Line** is the general interface concept; the **Shell** is the specific program (like Bash) that implements it and interprets commands.
    
3. The **`$PATH`** environment variable.
    
4. They indicate **optional** arguments or flags.
    
5. Because a space normally marks the end of one argument and the start of another; quotes tell the shell to treat the string as a single argument.
    
6. `which bash` or `command -v bash`.
    
7. Use a relative path by prepending a dot and a slash: **`./setup.sh`**.
   
   
   #### 🆕 New Commands

|**Command**|**Short Description**|**Main Options/Flags**|**Common Use Case**|
|:--|:--|:--|:--|
|`pwd`|Print working directory|N/A|Seeing which "room" of the filesystem you are currently in.|
|`ls`|List directory contents|`-l` (long), `-h` (human-readable)|Viewing files and folders in a specific location.|
|`cd`|Change directory|N/A|Teleporting from your current folder to a different one.|
|`find`|Find files|`-type`, `-name`, `-iname`, `-exec`|Searching the filesystem tree for specific files or directories,.|
|`less`|Page through a file|N/A|Reading file content one screen at a time.|

#### 📌 Concepts and Commands Explained

- Unix filesystem basics and text-based representation.
- The critical distinction between absolute and relative file paths.
- Navigating the filesystem tree with `pwd`, `ls`, and `cd`,,.
- Advanced filesystem searching using the `find` utility.
- Interactive file viewing with the `less` pager.

### 🔧 Unix Filesystem Basics

The "filesystem" is a collection of directories and files that organizes all data on the system. In a command-line environment, everything is represented as text, and contents are only shown when explicitly requested. While it seems difficult at first, holding a mental map of this tree structure is often more efficient than using a graphical interface.

### 🔧 Absolute vs. Relative Paths

Understanding paths is a prerequisite for almost every Linux command.

- **Absolute Paths:** Provide the full "driving directions" to a file starting from the root directory (`/`). You can recognize them because they always start with a `/` character.
- **Relative Paths:** These are partial paths that use your current location as the starting point. They are often more convenient to type and are recognized by the fact that they do **not** start with a `/`.

### 💻 pwd

`pwd` stands for "print working directory". It identifies which "room" of the filesystem tree your command-line environment is currently visiting. Most new shell sessions begin in your home directory (e.g., `/home/dave`).

**Example 1: Basic usage**

```
pwd
```

**Expected Output:**

```
/home/dave
```

### 💻 ls

The `ls` command allows you to "list" files in a directory. Without arguments, it lists the current directory; otherwise, it lists the path provided.

- `-l`: The "long" listing shows extra metadata, including permissions, owner, group, and size.
- `-h`: The "human-readable" flag converts file sizes from raw bytes into megabytes or gigabytes.

**Example 1: Long listing with human-readable sizes**

```
ls -lh /var/log
```

**Expected Output:**

```
-rw-r--r--  1 root  root   4.6K Jun 24 02:31 alternatives.log
```

### 💻 cd

`cd` lets you "change directory" to any location on the filesystem. After running `cd`, the `pwd` command will reflect your updated location.

**Example 1: Changing to a specific system path**

```
cd /etc/ssl/certs
```

**Expected Output:** _(No output, but running `pwd` will now show `/etc/ssl/certs`)_

### 💻 find

`find` is used to search for files and directories. Unlike most commands, its flags typically use a single dash (e.g., `-name`).

- `-type d`: Search specifically for directories.
- `-name`: Search by name (case-sensitive).
- `-iname`: Search by name (case-insensitive).
- `-exec`: Run a command (like `echo`) on every found item.

**Example 1: Searching for a directory named 'home' from the root**

```
find / -type d -name home
```

**Expected Output:**

```
/home
```

### 💻 less

`less` is a "pager" that allows you to read a file one "page" at a time based on your terminal window size. It is interactive and ideal for large files.

- **Navigation:** Use **Arrow Keys** to scroll by line or **Spacebar** to scroll by page.
- **Searching:** Type `/` followed by a search string and hit **Enter**. Use `n` for the next match and `SHIFT+n` for the previous.
- **Quitting:** Type `q` to exit.

**Example 1: Reading a system file**

```
less /etc/hosts
```

**Expected Output:** _(The terminal window fills with the contents of the file in an interactive view)_

#### 📋 Examples

```
# Verify where you are currently located
pwd

# List contents of the /usr/local directory in long format
ls -l /usr/local

# Use a relative path to move into a subfolder
cd Documents

# Find all files in the current directory (.) ending in .txt
find . -name "*.txt"

# Search for the word 'root' inside the passwd file using the pager
less /etc/passwd
# (Inside less, type /root and hit Enter)
```

#### 💡 Pro Tips & Best Practices

- **Path Logic:** If a path starts with `/`, it is absolute; otherwise, it is relative to where you are right now,.
- **The "Less" Rule:** Always prefer `less` over `cat` for viewing files of unknown size to prevent your terminal from being flooded with thousands of lines of text.
- **Permission Denied:** When running `find` on system directories as a normal user, it is normal to see "Permission denied" errors for folders you aren't allowed to access.
- **Tab Completion:** While navigating with `cd`, use the **Tab** key to autocomplete directory names to save time and avoid typos,.

#### 🧪 Practice Tests

**Questions**

1. What character identifies the "root" directory in an absolute path?
2. Which command would you use to verify your current location in the filesystem tree?
3. What is the difference between the `find` flags `-name` and `-iname`?
4. When using `ls -lh`, what does the `-h` specifically accomplish?
5. How do you exit the `less` pager to return to your command prompt?
6. (Practical) Write a command to change your directory to `/var/log`.
7. (Practical) Construct a `find` command to search the current directory for any file starting with "config".
8. (Scenario) You are in `/home/dave/Desktop`. You want to see the contents of a folder named `notes` inside your current directory. Provide both the absolute and relative path commands to list its files.

**Answers**

1. The forward slash (`/`).
    
2. The `pwd` (print working directory) command.
    
3. `-name` is case-sensitive, while `-iname` is case-insensitive.
    
4. It makes the file size "human-readable" by showing KB, MB, or GB instead of raw bytes.
    
5. Press the `q` key.
    
6. `cd /var/log`.
    
7. `find . -name "config*"`.
    
8. Relative: `ls notes`; Absolute: `ls /home/dave/Desktop/notes`,.
   
   
   #### 🆕 New Commands

|**Command**|**Short Description**|**Main Options/Flags**|**Common Use Case**|
|:--|:--|:--|:--|
|`touch`|Create a file or update timestamps|`-a` (access), `-m` (modification)|Creating a new empty file quickly.|
|`mkdir`|Create a new directory|`-p` (parents/idempotent)|Making nested folder structures in one go.|
|`rmdir`|Remove empty directories|N/A|Deleting a folder only if it contains nothing.|
|`rm`|Remove files and directories|`-r` (recursive), `-f` (force)|Deleting files or entire directory trees.|
|`mv`|Move or rename files|N/A|Renaming a file or moving it to a new path.|
|`man`|Display manual pages|`[section number]`|Looking up documentation for any command.|
|`apropos`|Search manual page descriptions|N/A|Finding a command when you only know a keyword.|

---

#### 📌 Concepts and Commands Explained

- Modifying the filesystem by creating, moving, and deleting objects.
- The versatility of the `mv` command for both moving and renaming.
- Safety precautions when deleting data (specifically the dangers of `rm -rf`).
- Navigating and searching the built-in manual (manpages) system.
- Increasing efficiency using Shell Autocompletion (Tab-completion).

### 💻 touch

The `touch` command is used to create a new, empty file if the specified path does not exist. If the file already exists, `touch` updates its access and modification timestamps to the current time.

**Example 1: Basic usage (creating a file)**

```
touch newfile.txt
ls
```

**Expected Output:**

```
newfile.txt
```

**Example 2: Updating timestamps**

```
touch existingfile.txt
```

**Expected Output:** _(No output, but `ls -l` would show a refreshed modification time)._

### 💻 mkdir

`mkdir` stands for "make directory." It creates a new folder at the path provided. The `-p` flag is particularly important as it allows you to create nested directories and is "idempotent," meaning it won't throw an error if the directory already exists.

**Example 1: Creating multiple directories**

```
mkdir project1 project2
```

**Expected Output:** _(Two new folders created in the current directory)._

**Example 2: Creating nested directories with -p**

```
mkdir -p storage/logs/current
```

**Expected Output:** _(Creates the entire path even if 'storage' or 'logs' did not exist)._

### 💻 rmdir and rm

`rmdir` is a safe command that only removes directories if they are empty. In contrast, `rm` (remove) is used to delete files. To delete directories that contain files, you must use the `-r` (recursive) flag.

**Example 1: Deleting a file**

```
rm old_notes.txt
```

**Expected Output:** _(File is deleted)._

**Example 2: Forcing a recursive deletion**

```
rm -rf folder_to_delete
```

**Expected Output:** _(The folder and all its contents are deleted without confirmation)._

### 💻 mv

The `mv` (move) command is unique because it handles two tasks: moving a file to a different location and renaming a file. If the destination is a directory, the file is moved; if the destination is a new filename, the file is renamed.

**Example 1: Renaming a file**

```
mv draft.txt final_report.txt
```

**Expected Output:** _(The file 'draft.txt' is now named 'final_report.txt')._

**Example 2: Moving and renaming simultaneously**

```
mv notes.txt archives/old_notes.txt
```

**Expected Output:** _(File is moved into 'archives' and renamed to 'old_notes.txt')._

### 🔧 Getting Help + man, apropos

Linux systems ship with **manpages** (manual pages). These are organized into sections: Section 1 for general commands, Section 5 for file formats, and Section 8 for system administration. If you don't know the exact command name, use `apropos` to search by keyword.

**Example 1: Reading a manual page**

```
man ls
```

**Expected Output:** _(The terminal opens an interactive viewer for the 'ls' documentation)._

**Example 2: Searching for commands related to 'partition'**

```
apropos partition
```

**Expected Output:**

```
fdisk (8) - manipulate disk partition table
...
```

### 🔧 Shell Autocompletion (Tab-completion)

This is a time-saving feature where the shell completes a command or filename after you press the **Tab** key. If there are multiple matches, pressing Tab twice will list all possibilities.

**Example 1: Using Tab to complete a directory name**

```
cd Doc[TAB]
```

**Expected Output:**

```
cd Documents/
```

---

#### 📋 Examples

```
# Create a nested project structure
mkdir -p dev/project/bin

# Create an initial README file
touch dev/project/README.md

# Move the project folder to a new name
mv dev/project dev/linux_basics

# Search for help on how to remove things
apropos remove

# Safely delete the bins folder
rmdir dev/linux_basics/bin

# Delete the entire dev tree
rm -rf dev
```

---

#### 💡 Pro Tips & Best Practices

- **The Idempotency of `mkdir -p`:** Always use the `-p` flag in scripts. It ensures your script doesn't crash just because a directory already exists.
- **The Danger of `rm -rf /`:** Never run this. It tells Linux to delete the root directory, which will destroy the entire operating system.
- **Manual Page Pagers:** Since `man` uses a pager (like `less`), you can use `/` to search for text and `q` to quit the manual.
- **Check the Section:** If you need configuration file help, specify the section: `man 5 crontab`.
- **Tab-Completion Habit:** Develop the habit of hitting Tab every few characters. It prevents typos and verifies that the file you think is there actually exists.

---

#### 🧪 Practice Tests

**Questions**

1. Which command is used to update the modification time of an existing file?
2. What is the main difference between `rmdir` and `rm -r`?
3. If you want to rename `file1.txt` to `file2.txt`, which command do you use?
4. Which manual section contains information on system administration commands?
5. What happens if you run `mkdir -p` on a directory that already exists?
6. (Practical) Write a command to create a directory named `backups` and an empty file inside it named `today.log`.
7. (Practical) Write a command to delete a non-empty directory named `temp_files` without being prompted for confirmation.
8. (Scenario) You need to find a command that handles "passwords" but you can't remember the name. What command do you run?

**Answers**

1. The `touch` command.
    
2. `rmdir` only removes empty directories, while `rm -r` removes a directory and everything inside it.
    
3. `mv file1.txt file2.txt`.
    
4. Section 8.
    
5. Nothing happens; it is "idempotent" and will not report an error.
    
6. `mkdir backups && touch backups/today.log`.
    
7. `rm -rf temp_files`.
    
8. `apropos password` or `man -k password`.