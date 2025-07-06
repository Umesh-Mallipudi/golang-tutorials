package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// getting userinput
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you Name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Name ", name)
	// Request a greeting message.
	message, err := greetings.Hello(name)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
		log.Print(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
