package main

import (
	"fmt"
)

// 1. Basic function
func sayHello() {
	fmt.Println("1. Hello from a basic function!")
}

// 2. Function with parameters
func greet(name string) {
	fmt.Println("2. Hello,", name)
}

// 3. Function with return value
func square(x int) int {
	return x * x
}

// 4. Function with multiple return values
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 5. Named return values
func add(a, b int) (sum int) {
	sum = a + b
	return
}

// 6. Function that takes another function as a parameter
func operate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 7. Closure: returns a function that remembers outer variable
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	// Call basic function
	sayHello()

	// Call function with parameter
	greet("Alice")

	// Call function with return
	fmt.Println("3. Square of 4 is:", square(4))

	// Multiple returns
	q, r := divide(10, 3)
	fmt.Printf("4. Divide 10 by 3 → Quotient: %d, Remainder: %d\n", q, r)

	// Named return value
	fmt.Println("5. Sum using named return:", add(5, 6))

	// Anonymous function immediately executed
	fmt.Println("6. Anonymous function output:")
	func() {
		fmt.Println("   Hello from an anonymous function!")
	}()

	// Assign anonymous function to a variable and call
	greetFunc := func(name string) {
		fmt.Println("   Hi,", name)
	}
	greetFunc("Bob")

	// Function as value (higher-order function)
	multiply := func(a, b int) int {
		return a * b
	}
	result := operate(3, 4, multiply)
	fmt.Println("7. 3 * 4 using operate():", result)

	// Closure example
	fmt.Println("8. Closure example:")
	next := counter()
	fmt.Println("   Count 1:", next())
	fmt.Println("   Count 2:", next())
	fmt.Println("   Count 3:", next())
}
