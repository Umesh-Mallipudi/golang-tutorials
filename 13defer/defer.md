Great question — `defer` is an important and powerful keyword in Go.

---

## 🔹 What is `defer` in Go?

The `defer` keyword is used to **delay the execution** of a function **until the surrounding function returns**.

In other words:

> You can say: “Run this *after everything else is done*, just before we leave the function.”

---

## ✅ Syntax

```go
defer someFunction()
```

The `someFunction()` will **not run immediately** — it runs *after* the current function finishes, just before returning.

---

## 🧪 Simple Example

```go
func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred: end of function")
    fmt.Println("Middle")
}
```

### Output:

```
Start
Middle
Deferred: end of function
```

See how the `defer` line runs **last**, even though it's placed in the middle?

---

## 🔄 Multiple `defer` Statements: LIFO Order

Go executes deferred calls in **Last-In, First-Out** (LIFO) order — like a stack.

```go
func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
}
```

### Output:

```
Third
Second
First
```

---

## 📦 Common Use Cases of `defer`

### 1. **Close files / resources**

```go
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close() // Ensures it's closed when the function exits
```

### 2. **Unlock mutexes**

```go
mu.Lock()
defer mu.Unlock()
// Do some work with the shared resource
```

### 3. **Recover from panics**

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fmt.Println(a / b) // Will panic if b == 0
}
```

---

## ⚠️ Gotcha: Arguments are Evaluated Immediately

Even though the function runs later, the **arguments** are evaluated at the time the `defer` is declared:

```go
func main() {
    x := 10
    defer fmt.Println("Deferred x =", x)
    x = 20
}
```

### Output:

```
Deferred x = 10
```

---

## ✅ Summary

| Feature         | Behavior                                      |
| --------------- | --------------------------------------------- |
| `defer`         | Runs a function after the current one returns |
| Execution order | Last-in, first-out (LIFO)                     |
| Use cases       | Cleanup (files, locks), recover from panic    |
| Eval timing     | Arguments are evaluated immediately           |

---

Let me know if you'd like to try a mini exercise with `defer`, or want to see how it's used in real projects.
