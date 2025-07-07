Great! Let's walk through how to **write to** and **read from** a file in Go using the standard `os` and `io/ioutil` or `os` and `bufio` packages.

---

## ✅ Part 1: Writing to a File

There are two common ways:

1. **`os.WriteFile` (simpler)** — writes the whole content at once.
2. **Using `os.OpenFile` + `WriteString()`** — more control.

### 🔹 Method 1: Quick & Simple

```go
package main

import (
    "os"
    "log"
)

func main() {
    content := []byte("Hello, Go file write!\nThis is line 2.")

    // Write content to a file
    err := os.WriteFile("example.txt", content, 0644)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("File written successfully.")
}
```

* `0644` is the permission mode (read/write for owner, read for others).

---

### 🔹 Method 2: Manual Write with More Control

```go
file, err := os.Create("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("Hello, Go file write!\n")
file.WriteString("This is line 2.\n")
```

---

## ✅ Part 2: Reading from a File

### 🔹 Method 1: `os.ReadFile` (simplest)

```go
data, err := os.ReadFile("example.txt")
if err != nil {
    log.Fatal(err)
}

fmt.Println("File contents:")
fmt.Println(string(data))
```

### 🔹 Method 2: Reading Line by Line (using `bufio`)

```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text()) // Print each line
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

---

## 🔁 Full Example: Write and Then Read a File

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)

func main() {
    // Writing to the file
    file, err := os.Create("myfile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    file.WriteString("Go is awesome!\n")
    file.WriteString("Let's learn file I/O.\n")

    fmt.Println("Writing done.")

    // Reading from the file
    file, err = os.Open("myfile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    fmt.Println("Reading file line by line:")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
```

---

## ⚠️ Notes

* Always use `defer file.Close()` after opening or creating a file.
* Use `log.Fatal(err)` or handle `err` properly to avoid crashes.
* Use `os.ReadFile` for small files and `bufio.Scanner` for line-by-line reading.

---