Here's a comprehensive Obsidian-friendly markdown documentation for your Linux user and group management commands:

# 👥 Linux User & Group Management Commands

[[DevOps/NaNa/0. Nana DevOps Table of Contents|📚 Nana DevOps Table of Contents]]

> [!toc]- 📑 Contents
>
> - [[#📝 Command Documentation & Reference Guide|📝 Command Documentation & Reference Guide]]
> - [[#1️⃣ Creating Users: adduser vs useradd|1️⃣ Creating Users: `adduser` vs `useradd`]]
> - [[#2️⃣ Viewing User Information: cat /etc/passwd|2️⃣ Viewing User Information: `cat /etc/passwd`]]
> - [[#3️⃣ Changing Passwords: passwd|3️⃣ Changing Passwords: `passwd`]]
> - [[#4️⃣ Switching Users: su|4️⃣ Switching Users: `su`]]
> - [[#5️⃣ Creating Groups: groupadd vs addgroup|5️⃣ Creating Groups: `groupadd` vs `addgroup`]]
> - [[#6️⃣ Deleting Users: deluser vs userdel|6️⃣ Deleting Users: `deluser` vs `userdel`]]
> - [[#7️⃣ Deleting Groups: delgroup vs groupdel|7️⃣ Deleting Groups: `delgroup` vs `groupdel`]]
> - [[#8️⃣ Modifying Users: usermod|8️⃣ Modifying Users: `usermod`]]
> - [[#9️⃣ Adding Users to Multiple Groups: usermod -G|9️⃣ Adding Users to Multiple Groups: `usermod -G`]]
> - [[#🔟 Displaying Groups: groups|🔟 Displaying Groups: `groups`]]
> - [[#1️⃣1️⃣ Adding User to Group During Creation: useradd -G|1️⃣1️⃣ Adding User to Group During Creation: `useradd -G`]]
> - [[#1️⃣2️⃣ Removing User from Group: gpasswd -d|1️⃣2️⃣ Removing User from Group: `gpasswd -d`]]
> - [[#📊 Quick Reference Table|📊 Quick Reference Table]]
> - [[#🎯 Pro Tips|🎯 Pro Tips]]
> - [[#🔐 Security Best Practices|🔐 Security Best Practices]]


## 📝 Command Documentation & Reference Guide

---

## 1️⃣ Creating Users: `adduser` vs `useradd`

### 🔍 Difference Explained
| Feature | `adduser` | `useradd` |
|---------|-----------|-----------|
| **Type** | Perl script wrapper | Native binary |
| **Interface** | Interactive & friendly | Non-interactive |
| **Home Directory** | Auto-creates | Requires `-m` flag |
| **Default Config** | Uses `/etc/adduser.conf` | Uses `/etc/default/useradd` |
| **Password Prompt** | Automatically prompts | Requires separate `passwd` command |
| **Ease of Use** | Beginner-friendly | More control for advanced users |

### 📋 Structure
```bash
# adduser (Interactive)
sudo adduser [username]

# useradd (Non-interactive with flags)
sudo useradd [options] [username]
```

### 💡 Examples
```bash
# Interactive creation (recommended for beginners)
sudo adduser john_doe

# Non-interactive with specific options
sudo useradd -m -s /bin/bash -c "John Doe" john_doe
```

### 🚀 Advanced Use Cases
```bash
# Create user with custom home directory
sudo useradd -m -d /custom/home/path username

# Create system user (no login)
sudo useradd -r -s /usr/sbin/nologin service_user

# Create user with specific UID
sudo useradd -u 1500 username
```

---

## 2️⃣ Viewing User Information: `cat /etc/passwd`

### 📋 Structure
```bash
cat /etc/passwd
# or filter for specific user
grep username /etc/passwd
```

### 🔍 Understanding the Format
```
username:password:UID:GID:comment:home_directory:shell
```

### 💡 Examples
```bash
# View all users
cat /etc/passwd

# Find specific user
grep "^john" /etc/passwd
# Output: john:x:1001:1001:John Doe:/home/john:/bin/bash
```

### 🚀 Advanced Use Cases
```bash
# List only usernames
cut -d: -f1 /etc/passwd

# Find users with bash shell
grep "/bin/bash" /etc/passwd

# Sort users by UID
sort -t: -k3 -n /etc/passwd
```

---

## 3️⃣ Changing Passwords: `passwd`

### 📋 Structure
```bash
passwd [options] [username]
```

### 💡 Examples
```bash
# Change current user's password
passwd

# Change another user's password (requires sudo)
sudo passwd john_doe

# Lock a user account
sudo passwd -l username
```

### 🚀 Advanced Use Cases
```bash
# Force password change at next login
sudo passwd -e username

# Set password expiration
sudo passwd -x 30 username  # expire in 30 days

# Unlock account
sudo passwd -u username
```

---

## 4️⃣ Switching Users: `su`

### 📋 Structure
```bash
su [options] [-] [username]
```

### 🔍 The `-` Flag Explained
- `su username`: Switch user but keep current environment
- `su - username`: Switch user and load target user's full environment

### 💡 Examples
```bash
# Switch to root with full environment
su -

# Switch to specific user
su - john_doe

# Run single command as another user
su -c "command" - username
```

### 🚀 Advanced Use Cases
```bash
# Switch user with specific shell
su -s /bin/zsh - username

# Preserve environment variables
su -m username  # or su -p
```

---

## 5️⃣ Creating Groups: `groupadd` vs `addgroup`

### 🔍 Difference Explained
| Feature | `groupadd` | `addgroup` |
|---------|------------|------------|
| **Type** | Native binary | Script wrapper |
| **Interface** | Command-line | Interactive options |
| **Flexibility** | More options | Simpler interface |
| **Configuration** | Command flags | Uses config files |

### 📋 Structure
```bash
# groupadd (non-interactive)
sudo groupadd [options] groupname

# addgroup (interactive)
sudo addgroup groupname
```

### 💡 Examples
```bash
# Create basic group
sudo groupadd developers

# Create group with specific GID
sudo groupadd -g 2000 developers

# Create system group
sudo groupadd -r system_group
```

### 🚀 Advanced Use Cases
```bash
# Create group with non-unique GID
sudo groupadd -o -g 1000 shared_group

# Force creation even if exists
sudo groupadd -f existing_group
```

---

## 6️⃣ Deleting Users: `deluser` vs `userdel`

### 🔍 Difference Explained
| Feature | `deluser` | `userdel` |
|---------|-----------|-----------|
| **Type** | Debian-based script | Standard command |
| **Config File** | Uses `/etc/deluser.conf` | No config file |
| **Home Directory** | Auto-removes (configurable) | Requires `-r` flag |
| **Flexibility** | More automated | More manual control |

### 📋 Structure
```bash
# deluser (Debian/Ubuntu)
sudo deluser [options] username

# userdel (standard)
sudo userdel [options] username
```

### 💡 Examples
```bash
# Delete user only
sudo deluser john_doe

# Delete user AND home directory
sudo userdel -r john_doe

# Delete user and all their files
sudo deluser --remove-all-files john_doe
```

### 🚀 Advanced Use Cases
```bash
# Force delete even if user is logged in
sudo userdel -f username

# Delete user from specific group only
sudo deluser username groupname

# Backup before deletion
sudo tar -czf backup.tar.gz /home/username && sudo userdel -r username
```

---

## 7️⃣ Deleting Groups: `delgroup` vs `groupdel`

### 📋 Structure
```bash
# delgroup (Debian/Ubuntu)
sudo delgroup groupname

# groupdel (standard)
sudo groupdel groupname
```

### 💡 Examples
```bash
# Delete group
sudo groupdel developers

# Delete group with delgroup
sudo delgroup developers
```

### ⚠️ Important Notes
- Cannot delete a user's primary group
- Ensure no files remain with the group ownership

---

## 8️⃣ Modifying Users: `usermod`

### 📋 Structure
```bash
sudo usermod [options] username
```

### 💡 Examples
```bash
# Change user's primary group
sudo usermod -g developers john_doe

# Change username
sudo usermod -l newname oldname

# Change home directory
sudo usermod -d /new/home/path -m username
```

### 🚀 Advanced Use Cases
```bash
# Lock user account
sudo usermod -L username

# Set account expiry date
sudo usermod -e 2024-12-31 username

# Change user's shell
sudo usermod -s /bin/zsh username
```

---

## 9️⃣ Adding Users to Multiple Groups: `usermod -G`

### 📋 Structure
```bash
sudo usermod -G group1,group2,group3 username
```

### ⚠️ Important: `-G` vs `-aG`
- `-G`: **Replaces** all secondary groups
- `-aG`: **Appends** to existing groups (recommended!)

### 💡 Examples
```bash
# REPLACES all groups (dangerous!)
sudo usermod -G developers,testers john_doe

# APPENDS to existing groups (safe!)
sudo usermod -aG developers,testers john_doe
```

### 🚀 Advanced Use Cases
```bash
# Add to multiple groups with verification
sudo usermod -aG docker,sudo,developers username
groups username  # Verify

# Remove from all secondary groups
sudo usermod -G "" username
```

---

## 🔟 Displaying Groups: `groups`

### 📋 Structure
```bash
groups [username]
```

### 💡 Examples
```bash
# Show current user's groups
groups

# Show specific user's groups
groups john_doe

# Show group members
getent group groupname
```

### 🚀 Advanced Use Cases
```bash
# List all groups on system
cut -d: -f1 /etc/group

# Show users in a group
grep "^groupname:" /etc/group | cut -d: -f4

# Count user's groups
groups username | tr ' ' '\n' | wc -l
```

---

## 1️⃣1️⃣ Adding User to Group During Creation: `useradd -G`

### 📋 Structure
```bash
sudo useradd -G groupname username
```

### 💡 Examples
```bash
# Create user and add to single group
sudo useradd -G developers john_doe

# Create user and add to multiple groups
sudo useradd -G developers,testers,design john_doe

# Create with primary and secondary groups
sudo useradd -g primary_group -G secondary1,secondary2 username
```

### 🚀 Advanced Use Cases
```bash
# Create user with groups and home directory
sudo useradd -m -G docker,sudo -s /bin/bash dev_user

# Create user with specific UID and groups
sudo useradd -u 2000 -G developers -m username
```

---

## 1️⃣2️⃣ Removing User from Group: `gpasswd -d`

### 📋 Structure
```bash
sudo gpasswd -d username groupname
```

### 💡 Examples
```bash
# Remove user from group
sudo gpasswd -d john_doe developers

# Verify removal
groups john_doe
```

### 🚀 Advanced Use Cases
```bash
# Remove from multiple groups
for group in group1 group2 group3; do
    sudo gpasswd -d username $group
done

# Remove all users from a group
sudo gpasswd -M "" groupname

# Set specific users as group members (replaces all)
sudo gpasswd -M user1,user2,user3 groupname
```

---

## 📊 Quick Reference Table

| Command | Purpose | Key Flag |
|---------|---------|----------|
| `adduser` | Interactive user creation | Auto-config |
| `useradd` | Non-interactive user creation | `-m` for home |
| `passwd` | Change password | `-l` lock, `-e` expire |
| `su -` | Switch user | `-` for full env |
| `groupadd` | Create group | `-g` for GID |
| `deluser` | Delete user (Debian) | `--remove-all-files` |
| `userdel` | Delete user (standard) | `-r` remove home |
| `usermod -G` | Replace groups | ⚠️ Overwrites! |
| `usermod -aG` | Append groups | ✅ Safe! |
| `groups` | Display groups | No flags |
| `gpasswd -d` | Remove from group | `-d` delete |

---

## 🎯 Pro Tips

1. **Always use `-aG`** instead of `-G` when adding users to groups
2. **Backup before deletion**: `sudo tar -czf backup.tar.gz /home/user`
3. **Check before creating**: `getent passwd username` to verify existence
4. **Use meaningful GID ranges**: 
   - System: 0-999
   - Users: 1000+
5. **Audit regularly**: `sudo awk -F: '$3 >= 1000 {print $1}' /etc/passwd`

---

## 🔐 Security Best Practices

- 🚫 Never share passwords via command line
- 📝 Document all user/group changes
- 🔄 Regular audit of user accounts
- ⏰ Set password expiration policies
- 👥 Use groups for permission management instead of individual users

---

This documentation provides a comprehensive reference for Linux user and group management. Save this in your Obsidian vault for quick access! 📚
