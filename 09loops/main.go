package main

import (
	"fmt"
)

func main() {

	// 1. Classic for loop
	fmt.Println("1. Classic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}

	// 2. While-style loop
	fmt.Println("\n2. While-style loop:")
	j := 1
	for j <= 5 {
		fmt.Println("j =", j)
		j++
	}

	// 3. Infinite loop with break
	fmt.Println("\n3. Infinite loop (break at 3):")
	k := 1
	for {
		fmt.Println("k =", k)
		if k == 3 {
			break
		}
		k++
	}

	// 4. Using continue and break
	fmt.Println("\n4. Continue and Break:")
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue // skip 2
		}
		if i == 4 {
			break // stop at 4
		}
		fmt.Println("Value:", i)
	}

	// 5. Range loop over a slice
	fmt.Println("\n5. Range over slice:")
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 6. Range loop over a map
	fmt.Println("\n6. Range over map:")
	data := map[string]int{"a": 1, "b": 2}
	for key, val := range data {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	// 7. Range over a string
	fmt.Println("\n7. Range over string:")
	for i, ch := range "go语言" {
		fmt.Printf("Index: %d, Char: %c\n", i, ch)
	}
}
