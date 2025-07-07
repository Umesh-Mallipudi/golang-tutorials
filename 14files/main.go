package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Note: ")

	note, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// ✅ Open the file in append mode
	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ✅ Add timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	noteWithTime := fmt.Sprintf("[%s] %s", timestamp, note)

	// ✅ Write note to file
	_, err = io.WriteString(file, noteWithTime)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n✅ Note saved successfully!")

	// ✅ Read and print all notes
	fmt.Println("📄 All Notes:")
	viewNotes()

	fmt.Print("Do you want to clear all notes? (y/n): ")
	choice, _ := reader.ReadString('\n')
	if choice == "y\n" {
		clearNotes()
	}

}

func viewNotes() {
	file, err := os.Open("notes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func clearNotes() {
	err := os.WriteFile("notes.txt", []byte(""), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("🧹 Notes cleared.")
}
