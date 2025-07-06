// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var score int

// 	fmt.Print("Enter your score (0-100): ")
// 	fmt.Scanln(&score)

// 	// Basic if-else if-else
// 	if score >= 90 {
// 		fmt.Println("Grade: A")
// 	} else if score >= 80 {
// 		fmt.Println("Grade: B")
// 	} else if score >= 70 {
// 		fmt.Println("Grade: C")
// 	} else if score >= 60 {
// 		fmt.Println("Grade: D")
// 	} else {
// 		fmt.Println("Grade: F")
// 	}

// 	// if with short statement
// 	if pass := score >= 60; pass {
// 		fmt.Println("Status: Passed ✅")
// 	} else {
// 		fmt.Println("Status: Failed ❌")
// 	}

// 	// Example of an exact match check using if only
// 	if score == 100 {
// 		fmt.Println("Perfect score! 🎉")
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	// Example 1: Switch on a string value
	fmt.Print("Enter a day of the week: ")
	var day string
	fmt.Scanln(&day)

	switch day {
	case "Monday":
		fmt.Println("Ugh, it's Monday 😩")
	case "Friday":
		fmt.Println("Yay! It's Friday 🎉")
	case "Saturday", "Sunday":
		fmt.Println("Weekend vibes 🏖️")
	default:
		fmt.Println("Just a regular day.")
	}

	// Example 2: Switch with conditions (no expression)
	fmt.Print("\nEnter your score (0-100): ")
	var score int
	fmt.Scanln(&score)

	switch {
	case score == 100:
		fmt.Println("Perfect score! 🏆")
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Needs improvement 📚")
	}

	// Example 3: Demonstrating fallthrough
	fmt.Print("\nEnter a number (1-2): ")
	var num int
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("You entered One")
		fallthrough
	case 2:
		fmt.Println("This is Two (fallthrough executed)")
	default:
		fmt.Println("Number not in range")
	}
}
