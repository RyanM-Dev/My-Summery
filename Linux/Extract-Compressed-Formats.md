# 📦 Extracting Compressed Formats on Ubuntu / Debian

Quick reference for unpacking common archive and compression formats on Linux (Ubuntu/Debian).

---

## 🧰 Install Common Tools

Most formats work out of the box. Install the rest as needed:

```bash
sudo apt update
sudo apt install -y \
  tar gzip bzip2 xz-utils zip unzip \
  p7zip-full p7zip-rar \
  unrar-free \
  zstd \
  cabextract
```

| Package        | Formats / use              |
|----------------|----------------------------|
| `tar`          | `.tar`, combined archives  |
| `gzip`         | `.gz`                      |
| `bzip2`        | `.bz2`                     |
| `xz-utils`     | `.xz`, `.lzma`             |
| `zip` / `unzip`| `.zip`                     |
| `p7zip-full`   | `.7z`, many others         |
| `unrar-free`   | `.rar` (basic)             |
| `zstd`         | `.zst`                     |
| `cabextract`   | `.cab`                     |

> For full RAR support (newer RAR5), you may need `unrar` from multiverse:
>
> ```bash
> sudo apt install unrar
> ```

---

## 📋 Quick Cheat Sheet

| Extension              | Extract command |
|------------------------|-----------------|
| `.tar`                 | `tar -xvf file.tar` |
| `.tar.gz` / `.tgz`     | `tar -xzvf file.tar.gz` |
| `.tar.bz2` / `.tbz2`   | `tar -xjvf file.tar.bz2` |
| `.tar.xz` / `.txz`     | `tar -xJvf file.tar.xz` |
| `.tar.zst`             | `tar --use-compress-program=zstd -xvf file.tar.zst` |
| `.gz` (single file)    | `gunzip file.gz` or `gzip -d file.gz` |
| `.bz2` (single file)   | `bunzip2 file.bz2` or `bzip2 -d file.bz2` |
| `.xz` (single file)    | `unxz file.xz` or `xz -d file.xz` |
| `.zip`                 | `unzip file.zip` |
| `.7z`                  | `7z x file.7z` |
| `.rar`                 | `unrar x file.rar` |
| `.zst` (single file)   | `zstd -d file.zst` |

**Common flags (tar):**

| Flag | Meaning |
|------|---------|
| `-x` | extract |
| `-c` | create |
| `-v` | verbose (list files) |
| `-f` | file name follows |
| `-z` | gzip |
| `-j` | bzip2 |
| `-J` | xz |
| `-C` | extract into directory |
| `-t` | list contents (no extract) |

---

## 📁 Extract Into a Specific Folder

```bash
# Create destination first (optional but clean)
mkdir -p ~/Downloads/extracted

# tar family
tar -xzvf archive.tar.gz -C ~/Downloads/extracted

# zip
unzip archive.zip -d ~/Downloads/extracted

# 7z
7z x archive.7z -o~/Downloads/extracted

# rar
unrar x archive.rar ~/Downloads/extracted/
```

> Note: `7z` has **no space** after `-o` (`-o/path/to/dir`).

---

## 🔍 Inspect Before Extracting

```bash
# tar
tar -tzvf archive.tar.gz

# zip
unzip -l archive.zip

# 7z
7z l archive.7z

# rar
unrar l archive.rar
```

---

## 📄 Format Details

### 🔹 `.tar` — tape archive (no compression)

```bash
tar -xvf archive.tar
```

### 🔹 `.tar.gz` / `.tgz` — tar + gzip

```bash
tar -xzvf archive.tar.gz
# modern tar auto-detects compression:
tar -xvf archive.tar.gz
```

### 🔹 `.tar.bz2` / `.tbz2` — tar + bzip2

```bash
tar -xjvf archive.tar.bz2
# or
tar -xvf archive.tar.bz2
```

### 🔹 `.tar.xz` / `.txz` — tar + xz

```bash
tar -xJvf archive.tar.xz
# or
tar -xvf archive.tar.xz
```

