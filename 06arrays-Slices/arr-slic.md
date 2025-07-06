
---

## 📦 **Array vs Slice in Go**

| Feature        | **Array**                    | **Slice**                 |
| -------------- | ---------------------------- | ------------------------- |
| Fixed Size?    | ✅ Yes (e.g., `[3]int`)       | ❌ No – dynamic size       |
| Growable?      | ❌ No                         | ✅ Yes, using `append()`   |
| Memory Usage   | Stores all elements directly | Points to a backing array |
| Most Used?     | 🚫 Rare                      | ✅ Commonly used in Go     |
| Syntax Example | `arr := [3]int{1, 2, 3}`     | `s := []int{1, 2, 3}`     |

---

## 🔎 **How Slice Works Internally**

A **slice** is not the actual data. It contains:

* A **pointer** to an underlying array
* A **length** (`len`) → how many elements are in use
* A **capacity** (`cap`) → how many elements can fit before resizing

---

## ➕ **append() in Go**

* `append(slice, item)` adds element(s) to a slice
* If `len < cap`, it uses the same underlying array
* If `len == cap`, Go creates a **new, bigger array**, copies elements, and returns a new slice

---

### ✅ Example

```go
s := []int{1, 2, 3}  // len=3, cap=3
s = append(s, 4)     // New array created (cap exceeded)
s = append(s, 5)     // Added to the new array (room available)
```

---

## 🧠 Summary Diagram (Conceptual)

```
s := []int{1, 2, 3}

Initial:
- len = 3
- cap = 3
- memory = [1, 2, 3]

After append(s, 4):
- len = 4
- cap = 6 (Go grows array)
- memory = [1, 2, 3, 4, _, _] ← new array

After append(s, 5):
- len = 5
- memory = [1, 2, 3, 4, 5, _]
```

---

## ✅ TL;DR Summary

* **Array** = fixed size, rarely used directly
* **Slice** = flexible, points to an array, can grow
* **append()** = adds data to a slice, may create a new array if full
* Slices are **more powerful** and **easier to use** in real Go programs

---

Let me know if you’d like a quick practice quiz or memory diagram!
