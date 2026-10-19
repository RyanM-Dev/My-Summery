# 🧵 Concurrency in Go

Study notes for Katherine Cox-Buday’s *Concurrency in Go: Tools and Techniques for Developers*.

## 📚 Chapter Map

| Chapter | Focus | Notes |
|---|---|---|
| 1 | Concurrency foundations: races, atomicity, memory access, synchronization | [[01 - Chapter One/Race Condition]], [[01 - Chapter One/Atomicity]] |
| 2 | Concurrency, parallelism, CSP, and Go’s design | [[02 - Chapter Two/Parallelism & Concurrency]] |
| 3 | Goroutines, `sync`, channels, and `select` | [[03 - Chapter Three/1.Goroutine]], [[03 - Chapter Three/2.Sync]], [[03 - Chapter Three/3.Channel]] |
| 4 | Composable patterns: confinement, pipelines, fan-out/fan-in, cancellation | [[04 - Chapter Four/1.Confinement]], [[04 - Chapter Four/2.For Select Loop]], [[04 - Chapter Four/3.Preventing Goroutine Leaks]], [[04 - Chapter Four/4.Pipelines]], [[04 - Chapter Four/5.Fan Out,Fan In]] |
| 5 | Error propagation, timeouts, heartbeats, replicated requests, rate limiting, recovery | [[05 - Chapter Five/Concurrency at Scale]] |
| 6 | Runtime scheduling, work stealing, goroutines, and `GOMAXPROCS` | [[06 - Chapter Six/Goroutines and the Go Runtime]] |

## 🧭 Suggested Learning Path

1. Understand what makes concurrent access unsafe in Chapters 1–2.
2. Learn Go’s synchronization building blocks in Chapter 3.
3. Compose them into leak-resistant patterns in Chapter 4.
4. Add failure, cancellation, and capacity policies in Chapter 5.
5. Use the runtime model in Chapter 6 to interpret performance and diagnostics.

> [!tip] Link note
> Existing notes use several legacy filenames without spaces. Links above target those current filenames; rename only with the vault’s link-update workflow.

