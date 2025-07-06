Absolutely! Let's go over **if**, **else if**, and **else** statements in Go, along with a few related features. Go keeps conditionals simple and clean — here's everything you need to know.

---

## ✅ Basic `if` Statement

```go
if condition {
    // code block runs if condition is true
}
```

### Example:

```go
age := 20
if age >= 18 {
    fmt.Println("You are an adult.")
}
```

---

## 🔁 `if-else` Statement

```go
if condition {
    // true case
} else {
    // false case
}
```

### Example:

```go
num := 5
if num%2 == 0 {
    fmt.Println("Even number")
} else {
    fmt.Println("Odd number")
}
```

---

## 🔁 `if-else if-else` Chain

```go
if condition1 {
    // executes if condition1 is true
} else if condition2 {
    // executes if condition2 is true
} else {
    // executes if none are true
}
```

### Example:

```go
score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 75 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: C or lower")
}
```

---

## 💡 Short Statement in `if`

Go allows you to declare a short variable **only for use inside the `if` statement**:

```go
if x := compute(); x > 10 {
    fmt.Println("x is big")
} else {
    fmt.Println("x is small")
}
```

* `x` is only accessible inside the `if` and `else` blocks.
* This is useful when the value is only needed for condition checking.

---

## 🔒 No Ternary Operator in Go

Go **does not** have the `condition ? trueVal : falseVal` syntax you see in C or JavaScript.

This won't work:

```go
// ❌ Invalid in Go
result := a > b ? a : b
```

You must use `if-else`:

```go
if a > b {
    result = a
} else {
    result = b
}
```

---

## 🧠 Summary

| Construct                      | Description                                  |
| ------------------------------ | -------------------------------------------- |
| `if`                           | Runs block if condition is true              |
| `if-else`                      | Runs alternative block if condition is false |
| `if-else if-else`              | Chains multiple conditions                   |
| `if shortStatement; condition` | Declares temporary variable in condition     |
| ❌ No ternary operator          | Use `if-else` instead                        |

---

