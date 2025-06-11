# ⚛️ Atomicity 

**Atomicity** means that something is **indivisible** and **uninterruptible** **within a defined context**. It will fully complete without anything else occurring in that same context while it's executing.

---

## 📌 Context Defines Atomicity

- Whether an operation is atomic **depends on the scope** you're working in.
    
- This dependency can **help or hurt** you depending on the situation.
    
- So, the **first step** in thinking about atomicity is always: **Define the context**.
    

---

## 🧪 Example: `i++` — Is it Atomic?

Even though `i++` looks atomic, it's actually made up of **three operations**:

1. Retrieve the value of `i`
    
2. Increment it
    
3. Store the new value back in `i`
    

While each of these steps might be atomic on its own, **the full `i++` is not necessarily atomic** unless the context guarantees it.

- **In a single-threaded program**: It’s atomic **in that context**.
    
- **In a goroutine where `i` isn’t shared**: Still atomic.
    
- **In a concurrent context where `i` is shared**: Not atomic — race conditions can happen.
    

---

## 🚦 Why Atomicity Matters

- **In concurrent contexts**, atomic operations are **safe by design**.
    
- They allow you to build **logically correct programs** without unexpected interference.
    
- Atomic operations can also help **optimize performance** in multi-threaded environments.
    

---

🧠 **Bottom Line**: Atomicity is **always context-dependent**. Without defining the scope, you can't assume an operation is truly atomic.