package main

import "fmt"

func main() {
	fmt.Println("Welecome to Pointers....")

	myNumber := 25
	ptr := &myNumber
	fmt.Println("My value is : ", myNumber)
	fmt.Println("My ptr value : ", ptr)
	fmt.Println("My ptr value : ", *ptr)
}
