
# 🛑 Critical Section & Synchronization

A **critical section** is a part of your program that requires **exclusive access** to a shared resource to avoid data races.

---

## 🧩 Synchronizing Memory Access Example

The following Go code demonstrates **synchronization** using a `sync.Mutex` to prevent memory access issues:
(it's not idiomatic Go don't use it)

```go
var memoryAccess sync.Mutex
var value int

go func() {
    memoryAccess.Lock()
    value++
    memoryAccess.Unlock()
}()

memoryAccess.Lock()
if value == 0 {
    fmt.Printf("the value is %v.\n", value)
} else {
    fmt.Printf("the value is %v.\n", value)
}
memoryAccess.Unlock()
```

### What's Happening:

- The `sync.Mutex` is used to **lock** and **unlock** the critical section, ensuring **exclusive access** to `value`.
    
- This **solves the data race** problem by synchronizing memory access.
    

---

## 🚨 Race Condition Still Exists

- While we've solved the **data race**, we still have a **race condition**.
    
- The order of operations is **non-deterministic**:
    
    - We don’t know if the **goroutine** will run first or if the **if-else** block will run first.
        
- The **nondeterminism** is only partially narrowed, but not fully resolved.
    

---

### **Takeaway**: While `sync.Mutex` helps with memory synchronization, it **doesn’t solve all race conditions** — particularly when the **order of execution** is still uncertain.

---
# ⏱️ Performance Cost of Synchronization

Using synchronization techniques like `Lock()` and `Unlock()` ensures **safe memory access**, but it also comes with **performance trade-offs**.

---

## 🐢 Why It Can Slow Things Down

- Every time you call `Lock`, the program **pauses** until it acquires the lock.
    
- Repeatedly entering and exiting critical sections increases **overhead**.
    

---

## ❓ Two Key Questions to Ask

1. **Are critical sections accessed frequently?**  
    Frequent locking can significantly slow your program down.
    
2. **How large should each critical section be?**
    
    - **Too small** → frequent locking/unlocking = overhead.
        
    - **Too large** → longer lock hold = reduced concurrency.
        

> There’s **no fixed rule** — tuning critical sections requires experience and context-specific judgment. It’s as much an **art** as it is a science.

## ⚠️ Note

- Synchronization, like other concurrency models, has **complex trade-offs**.
---
# 🔒 Deadlocks in Go 

Deadlocks, livelocks, and starvation are concurrency issues that prevent your program from making forward progress. Let’s focus on **deadlocks**.

---

## 🧊 What Is a Deadlock?

A **deadlock** occurs when **all concurrent processes are waiting on each other**, and **none can proceed**. The program will never recover unless something external intervenes.

### 🛑 In Go:

The Go runtime can detect **some** deadlocks (e.g., when all goroutines are blocked), but it doesn’t prevent them.

---

## 🧪 Example of Deadlock in Go

```go
type value struct {
    mu sync.Mutex
    value int
}

var a, b value
var wg sync.WaitGroup

printSum := func(v1, v2 *value) {
    defer wg.Done()
    v1.mu.Lock()
    defer v1.mu.Unlock()
    time.Sleep(2 * time.Second)
    v2.mu.Lock()
    defer v2.mu.Unlock()
    fmt.Printf("sum = %v\n", v1.value + v2.value)
}

wg.Add(2)
go printSum(&a, &b)
go printSum(&b, &a)
wg.Wait()
```

### ❗ What happens here?

- First goroutine locks `a`, then waits for `b`.
    
- Second goroutine locks `b`, then waits for `a`.
    
- Both are waiting **indefinitely** — **deadlock**.
    

---


### ☑️ The Coffman Conditions for Deadlock

1. **Mutual Exclusion**  
    At least one resource must be held in a non-shareable mode; that is, only one process can use the resource at a time.
    
2. **Hold and Wait** _(also known as Wait For Condition)_  
    A process is currently holding at least one resource and is waiting to acquire additional resources that are currently held by other processes.
    
3. **No Preemption**  
    Resources cannot be forcibly taken away from a process; they can only be released voluntarily by the process holding them.
    
