Welcome to **Piscine Go**.

Inspired by the intensive, project-driven methodology of 42 school, this boot camp throws you straight into the deep end. There is no hand-holding, no video lectures, and no introductory tutorials. Your only tools are the Go compiler, the official documentation, unit tests, and your own logical deduction.

In Go, simplicity is deliberate. You won't find traditional classes, classical inheritance, or exceptions. Instead, you will master composition, strict typing, explicit error handling, and a revolutionary concurrency model.

---

## Piscine Rules & Guidelines

1. **The Standard Library is Your Limit:** Unless explicitly stated otherwise, you may only use Go's built-in packages (`fmt`, `os`, `strings`, `sync`, etc.). No external dependencies (`go mod init piscine`).
2. **Strict Compilation:** All code must compile with `go build` with zero warnings or errors.
3. **Test-Driven Rigor:** Every exercise must include or pass a suite of unit tests (`go test -v`). Writing your own tests is half the battle.
4. **Idiomatic Go:** Code must adhere to `gofmt` standards.

---

## Module 1: Slices & Maps (Memory & Collections)

*In Go, arrays have fixed lengths. Slices provide dynamic, flexible views over arrays, while maps provide hash-table lookups. Mastering how slices share underlying arrays is critical to avoiding subtle memory bugs.*

### Exercises:

* **`exercise/slice_trim`**: Write a function `TrimEdges(s []int, n int) []int` that returns a slice removing the first $n$ and last $n$ elements **without allocating a new underlying array** if possible, modifying or slicing the window directly. Understand slice headers (pointer, length, capacity).
* **`exercise/word_count`**: Write a function `WordFrequency(text string) map[string]int` that counts the frequency of each case-insensitive word in a string, stripping punctuation without using regex.
* **`exercise/group_by`**: Implement a generic-adjacent or slice-based grouping function: given a slice of strings, group them by their length into a `map[int][]string`.

---

## Module 2: Structs & Methods (Behavioral Encapsulation)

*Go does not have classes. Instead, structs hold state, and methods are functions with a receiver argument. You will learn value receivers vs. pointer receivers.*

### Exercises:

* **`exercise/vector2d`**: Create a `Vector2D` struct with `X` and `Y` float64 fields. Implement methods:
* `Add(v Vector2D) Vector2D` (value receiver)
* `Scale(factor float64)` (**pointer receiver** to mutate in place)
* `Magnitude() float64`


* **`exercise/nested_json`**: Define a nested struct hierarchy representing an organizational chart (Company $\rightarrow$ Department $\rightarrow$ Employee). Implement a custom method that serializes the tree into a formatted, indented string layout manually (without relying entirely on `encoding/json` magic initially, then compare with it).

---

## Module 3: Interfaces (Implicit Polymorphism)

*Go interfaces are satisfied implicitly—no `implements` keyword. If a type implements the methods, it fits the interface. This decouples definitions from implementations.*

### Exercises:

* **`exercise/plugin_system`**: Define an interface `Transformer`:
```go
type Transformer interface {
    Transform(input string) (string, error)
}

```


Implement three types that satisfy this interface: `UppercaseTransformer`, `Rot13Transformer`, and `ReverseTransformer`. Write a pipeline executor that chains multiple transformers together sequentially.
* **`exercise/shape_calc`**: Create a geometry interface with `Area()` and `Perimeter()` methods. Implement it for `Rectangle`, `Circle`, and `Triangle`. Write a function `PrintReport(s Shape)` that uses type assertions (`s.(Circle)`) to extract specific properties when available.

---

## Module 4: Generics (Type Parameters)

*Introduced to reduce code duplication, Go generics allow functions and data structures to operate on typed parameters without losing static type safety.*

### Exercises:

* **`exercise/generic_stack`**: Implement a generic LIFO Stack data structure `Stack[T any]` supporting `Push(T)`, `Pop() (T, bool)`, `Peek() (T, bool)`, and `IsEmpty() bool`.
* **`exercise/slice_utils`**: Implement generic utility functions:
* `Map[T, U any](s []T, f func(T) U) []U`
* `Filter[T any](s []T, pred func(T) bool) []T`
* `Reduce[T, U any](s []T, initial U, f func(U, T) U) U`



---

## Module 5: Concurrency (Goroutines & Channels)

*Do not communicate by sharing memory; instead, share memory by communicating. Goroutines are lightweight threads managed by the Go runtime, and channels coordinate them.*

### Exercises:

* **`exercise/worker_pool`**: Build a concurrent worker pool. Given a channel of jobs (integers representing file sizes to process) and a worker count $N$, spin up $N$ worker goroutines that process jobs concurrently and send results to a shared results channel. Ensure clean shutdown without leaking goroutines.
* **`exercise/ping_pong`**: Implement a synchronization exercise using channels where two goroutines ("Ping" and "Pong") bounce a ball counter back and forth up to 100 times before gracefully terminating using `sync.WaitGroup` or closed channels.
* **`exercise/rate_limiter`**: Build a token-bucket rate limiter using `time.Ticker` and channels that allows a maximum of $X$ requests per second.

---

## Module 6: The Rush (Final Capstone Project)

*The Piscine ends with a 48-hour "Rush"—a collaborative or individual intense project where you build a fully functional application from scratch.*

### Project: **GoKV — A Concurrent In-Memory Key-Value Store**

Build a lightweight, thread-safe, TCP-based or CLI-driven Key-Value store in Go.

* **Requirements:**
1. **Data Storage:** Uses custom internal maps protected by `sync.RWMutex` (or sharded maps for performance).
2. **Generics/Interfaces:** Supports multiple data types (strings, integers, lists) using interface values or generic abstractions.
3. **Concurrency:** Handles multiple concurrent client operations safely.
4. **Expiration (TTL):** Background goroutines running periodically to sweep and clean up expired keys.
5. **Test Suite:** Must pass 100% of a provided black-box integration test suite checking race conditions (`go test -race`).