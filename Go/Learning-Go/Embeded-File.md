# 📦 Go `//go:embed` 

The `//go:embed` directive was introduced in **Go 1.16** as part of the `embed` package. It allows you to **embed files directly into your Go binary**, so the program doesn't rely on external files at runtime.

---

## 🧪 Your Code Explained

```go
//go:embed service.toml
var _svcBuf []byte
```

This line tells Go:

> Take the `service.toml` file (placed next to this source file) and embed its contents into the `_svcBuf` variable as a `[]byte`.

---

## 🧱 How It Works

1. **Import the `embed` package** (required):
    
    ```go
    import _ "embed"
    ```
    
2. **Use `//go:embed` directly before a variable**:
    
    - Only allowed types: `string`, `[]byte`, or `embed.FS`.
        
    - `string`: embeds the file as a text string.
        
    - `[]byte`: embeds the file as raw binary data.
        
    - `embed.FS`: used to embed multiple files or directories.
        

---

## 🔍 In Your Case

- `_svcBuf` contains the raw bytes of the `service.toml` file.
    
- The content is parsed using a config loader like `koanf` with `rawbytes.Provider`.
    

---

## ✅ Benefits of Using `embed`

- No external file dependencies at runtime.
    
- Ideal for **CLI tools and services** that need to be self-contained.
    
- Easier and more secure for deployment.
    

---

Let me know if you’d like examples of embedding multiple files or directories too.