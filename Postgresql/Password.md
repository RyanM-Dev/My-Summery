# 🔐 Step-by-Step Guide: Change PostgreSQL Password and Enable md5 Authentication on Linux

---

## 1. Log in to `psql` as the `postgres` user

```bash
sudo -u postgres psql
```

- Logs in without password due to `peer` authentication.
    

---

## 2. Change the `postgres` user password

```sql
ALTER USER postgres WITH PASSWORD 'your_new_password';
```

- Replace `'your_new_password'` with your secure password.
    

---

## 3. Exit the `psql` shell

```sql
\q
```

---

## 4. Edit `pg_hba.conf` to enable md5 authentication

- Typical file locations:
    
    ```
    /etc/postgresql/16/main/pg_hba.conf
    ```
    
    or
    
    ```
    /etc/postgresql/17/main/pg_hba.conf
    ```
    
- Open with nano:
    

```bash
sudo nano /etc/postgresql/16/main/pg_hba.conf
```

- Find the line:
    

```
local   all             postgres                                peer
```

- Change to:
    

```
local   all             postgres                                md5
```

---

## 5. Save and exit editor

- In nano: press `Ctrl + O` to save, then `Ctrl + X` to exit.
    

---

## 6. Restart PostgreSQL to apply changes

```bash
sudo systemctl restart postgresql
```

---

## 7. Test password authentication

```bash
psql -U postgres -W
```

- Enter the new password to connect.
    

---

## ⚠️ Notes

- `peer`: Authenticates OS user without password.
    
- `md5`: Password authentication, recommended for security.
    
- Enabling `md5` on local connections requires password authentication for all local clients.
    