4. **Circular Wait**  
    There exists a set of processes `{P1, P2, ..., Pn}` where each process is waiting for a resource held by the next process in the chain, and the last process is waiting for a resource held by the first, forming a circular dependency.
    

	Our example meets **all four conditions**, so it's guaranteed to deadlock.

---

## 🛡️ Preventing Deadlocks

To avoid deadlocks, **break at least one Coffman Condition**. For example:

- Enforce a consistent **locking order**.
    
- Use **timeouts** or **try-lock** strategies.
    
- Avoid holding locks longer than needed.
    
![[ChatGPT Image May 8, 2025, 09_03_50 AM.png]]
---

🔑 **Takeaway**: Understand the Coffman conditions. If you prevent just **one**, you eliminate the risk of deadlocks in your concurrent program.

# 🤹‍♀️ Livelocks  
---

## 🧠 Concept

**Livelock** occurs when multiple concurrent operations **keep changing state** in response to each other, but **never make actual progress**.

> Like two people trying to avoid each other in a hallway by stepping side to side in perfect sync — **forever stuck in motion, never moving forward**.

---

## 🧪 Example Code: Simulating Livelock

```go
// Create a condition variable with a mutex for synchronization
cadence := sync.NewCond(&sync.Mutex{})

// Broadcast cadence signal every 1ms (same rate for all)
go func() {
    for range time.Tick(1 * time.Millisecond) {
        cadence.Broadcast()
    }
}()

// Simulate waiting for cadence before taking a step
takeStep := func() {
    cadence.L.Lock()
    cadence.Wait()
    cadence.L.Unlock()
}

// Try to move in a direction (left or right)
tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
    fmt.Fprintf(out, " %v", dirName)

    // Declare intent to move by incrementing the direction counter
    atomic.AddInt32(dir, 1)

    takeStep() // wait for cadence

    // If only one person is trying this direction, move successfully
    if atomic.LoadInt32(dir) == 1 {
        fmt.Fprint(out, ". Success!")
        return true
    }

    takeStep() // failed to move, wait again

    // Undo intent to move (give up on that direction)
    atomic.AddInt32(dir, -1)
    return false
}
```

---

## 🧍 Walk Function (Represents One Person)

```go
walk := func(walking *sync.WaitGroup, name string) {
    var out bytes.Buffer
    defer func() { fmt.Println(out.String()) }()
    defer walking.Done()

    fmt.Fprintf(&out, "%v is trying to scoot:", name)

    // Try up to 5 times to step left or right
    for i := 0; i < 5; i++ {
        if tryLeft(&out) || tryRight(&out) {
            return
        }
    }

    // Give up after 5 failed attempts
    fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
}
```

---

## 🧍‍♀️🧍‍♀️ Launch Two People Simultaneously

```go
var peopleInHallway sync.WaitGroup
peopleInHallway.Add(2)

go walk(&peopleInHallway, "Alice")
go walk(&peopleInHallway, "Barbara")

peopleInHallway.Wait()
```

---

## 🖨️ Output

```
Alice is trying to scoot: left right left right left right left right left right
Alice tosses her hands up in exasperation!
Barbara is trying to scoot: left right left right left right left right left right
Barbara tosses her hands up in exasperation!
```

> Alice and Barbara keep switching directions at the same cadence, **blocking each other endlessly**.

---

## 🎯 Key Takeaways

- **Livelocks occur when concurrent processes try to avoid conflict (like a deadlock), but without coordination.**
    
- The processes are _active_, but their efforts are **futile** — the program doesn’t move forward.
    
- **Difficult to detect**: CPU usage may appear normal, logs may be active — but nothing productive is happening.
    
- Livelocks are a type of **starvation**, where processes are denied progress over time.
    

---

# 🍽️ Starvation in Go — Notes (Structured Format)

---

## 🧠 Concept

**Starvation** occurs when a concurrent process **cannot access the resources it needs** to perform its work. As a result, it either works inefficiently — or not at all.

- In **livelock**, _all_ processes are equally starved, doing no real work.
    
- In **starvation**, _some_ processes may dominate access to resources, **blocking others unfairly**.
    

