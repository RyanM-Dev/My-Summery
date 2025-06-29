# 🐧 Linux Log & File Viewing Commands 

---

## 📄 `head` — Show First Lines of a File

```bash
head filename
```

- **Default**: Shows the **first 10 lines**.
    
- **Customize line count**:
    

```bash
head -n 20 filename  # shows first 20 lines
```

---

## 📄 `tail` — Show Last Lines of a File

```bash
tail filename
```

- **Default**: Shows the **last 10 lines**.
    
- **Follow the log file (live updates)**:
    

```bash
tail -f /var/log/syslog
```

---

## 📝 `logger` — Write to System Log (`/var/log/syslog`)

```bash
logger "This is a test log message"
```

- Appends the message to the system log.
    
- Useful for scripting, debugging, or automation logs.
    

You can view it with:

```bash
tail /var/log/syslog
```

---

Let me know if you want to log with timestamps, tags, or priorities!