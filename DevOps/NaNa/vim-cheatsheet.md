# 📝 Vim Cheat Sheet

A practical Vim reference with the commands you’ll actually use.

> ⚠️ **Important: `:` is only used for Vim command-line (Ex) commands.**
>
> Examples that **need `:`**:
>
> ```vim
> :w
> :q
> :set number
> :echo @+
> ```
>
> Normal-mode commands **do not use `:`**:
>
> ```vim
> yy
> dd
> p
> "+yy
> "+p
> ```
>
> For clipboard commands such as `"+yy` or `"+p`, first press `Esc` to make sure you are in **Normal mode**, then type the keys directly. Do **not** type `:` first.

---

## 🚪 Open, Save & Quit

> These are Vim command-line (Ex) commands, so they **do start with `:`**.

| Command | Action |
|---|---|
| `vim file.txt` | Open a file |
| `:w` | Save |
| `:q` | Quit |
| `:wq` | Save and quit |
| `:x` | Save and quit |
| `:q!` | Quit without saving |
| `:w filename` | Save as another file |

---

## 🎛️ Vim Modes

| Key | Mode / Action |
|---|---|
| `Esc` | Return to **Normal mode** |
| `i` | Insert before cursor |
| `a` | Insert after cursor |
| `I` | Insert at beginning of line |
| `A` | Insert at end of line |
| `o` | New line below |
| `O` | New line above |
| `v` | Visual character selection |
| `V` | Visual line selection |
| `Ctrl+v` | Visual block selection |

> 💡 When confused, press `Esc` first.

---

## 🧭 Movement

### Basic

| Key | Action |
|---|---|
| `h` | Left |
| `j` | Down |
| `k` | Up |
| `l` | Right |

### Words

| Key | Action |
|---|---|
| `w` | Next word |
| `b` | Previous word |
| `e` | End of word |

### Lines

| Key | Action |
|---|---|
| `0` | Beginning of line |
| `^` | First non-space character |
| `$` | End of line |

### File Navigation

| Key | Action |
|---|---|
| `gg` | First line |
| `G` | Last line |
| `10G` | Go to line 10 |
| `:10` | Go to line 10 |
| `Ctrl+f` | Page down |
| `Ctrl+b` | Page up |
| `Ctrl+d` | Half-page down |
| `Ctrl+u` | Half-page up |

---

## ✂️ Copy, Cut & Paste

> 🧠 **Rule:** Copy/paste commands are **Normal-mode commands**, so they do **not** start with `:`.
>
> Before using them, press:
>
> ```text
> Esc
> ```
>
> Then type the command directly.

### Inside Vim

| Command | Action |
|---|---|
| `yy` | Copy current line |
| `3yy` | Copy 3 lines |
| `y` | Copy selected text |
| `dd` | Cut/delete current line |
| `3dd` | Cut/delete 3 lines |
| `p` | Paste after cursor |
| `P` | Paste before cursor |

Example:

```text
Esc
yy
p
```

There is **no `:`** before `yy` or `p`.

### 🖥️ System Clipboard

#### Copy one full line from Vim → system clipboard

First press:

```text
Esc
```

Then type:

```vim
"+yy
```

That means pressing these keys one by one:

```text
"
+
y
y
```

Do **not** type:

```vim
:"+yy
```

That would be incorrect because `"+yy` is a Normal-mode command.

---

#### Copy selected text from Vim → system clipboard

1. Press `Esc`
2. Press `v`
3. Select the text with `h`, `j`, `k`, `l` or the arrow keys
4. Type:

```vim
"+y
```

Again, there is **no `:`**.

---

#### Paste system clipboard → Vim

First press:

```text
Esc
```

Then type:

```vim
"+p
```

That means:

```text
"
+
p
```

To paste before the cursor:

```vim
"+P
```

Do **not** type `:` before these commands.

---

#### Check clipboard contents with an Ex command

This **does** require `:` because `echo` is a Vim command-line command:

```vim
:echo @+
```

To manually put text into the system clipboard register:

```vim
:let @+ = "vim clipboard test"
```

So remember:

```text
"+yy       ← Normal-mode command → NO :
"+p        ← Normal-mode command → NO :
:echo @+   ← Ex command          → YES :
:let @+... ← Ex command          → YES :
```

### ⭐ Recommended Clipboard Setting

Add this to `~/.vimrc`:

```vim
set clipboard=unnamedplus
```

When writing it interactively inside Vim, use:

```vim
:set clipboard=unnamedplus
```

because `:set` is an Ex command.

After enabling this setting, normal copy/paste commands:

```vim
yy
p
```

use the system clipboard automatically.

---

## 🗑️ Delete

| Command | Action |
|---|---|
| `x` | Delete character |
| `dd` | Delete line |
| `dw` | Delete word |
| `d$` | Delete to end of line |
| `d0` | Delete to beginning of line |
| `D` | Delete to end of line |

Useful combinations:

```vim
5dd
```

Delete 5 lines.

```vim
dG
```

Delete from current line to end of file.

```vim
dgg
```

Delete from current line to beginning of file.

---

## ↩️ Undo & Redo

| Command | Action |
|---|---|
| `u` | Undo |
| `Ctrl+r` | Redo |
| `U` | Undo changes on current line |

---

## 🔍 Search

Search forward:

```vim
/text
```

Search backward:

```vim
?text
```

Then:

| Key | Action |
|---|---|
| `n` | Next result |
| `N` | Previous result |

Search for the word under cursor:

```vim
*
```

Search backward for word under cursor:

```vim
#
```

Remove search highlighting:

```vim
:noh
```

---

## 🔄 Find & Replace

