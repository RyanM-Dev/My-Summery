# 🖥️ Vim Cheatsheet

## 📝 The Ultimate Reference for Efficient Text Editing

---

## 1️⃣ Exiting & Saving Commands

### 💾 Basic Commands
| Command | Action |
|---------|--------|
| `:w` | Save (write) file |
| `:q` | Quit (fails if unsaved changes) |
| `:wq` or `:x` | Save and quit |
| `ZZ` | Save and quit (Normal mode) |
| `ZQ` | Quit without saving (Normal mode) |

### 🚀 Advanced Commands
| Command | Action |
|---------|--------|
| `:q!` | Quit without saving (force) |
| `:w!` | Force write (override read‑only) |
| `:wq!` | Force save and quit |
| `:w filename` | Save as new file (like “Save As”) |
| `:saveas filename` | Save under a new name |
| `:wall` | Save all open buffers |
| `:qall` | Quit all buffers |
| `:wqall` | Save and quit all buffers |
| `:x` | Save only if changes were made, then quit |
| `:exit` | Same as `:x` |
| `:enew` | Discard changes and open a new empty buffer |

---

## 2️⃣ Editing Commands (Basic)

### ✏️ Entering Insert Mode
| Command | Action |
|---------|--------|
| `i` | Insert before cursor |
| `I` | Insert at beginning of line |
| `a` | Append after cursor |
| `A` | Append at end of line |
| `o` | Open new line below cursor |
| `O` | Open new line above cursor |

### 🗑️ Deleting Text
| Command | Action |
|---------|--------|
| `x` | Delete character under cursor |
| `X` | Delete character before cursor |
| `dd` | Delete entire line |
| `D` | Delete from cursor to end of line |
| `dw` | Delete a word (from cursor to end of word) |
| `d$` | Delete to end of line |

### 🔁 Undo & Redo
| Command | Action |
|---------|--------|
| `u` | Undo last change |
| `Ctrl+r` | Redo |
| `U` | Undo all changes on current line (if cursor hasn’t moved) |

### 📋 Copy, Cut & Paste
| Command | Action |
|---------|--------|
| `yy` | Yank (copy) current line |
| `Y` | Same as `yy` |
| `yw` | Yank word |
| `y$` | Yank to end of line |
| `dd` | Cut line (delete into buffer) |
| `p` | Paste after cursor (or below line) |
| `P` | Paste before cursor (or above line) |

---

## 3️⃣ Editing Commands (Advanced)

### 🔧 Change Commands
| Command | Action |
|---------|--------|
| `cc` | Change entire line |
| `C` | Change from cursor to end of line |
| `cw` | Change word (from cursor to end of word) |
| `c$` | Change to end of line |
| `r<char>` | Replace single character under cursor |
| `R` | Enter Replace mode (overwrite until Esc) |
| `s` | Delete character and enter Insert mode |
| `S` | Delete line and enter Insert mode |

### 📐 Indentation & Shifting
| Command | Action |
|---------|--------|
| `>>` | Indent current line |
| `<<` | Unindent current line |
| `>%` | Indent block (when cursor on bracket) |
| `=%` | Auto‑indent block (when cursor on bracket) |

### 🧩 Text Objects (Combine with d, c, y, v)
| Command | Action |
|---------|--------|
| `di(` | Delete inside parentheses |
| `da(` | Delete around parentheses (includes brackets) |
| `ci"` | Change inside double quotes |
| `yi{` | Yank inside curly braces |
| `ca[` | Change around square brackets |

### 🔁 Macros & Repetition
| Command | Action |
|---------|--------|
| `q<letter>` | Start recording macro into register `<letter>` |
| `q` | Stop recording |
| `@<letter>` | Execute macro from register `<letter>` |
| `@@` | Repeat last macro |
| `10@<letter>` | Execute macro 10 times |
| `.` | Repeat last change (incredibly powerful) |

### 📦 Registers
| Command | Action |
|---------|--------|
| `"ayy` | Yank line into register `a` |
| `"ap` | Paste from register `a` |
| `"+y` | Yank to system clipboard (if compiled with +clipboard) |
| `:reg` | Display all register contents |

---

## 4️⃣ Navigation (Normal Mode)

### 🔹 Basic Movement
| Key | Action |
|-----|--------|
| `h` / `j` / `k` / `l` | Left / Down / Up / Right |
| `w` / `b` | Next / previous word |
| `e` | End of word |
| `0` / `^` | Start of line / first non‑blank |
| `$` | End of line |
| `gg` / `G` | Go to first / last line |
| `:N` or `NG` | Go to line N |
| `Ctrl+f` / `Ctrl+b` | Page down / up |
| `Ctrl+d` / `Ctrl+u` | Half‑page down / up |