### 🔹 `.tar.zst` — tar + zstd

```bash
tar --use-compress-program=zstd -xvf archive.tar.zst
# or pipe
zstd -d -c archive.tar.zst | tar -xvf -
```

### 🔹 `.gz` / `.bz2` / `.xz` / `.zst` — single-file compression

These compress **one file**, not a whole folder tree:

```bash
gunzip  file.gz      # removes .gz, leaves file
bunzip2 file.bz2
unxz    file.xz
zstd -d file.zst
```

Keep the original compressed file:

```bash
gunzip -k file.gz
# or
gzip -dk file.gz
```

### 🔹 `.zip`

```bash
unzip archive.zip
unzip archive.zip -d /path/to/output
unzip archive.zip file1.txt path/inside/file2.txt   # extract specific files
```

### 🔹 `.7z`

```bash
7z x archive.7z          # keep full paths
7z e archive.7z          # extract all files into current dir (flatten)
7z x archive.7z -o/tmp/out
```

### 🔹 `.rar`

```bash
unrar x archive.rar           # keep paths
unrar e archive.rar           # flatten into current dir
unrar x archive.rar ./outdir/
```

### 🔹 Multi-part archives

```bash
# zip split parts: file.zip, file.z01, file.z02 ...
unzip file.zip

# 7z multi-part: archive.7z.001, archive.7z.002 ...
7z x archive.7z.001

# rar multi-part: archive.part1.rar, archive.part2.rar ...
unrar x archive.part1.rar
```

---

## 🛠️ Universal Option: `7z`

`7z` can open many formats (zip, tar, gzip, bzip2, xz, 7z, rar, iso, …):

```bash
7z x anything.whatever
7z l anything.whatever   # list
```

Useful when you are unsure of the format:

```bash
file archive.unknown
```

---

## 🗜️ Create Archives (bonus)

```bash
# tar.gz
tar -czvf archive.tar.gz folder/

# tar.xz (often smaller)
tar -cJvf archive.tar.xz folder/

# zip
zip -r archive.zip folder/

# 7z (strong compression)
7z a archive.7z folder/
```

---

## ⚠️ Tips & Gotchas

1. **Always list first** if the archive is untrusted — malware can use path tricks:
   ```bash
   tar -tzvf archive.tar.gz | head
   ```
2. Prefer extracting into an **empty dedicated directory** so files do not scatter into `$PWD`.
3. Modern `tar` often **auto-detects** compression — `tar -xvf file.tar.*` is usually enough.
4. **`.deb` packages** are not normal “extract and run” archives. To peek inside:
   ```bash
   dpkg-deb -x package.deb ./out/
   dpkg-deb -e package.deb ./out/DEBIAN/   # control files
   ```
5. **Password-protected** archives:
   ```bash
   unzip -P 'password' secret.zip
   7z x secret.7z -p'password'
   unrar x -p'password' secret.rar
   ```

---

## 🧠 One-liner: extract by extension (bash)

Handy alias/function idea:

```bash
extract() {
  if [ ! -f "$1" ]; then
    echo "File not found: $1"
    return 1
  fi
  case "$1" in
    *.tar.bz2|*.tbz2) tar -xjvf "$1" ;;
    *.tar.gz|*.tgz)   tar -xzvf "$1" ;;
    *.tar.xz|*.txz)   tar -xJvf "$1" ;;
    *.tar.zst)        tar --use-compress-program=zstd -xvf "$1" ;;
    *.tar)            tar -xvf "$1" ;;
    *.bz2)            bunzip2 "$1" ;;
    *.gz)             gunzip "$1" ;;
    *.xz)             unxz "$1" ;;
    *.zst)            zstd -d "$1" ;;
    *.zip)            unzip "$1" ;;
    *.7z)             7z x "$1" ;;
    *.rar)            unrar x "$1" ;;
    *)                echo "Unknown format: $1" ;;
  esac
}
```

Add to `~/.bashrc` or `~/.zshrc`, then:

```bash
extract archive.tar.gz
```
