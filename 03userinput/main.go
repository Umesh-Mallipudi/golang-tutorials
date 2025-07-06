package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welecomeMsg := "Welecome to Pizza Rating App (:"
	fmt.Println(welecomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Rating For OO Pizza: ")

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the Rating : ", rating)
	fmt.Printf("Type of the rating variable: %T ", rating)
}