> Find/replace uses Vim's command line, so these commands **start with `:`**.

Replace first occurrence on current line:

```vim
:s/old/new/
```

Replace every occurrence on current line:

```vim
:s/old/new/g
```

Replace throughout entire file:

```vim
:%s/old/new/g
```

Ask for confirmation before each replacement:

```vim
:%s/old/new/gc
```

Useful confirmation keys:

| Key | Action |
|---|---|
| `y` | Replace |
| `n` | Skip |
| `a` | Replace all |
| `q` | Quit |

---

## ✏️ Editing

| Command | Action |
|---|---|
| `r` | Replace one character |
| `R` | Replace mode |
| `cw` | Change word |
| `cc` | Replace entire line |
| `C` | Change to end of line |
| `~` | Toggle character case |
| `J` | Join current line with next line |

Example:

```vim
cw
```

Deletes the current word and enters Insert mode.

---

## 📋 Visual Selection

### Character selection

```vim
v
```

### Line selection

```vim
V
```

### Column / block selection

```vim
Ctrl+v
```

After selecting:

| Key | Action |
|---|---|
| `y` | Copy |
| `d` | Delete |
| `c` | Change |
| `>` | Indent |
| `<` | Unindent |

---

## 📐 Indentation

| Command | Action |
|---|---|
| `>>` | Indent line |
| `<<` | Unindent line |
| `=` | Auto-indent selection |
| `gg=G` | Auto-indent entire file |

For example:

```vim
5>>
```

Indent the next 5 lines.

---

## 🔢 Repeat Commands

Most Vim commands accept a number.

Examples:

```vim
5j
```

Move down 5 lines.

```vim
3yy
```

Copy 3 lines.

```vim
10dd
```

Delete 10 lines.

```vim
5w
```

Move forward 5 words.

---

## ⚡ Extremely Useful Commands

### Repeat last change

```vim
.
```

The `.` command is one of Vim's most useful features.

Example:

```vim
dd
```

Then move somewhere else and press:

```vim
.
```

It repeats the delete command.

---

### Jump between matching brackets

Put the cursor on:

```text
( )
{ }
[ ]
```

and press:

```vim
%
```

---

### Show line numbers

```vim
:set number
```

Disable them:

```vim
:set nonumber
```

Relative line numbers:

```vim
:set relativenumber
```

---

## 🪟 Multiple Files / Buffers

> Buffer/file-management commands below are Ex commands, so they **start with `:`**.

Open another file:

```vim
:e file.txt
```

List buffers:

```vim
:ls
```

Next buffer:

```vim
:bn
```

Previous buffer:

```vim
:bp
```

Delete/close current buffer:

```vim
:bd
```

---

## 🖼️ Split Windows

> Split commands below are Ex commands, so they **start with `:`**.

Horizontal split:

```vim
:split
```

or:

```vim
:sp
```

Vertical split:

```vim
:vsplit
```

or:

```vim
:vsp
```

Move between windows:

```text
Ctrl+w h
Ctrl+w j
Ctrl+w k
Ctrl+w l
```

Quickly switch to the next window:

```text
Ctrl+w w
```

---

## 🐚 Run Shell Commands

> Shell commands launched through Vim's command line **start with `:`**.

Run a command:

```vim
:!ls
```

Example:

```vim
:!git status
```

Open a shell:

```vim
:shell
```

Return to Vim:

```bash
exit
```

---


---

## `:` — When Do You Use It?

Use `:` when entering a **Vim command-line / Ex command**.

Examples:

```vim
:w
:q
:wq
:set number
:set clipboard=unnamedplus
:%s/old/new/g
:echo @+
:!git status
```

Do **not** use `:` for Normal-mode editing commands:

```vim
yy
dd
p
u
w
b
gg
G
"+yy
"+p
```

A simple way to remember it:

```text
Editing/moving/copying with keys → no :
Configuration/file/search commands → usually :
```

## 🧠 Commands Worth Memorizing

```text
Esc      Normal mode

i        Insert
a        Append
o        New line

h j k l  Move
w / b    Next / previous word
0 / $    Start / end of line
gg / G   Start / end of file

yy       Copy line
dd       Delete line
p        Paste

u        Undo
Ctrl+r   Redo

/text    Search
n / N    Next / previous match

:%s/a/b/g   Replace all

.        Repeat last change

:w       Save
:q       Quit
:wq      Save + quit
:q!      Quit without saving
```

---

## 💡 Good `~/.vimrc` Basics

A simple useful configuration:

```vim
set number
set relativenumber
set clipboard=unnamedplus

set tabstop=4
set shiftwidth=4
set expandtab
set autoindent

set ignorecase
set smartcase

set hlsearch
set incsearch

syntax on
```

### What these do

- `number` → show line numbers
- `relativenumber` → easier movement with `5j`, `7k`, etc.
- `clipboard=unnamedplus` → use system clipboard
- `expandtab` → spaces instead of tab characters
- `autoindent` → preserve indentation
- `ignorecase` → case-insensitive search
- `smartcase` → uppercase search becomes case-sensitive
- `hlsearch` → highlight search results
- `incsearch` → search while typing

---

## 🎯 Daily Vim Workflow

A typical editing flow:

```text
vim file.go

i               → start writing
Esc             → Normal mode

w / b           → move by word
0 / $           → start/end of line
gg / G          → top/bottom

yy              → copy
dd              → delete
p               → paste
u               → undo
Ctrl+r          → redo

/text           → search
n               → next match

:w              → save
:wq             → save and exit
```

---

> 🚀 **Vim principle:** Stay in Normal mode most of the time. Enter Insert mode only when you actually need to type text.