### 🔹 Advanced Movement
| Command | Action |
|---------|--------|
| `f<char>` | Move **to** next occurrence of char (within line) |
| `F<char>` | Move **to** previous occurrence of char |
| `t<char>` | Move **until** next occurrence of char |
| `T<char>` | Move **until** previous occurrence of char |
| `%` | Jump to matching bracket/parenthesis |
| `*` / `#` | Search forward / backward for word under cursor |
| `m{a-z}` | Set mark (e.g., `ma`) |
| `'a` | Jump to mark `a` (line) |
| `` `a `` | Jump to exact position of mark `a` |

---

## 5️⃣ Search & Replace

### 🔍 Searching
| Command | Action |
|---------|--------|
| `/pattern` | Search forward for pattern |
| `?pattern` | Search backward |
| `n` | Next match |
| `N` | Previous match |
| `*` | Search forward for word under cursor |
| `#` | Search backward for word under cursor |
| `:set hlsearch` | Highlight all matches |
| `:noh` | Clear highlighting |

### 🔄 Search & Replace (`:substitute`)
```vim
:[range]s/pattern/replacement/[flags]
```
- **range** – e.g., `%` for whole file, `1,5` for lines 1‑5
- **flags** – `g` replace all in line, `c` confirm each, `i` ignore case

**Examples:**
```vim
:s/old/new/          " replace first occurrence in current line
:s/old/new/g         " replace all in current line
:%s/old/new/g        " replace all in entire file
:%s/old/new/gc       " replace all with confirmation
:5,10s/foo/bar/g     " replace in lines 5 to 10
```

---

## 6️⃣ Visual Mode

### 🔹 Selecting Text
| Command | Action |
|---------|--------|
| `v` | Character‑wise visual mode |
| `V` | Line‑wise visual mode |
| `Ctrl+v` | Block‑wise visual mode |

### 🔹 Operations on Selection
- `d` – delete selection
- `y` – yank (copy)
- `>` – indent right
- `<` – indent left
- `~` – toggle case
- `u` – convert to lowercase (selection)
- `U` – convert to uppercase

---

## 7️⃣ Working with Multiple Files, Windows & Tabs

### 🔹 Buffers
| Command | Action |
|---------|--------|
| `:e filename` | Open file in new buffer |
| `:ls` | List buffers |
| `:bnext` / `:bprev` | Next / previous buffer |
| `:bdelete` | Close buffer |
| `Ctrl+^` | Toggle between last two buffers |

### 🔹 Windows (Splits)
| Command | Action |
|---------|--------|
| `:split` / `:sp` | Horizontal split |
| `:vsplit` / `:vsp` | Vertical split |
| `Ctrl+w s` | Horizontal split (Normal mode) |
| `Ctrl+w v` | Vertical split |
| `Ctrl+w w` | Cycle between windows |
| `Ctrl+w h/j/k/l` | Move to window left/down/up/right |
| `Ctrl+w q` | Close current window |
| `Ctrl+w =` | Equalize window sizes |

### 🔹 Tabs
| Command | Action |
|---------|--------|
| `:tabnew` | Open new tab |
| `:tabnext` / `:tabprev` | Go to next / previous tab |
| `gt` / `gT` | Same as above (Normal mode) |
| `:tabclose` | Close current tab |

---

## 📊 Quick Reference Table

| Task | Command |
|------|---------|
| **Save** | `:w` |
| **Quit** | `:q` |
| **Save & Quit** | `:wq` or `ZZ` |
| **Quit without saving** | `:q!` or `ZQ` |
| **Insert before cursor** | `i` |
| **Append after cursor** | `a` |
| **Open line below** | `o` |
| **Delete character** | `x` |
| **Delete line** | `dd` |
| **Copy line** | `yy` |
| **Paste** | `p` |
| **Undo** | `u` |
| **Redo** | `Ctrl+r` |
| **Search** | `/pattern` |
| **Replace all** | `:%s/old/new/g` |
| **Go to line** | `:N` |
| **Start of file** | `gg` |
| **End of file** | `G` |
| **Indent line** | `>>` |
| **Split window** | `:split` / `:vsplit` |

---

## 🎯 Pro Tips

> [!TIP] **Learn touch typing in Vim** – The real power comes when you no longer think about keys.

> [!NOTE] **Use `.` and macros** – They turn repetitive tasks into single keystrokes.

> [!WARNING] **Don't stay in Insert mode** – Always return to Normal mode with `Esc` after typing. It’s the key to efficiency.

> [!IMPORTANT] **Customize your `.vimrc`** – Set options like `set number`, `set relativenumber`, `set mouse=a`, etc.

> [!EXAMPLE] **Combine motions and actions** – `d3w` deletes three words, `c$` changes to end of line, `yiw` copies a word.

---

## 🔧 Essential `.vimrc` Settings

```vim
set number              " show line numbers
set relativenumber      " relative line numbers
set autoindent          " auto-indent new lines
set smartindent         " smart indenting
set tabstop=4           " tab width = 4 spaces
set shiftwidth=4        " indent width = 4 spaces
set expandtab           " use spaces instead of tabs
set ignorecase          " case-insensitive search
set smartcase           " override ignorecase if uppercase used
set hlsearch            " highlight search results
set incsearch           " incremental search
set mouse=a             " enable mouse support
set clipboard=unnamedplus " use system clipboard
syntax on               " syntax highlighting
```

---

This cheatsheet covers the essential Vim commands, starting with saving/exiting and editing basics, then expanding to advanced editing, navigation, search, and more. Keep it in your Obsidian vault for quick reference! 🚀