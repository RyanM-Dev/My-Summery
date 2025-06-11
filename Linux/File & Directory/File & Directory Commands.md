

# 📁 File Management in Linux

## 📦 Copying Files and Directories

### 🔹 Copy a file

```bash
cp source.txt /path/to/destination/
```

### 🔹 Copy and rename a file

```bash
cp source.txt /path/to/destination/newname.txt
```

### 🔹 Copy a folder (recursively)

```bash
cp -r source_folder/ /path/to/destination/
```

- `-r` flag is used for copying directories and their contents.
    

---

## 🚚 Moving (Transferring) Files and Folders

### 🔹 Move a file

```bash
mv source.txt /path/to/destination/
```

### 🔹 Rename a file

```bash
mv oldname.txt newname.txt
```

### 🔹 Move a folder

```bash
mv folder_name/ /path/to/destination/
```

---

## 🧰 Quick Examples

```bash
cp myfile.txt ~/Documents/
cp -r Videos/Series/ /mnt/media_game/

mv notes.txt ~/Desktop/
mv Projects/AI/ /mnt/media_game/Backup/
mv oldname.mp4 newname.mp4
```

---

## 📁 Creating a Directory

### 🔹 Create a new directory in current location:

```bash
mkdir folderName
```

### 🔹 Create a directory at a specific path:

```bash
mkdir /path/to/folderName
```

---

## 🗑️ Removing a Directory

### 🔹 Remove an **empty** directory:

```bash
rmdir folderName
```

> ⚠️ `rmdir` only works if the directory is completely empty.

- If the directory contains **files or other directories**, `rmdir` will **not work**.
  
  # 📁  File & Directory 

---

## ✏️ Rename a File

```bash
mv currentFileName nextFileName
```

---

## 🚚 Move a File

```bash
mv fileName /path/to/destination/
```

---

## 📄 Copy a File

```bash
cp fileName /path/to/destination/
```

---

## 📂 Move/Copy a Directory (Including Contents)

### 🔸 Move:

```bash
mv -r directoryName /path/to/destination/
```

### 🔸 Copy:

```bash
cp -r directoryName /path/to/destination/
```

---

## 📦 Move/Copy Only Contents of a Directory

```bash
mv -r directoryName/* /path/to/destination/
cp -r directoryName/* /path/to/destination/
```

---

## 📍 Special Symbol

- `.` → refers to the **current directory**
    



---

## ❌ Remove a File

```bash
rm fileName
```

---

## 🧹 Remove a Directory and All Its Contents (Recursively & Forcibly)

```bash
rm -rf directoryName
```

---

## ⚠️ Remove Everything in the Current Directory

```bash
rm -rf *
```

---

## 🚨 CAUTION

> **Do NOT run** `rm -rf *` or similar commands in **critical directories** like:

- `/`
    
- `/root`
    
- `/etc`
    

Doing so can **permanently delete your operating system files** and make the system unbootable.