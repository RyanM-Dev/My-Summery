# 🛠️ Fixing 404 Errors When Installing Packages on Ubuntu

If you're encountering 404 errors while fetching packages, it usually means the repository mirror is outdated or unreachable.

---

## ✅ Step 1: Update Package Repositories

Update your local package list:

```bash
sudo apt update
```

Then try installing the package again:

```bash
sudo apt install [your_package_name]
```

---

## 🔄 Step 2: Change Repository Mirror

If the issue persists, update your software source mirror:

1. Open the sources list:
    

```bash
sudo nano /etc/apt/sources.list
```

2. Replace all occurrences of:
    

```
http://ir.archive.ubuntu.com/ubuntu
```

with:

```
http://archive.ubuntu.com/ubuntu
```

3. Save and exit (`Ctrl + X`, then `Y`, then `Enter`).
    
4. Update again:
    

```bash
sudo apt update
```

---

## 🧰 Step 3: Use `--fix-missing` Option

Try installing with the `--fix-missing` flag to skip temporarily unavailable packages:

```bash
sudo apt install --fix-missing
```