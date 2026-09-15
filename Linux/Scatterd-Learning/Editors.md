
# 📝  `nano` 

## 🧭 Basic Behavior

- `nano fileName`  
    Opens the file if it exists, or **creates it** if it doesn't.  
    **Example:**
    
    ```bash
    nano /home/ryan/help.txt
    ```
    
- `nano` _(without filename)_  
    Opens an empty editor. You’ll be asked to enter a filename **only when you try to save/exit**.
    

---

## 🧩 Shortcut Symbols in `nano`

- `^` = **CTRL** key  
    Example: `^X` means `CTRL + X`
    
- `M` = **ALT (Meta)** key  
    Example: `M-U` means `ALT + U`
    

---

## 💾 Saving and Exiting

- To **exit**, press:
    
    ```text
    CTRL + X
    ```
    
    - If there are unsaved changes, it will **ask if you want to save**.
        
    - Then it will prompt for a **file name** if needed.