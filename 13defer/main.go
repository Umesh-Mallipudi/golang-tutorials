package main

import "fmt"

func main() {

	fmt.Println("Welecome to go lang defer world :)")
	defer fmt.Println("This is the defered function")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("This is the main function")
	myDefer()
	fmt.Println("")
}

func myDefer() {
	fmt.Println("myDefered function")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// Welecome , this is main
// 0,1,2,3,4
// this is defered , one, two