---

## 🧪 Example Code: Greedy vs. Polite Worker

```go
var wg sync.WaitGroup
var sharedLock sync.Mutex
const runtime = 1 * time.Second

// Greedy worker grabs the lock once per loop and holds it
greedyWorker := func() {
    defer wg.Done()
    var count int
    for begin := time.Now(); time.Since(begin) <= runtime; {
        sharedLock.Lock()
        time.Sleep(3 * time.Nanosecond)
        sharedLock.Unlock()
        count++
    }
    fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}

// Polite worker breaks work into smaller chunks, releasing the lock more often
politeWorker := func() {
    defer wg.Done()
    var count int
    for begin := time.Now(); time.Since(begin) <= runtime; {
        sharedLock.Lock()
        time.Sleep(1 * time.Nanosecond)
        sharedLock.Unlock()

        sharedLock.Lock()
        time.Sleep(1 * time.Nanosecond)
        sharedLock.Unlock()

        sharedLock.Lock()
        time.Sleep(1 * time.Nanosecond)
        sharedLock.Unlock()

        count++
    }
    fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
}

wg.Add(2)
go greedyWorker()
go politeWorker()
wg.Wait()
```

---

## 🖨️ Sample Output

```
Polite worker was able to execute 289777 work loops.
Greedy worker was able to execute 471287 work loops
```

> Even though both simulate the **same total work time** (3ns per loop), the greedy worker **monopolizes the shared lock**, starving the polite worker.

---

## 🎯 Key Observations

- **Both workers do the same amount of work**, but the greedy one gets almost **2× more done**.
    
- This isn’t due to faster code — it’s because the **greedy worker holds the lock longer**, preventing the polite one from accessing the resource.
    

---

## 🔍 Detecting Starvation

- Use **metrics and logging**: measure how often each process completes its work.
    
- Compare performance: a significant difference in throughput may indicate starvation.
    

---

## 🛠️ Other Forms of Starvation

Starvation isn’t limited to Go mutexes. It can happen with:

- **CPU time**
    
- **Memory**
    
- **File handles**
    
- **Database connections**
    
- Any **shared, limited resource**
    

---

## 🔑 Takeaway

Starvation leads to **inefficient or unfair execution** in concurrent programs. Be cautious of holding locks longer than necessary, even unintentionally — it may **block others from progressing**.


# ⚙️ Go's Concurrency & Garbage Collection 

---

## 🗑️ Go’s Low-Latency Garbage Collector

### 💬 The Debate

- Some developers argue **garbage collectors (GC)** are unsuitable for:
    
    - **Real-time system performance**
        
    - **Deterministic performance**  
        Because GCs typically pause the program to reclaim memory.
        

### ✅ Go’s Advantage

- Since **Go 1.8**, GC **pause times are typically between 10–100 microseconds**, which is extremely low-latency.
    
- In most applications, **you don’t need to worry** about GC internals anymore.
    

---

## 🧵 Concurrency Management by the Runtime

- Go **automatically multiplexes goroutines onto OS threads**.
    
- You don’t need to:
    
    - Manually manage thread pools
        
    - Distribute tasks across threads
        
    - Write custom logic to pause or resume task execution
        

### 🧪 Example Scenario: Web Server

- In many languages:
    
    - You must create a **thread pool**.
        
    - Assign each incoming connection to a thread.
        
    - Implement your own CPU-sharing and scheduling logic.
        
- In **Go**:
    
    - Just use **`go func()`** to handle each connection.
        
    - The runtime takes care of:
        
        - Scheduling
            
        - Fairness
            
        - Thread assignment
            

---

## 🔧 Simplicity in Composing Concurrency

- Go’s built-in concurrency primitives (like goroutines, channels, `select`) make it **easier to model and compose** complex concurrent systems.
    
- Developers can focus on **problem-solving**, not low-level thread management.
    

---

## 🧠 Takeaway

Go combines **efficient garbage collection** and **powerful concurrency primitives** to simplify the development of **high-performance**, **scalable**, and **elegant concurrent systems** — without requiring low-level thread control.