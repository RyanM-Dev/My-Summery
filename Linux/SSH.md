# 🖥️ **Complete SSH Tutorial: Key Management and Setup** 🔑

This tutorial covers everything you need to know about using **SSH keys**, including how to generate, set up, and use them on remote servers, all explained up to a **moderate level**.

---

## **What is SSH and SSH Keys?** 🔒

SSH (Secure Shell) is a cryptographic protocol for securely accessing and managing remote servers. It uses a pair of **public** and **private** keys for authentication:

- **Private Key**: Kept secure on your machine. **Do not share** it. 🔑
    
- **Public Key**: Shared with the server. It is added to the server's `~/.ssh/authorized_keys` file. 🌍
    

---

## **1. Generating SSH Keys** 🛠️

To generate an SSH key pair (private and public), use the following command:

```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

### **Explanation**:

- `-t rsa`: Specifies the key type (RSA).
    
- `-b 4096`: Sets the key length to 4096 bits (stronger security).
    
- `-C "your_email@example.com"`: Adds a comment (usually your email) to the key.
    

### **Prompts after running the command**:

1. **File location**: Press **Enter** to accept the default (`~/.ssh/id_rsa`).
    
2. **Passphrase**: Optionally set a passphrase for added security. 🔐
    

Your **public key** will be stored in `~/.ssh/id_rsa.pub` and your **private key** in `~/.ssh/id_rsa`.

---

## **2. Viewing Your SSH Public Key** 👀

To view your **public key**, run:

```bash
cat ~/.ssh/id_rsa.pub
```

Copy the full output to add it to the server. 📋

---

## **3. Copying Your Public Key to the Server** 🚀

### **Option 1: Using `ssh-copy-id`** ⚡

If you can access the server with a username and password, use `ssh-copy-id`:

```bash
ssh-copy-id username@server_ip
```

After entering your password, the public key will automatically be added to the server's `~/.ssh/authorized_keys` file. ✅

### **Option 2: Manually Copying the Public Key** 📝

If `ssh-copy-id` is unavailable or you prefer manual copying:

1. **Copy the public key**:
    
    ```bash
    cat ~/.ssh/id_rsa.pub
    ```
    
    Copy the entire output. 📋
    
2. **Log into the server**:
    
    ```bash
    ssh username@server_ip
    ```
    
3. **Create the `.ssh` directory**:
    
    ```bash
    mkdir -p ~/.ssh
    ```
    
4. **Edit the `authorized_keys` file**:
    
    ```bash
    nano ~/.ssh/authorized_keys
    ```
    
5. **Paste the public key** into the file.
    
6. **Set the correct permissions**:
    
    ```bash
    chmod 700 ~/.ssh
    chmod 600 ~/.ssh/authorized_keys
    ```
    

---

## **4. Testing SSH Login Using Keys** ✅

Once the public key is added to the server, log in without a password:

```bash
ssh username@server_ip
```

If you set a passphrase for your private key, you’ll need to enter it. 🔑

---

## **5. Troubleshooting** 🛠️

If SSH login doesn't work, ensure:

1. **Permissions**: Make sure `.ssh` folder and `authorized_keys` file have the correct permissions (`700` for `.ssh`, `600` for `authorized_keys`).
    
2. **Correct File Format**: Ensure the public key is correctly pasted into the `authorized_keys` file without extra spaces or line breaks.
    
3. **Server Logs**: Check SSH server logs (`/var/log/auth.log` on Linux) for errors. 📜
    

---

## **6. Using SSH Config File for Convenience** ⚙️

To make frequent SSH logins easier, configure the `~/.ssh/config` file:

1. Edit or create the config file:
    
    ```bash
    nano ~/.ssh/config
    ```
    
2. Add the server configuration:
    
    ```
    Host myserver
      HostName server_ip
      User username
      IdentityFile ~/.ssh/id_rsa
    ```
    

Now, simply use:

```bash
ssh myserver
```

---

## **7. Managing Multiple SSH Keys** 🔑

If you use multiple SSH keys, specify which one to use for a server:

```bash
ssh -i ~/.ssh/other_key username@server_ip
```

Alternatively, define the key in your `~/.ssh/config` file.

---

## **8. Adding and Removing SSH Keys from the Server** ➕

- **To add a new key**, repeat the process of copying the public key to the `authorized_keys` file.
    
- **To remove a key**, edit `~/.ssh/authorized_keys` and delete the line corresponding to the key.
    

---

## **9. SSH Key Management Tools** 🧰

- **`ssh-agent`**: Manages private keys in memory to avoid repeatedly entering passphrases.
    
    - Start the agent: `eval $(ssh-agent)`
        
    - Add a key to the agent: `ssh-add ~/.ssh/id_rsa`
        
- **`keychain`**: A more advanced tool for managing SSH keys across multiple sessions.
    

---

## **Summary of Key Commands** 📋

|**Action**|**Command**|
|---|---|
|**Generate SSH Key Pair**|`ssh-keygen -t rsa -b 4096 -C "your_email@example.com"`|
|**View Public Key**|`cat ~/.ssh/id_rsa.pub`|
|**Copy Public Key to Server (Automatic)**|`ssh-copy-id username@server_ip`|
|**Manually Copy Public Key to Server**|`nano ~/.ssh/authorized_keys` (paste key)|
|**Test SSH Key Login**|`ssh username@server_ip`|
|**Set SSH Config File**|`nano ~/.ssh/config`|
|**Use Specific SSH Key**|`ssh -i ~/.ssh/other_key username@server_ip`|
|**Manage SSH Agent**|`ssh-add ~/.ssh/id_rsa`|
|**Remove SSH Key from Server**|Edit `~/.ssh/authorized_keys` and delete the line|

---

## **Using a Custom Port for SSH Connection** 🌐

If your server is using a port other than the default `22`, use the `-p` option:

```bash
ssh -p port_number username@server_ip
```

### **Example**:

```bash
ssh -p 2222 username@server_ip
```

---

## **Using a Custom Port with SSH Config File** ⚙️

To use a custom port with an SSH config file:

1. Edit `~/.ssh/config`:
    
    ```bash
    nano ~/.ssh/config
    ```
    
2. Add the configuration:
    
    ```
    Host myserver
      HostName server_ip
      User username
      Port 2222
      IdentityFile ~/.ssh/id_rsa
    ```
    
3. Now, use:
    
    ```bash
    ssh myserver
    ```
    

This will automatically use port `2222` for the connection. 🔄

---

By following these steps, you can securely set up and manage SSH keys, streamline your login process, and efficiently handle custom ports. Let me know if you need further clarification or encounter any issues! 🖥️