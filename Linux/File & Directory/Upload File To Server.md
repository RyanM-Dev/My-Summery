# 📤 SCP File Upload — Step-by-Step Instruction

Use this guide to **upload a local folder to a remote server** via `scp` on a custom SSH port.

---

## ✅ Step 1: Use the Correct `scp` Syntax

To upload a **directory**, run:

```bash
scp -P [PORT] -r /path/to/local/folder user@host:/path/to/remote/destination/
```

### Example:

```bash
scp -P 6767 -r /home/ryan/Pictures/sample-photo/ root@95.416.170.57:/root/ryan/
```

> 🔸 `-P 6767`: specifies the SSH port  
> 🔸 `-r`: recursively uploads the folder  
> 🔸 Replace `root@95.416.170.57` with your actual server IP and username

---

## ✅ Step 2: Make Sure the Source Folder Exists

Check the folder on your local machine:

```bash
ls /home/ryan/Pictures/sample-photo/
```

---

## ✅ Step 3: Ensure the Destination Folder Exists (on the server)

SSH into your server:

```bash
ssh -p 6767 root@95.416.170.57
```

Create the target folder if it doesn't exist:

```bash
mkdir -p /root/ryan
```

---

## ⚠️ Common Mistake

Do **not** use `-p` (lowercase) — it means **preserve timestamps**, **not port**!

Always use `-P` (uppercase) for setting the SSH port.

---

Need an easier or more flexible alternative? Let me know — I can help you set up `rsync`, `sftp`, or a graphical tool like `FileZilla`.