

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


---

## 📂 Navigating Directories

### 🔄 **Go Back to Previous Directory**

```bash
cd -
```

- Moves you to the **last directory** you were in.
    
- Example:
    

```bash
cd /var/log
cd /etc
cd -   # ⬅ Back to /var/log
```

---

## 📋 Listing Files

### 1️⃣ **Show All Files (Including Hidden)**

```bash
ls -a
```

- Displays **all files**, including hidden files starting with `.`.
    
- Example hidden files: `.bashrc`, `.git`
    

---

### 2️⃣ **Show Detailed Info (Long Format)**

```bash
ls -l
```

- Shows:
    
    - **Permissions (Access Control)**
        
    - **Owner / Group**
        
    - **Size**
        
    - **Date Modified**
        
- Example output:
    

```
-rw-r--r--  1 user user  2340 Jul 30 10:12 myfile.txt
```

---

### 3️⃣ **Show Human-Readable Sizes**

```bash
ls -lh
```

- Adds **KB, MB, GB** to file sizes.
    
- Example:
    

```
-rw-r--r--  1 user user  2.3K Jul 30 10:12 myfile.txt
```

---

### 4️⃣ **List Directories with `/` Suffix**

```bash
ls -p
```

- Appends `/` after folder names.
    
- Helps distinguish between **files** and **directories** quickly.
    

---

### 💡 Bonus: Get Command Help or Human-Readable Option

```bash
ls -h
```

- **If used with `-l`**, it makes sizes **human-readable**.
    
- **If used alone**, in some distros `-h` is part of `-lh` for human size view.
    
- Use `man ls` for full help.
    


---

## 📝 Creating Files

### 1️⃣ **Create an Empty File with `touch`**

```bash
touch filename.txt
```

- Creates an **empty file** if it doesn’t exist.
    

---

### 2️⃣ **Create a File with `>` Operator**

```bash
> filename.txt
```

- **Also creates an empty file**, like `touch`.
    
- ⚠ **Warning:** If the file exists, `>` **empties its content!**
    

**Example:**

```bash
> mylog.txt    # Creates or empties mylog.txt
```

---

## 🗑 Removing Files & Directories

### 1️⃣ **Remove a File**

```bash
rm file.txt
```

- Deletes the specified **file only**.
    

---

### 2️⃣ **Remove an Empty Directory**

```bash
rmdir myfolder
```

- Deletes the **directory** if it is **empty**.
    
- Fails if the folder has files inside.
    

---

### 3️⃣ **Remove Files and Folders Recursively**

```bash
rm -fr foldername
```

- **f** = force (no confirmation)
    
- **r** = recursive (delete folder and all contents)
    
- ⚠ **Very Dangerous!**
    
    - Running `rm -fr /` will **erase your entire system**.
        
    - Always **double-check the path** before using.
        

**Safe Practice:**

```bash
rm -r foldername       # safer, asks for confirmation
rm -ri foldername      # interactive deletion
```

---

I can now **merge this with the first Linux commands** into a **complete Obsidian cheat sheet** for you.  
Do you want me to make it as a **single tidy markdown with tables + emojis**? It’ll be perfect for quick reference.

# 🐧 Linux File Redirection & Output Management

Here’s a **well-organized cheat sheet** for file creation, overwriting, appending, and error redirection with **examples and warnings**:

---

## ✍ Creating & Writing to Files

### 1️⃣ **Overwrite or Create New File (`>`)**

```bash
echo "Hello World" > file.txt
```

- **If file doesn’t exist:** creates it.
    
- **If file exists:** ❗ **Overwrites** its contents.
    

---

### 2️⃣ **Append to File (`>>`)**

```bash
echo "New Line" >> file.txt
```

- **If file exists:** adds content at the **end**.
    
- **If file doesn’t exist:** creates a **new file**.
    
- ✅ Safer if you **don’t want to lose existing data**.
    

---

## ⚙ Redirecting Errors

Linux distinguishes:

- `1` → Standard Output (**stdout**)
    
- `2` → Standard Error (**stderr**)
    

---

### 3️⃣ **Redirect Errors and Overwrite (`2>`)**

```bash
ls /no/such/path 2> error.log
```

- Sends **only errors** to `error.log`.
    
- **Overwrites** the file if it exists.
    

---

### 4️⃣ **Redirect Errors and Append (`2>>`)**

```bash
ls /no/such/path 2>> error.log
```

- Sends **errors to the end** of `error.log`.
    
- **Does not overwrite** existing content.
    

---

### 💡 Combined Example: Save Output & Errors Separately

```bash
command > out.log 2> error.log
```

- **Output** → `out.log`
    
- **Errors** → `error.log`
    

---

I can now **combine all three cheat sheets (navigation, deletion, redirection)** into **one professional Linux quick reference** with emojis and a small **examples table** for Obsidian.

Do you want me to prepare that **final master version**? It will be super clean for your notes.