package main

import "fmt"

const pi = 3.14

func main() {
	fmt.Println("Variables")

	var username string = "Umesh"
	fmt.Println("WELECOME TO THE GO LANG WORLD : ", username)
	fmt.Printf("Variable is of Type  %T \n", username)

	var smallInt uint = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type : %T \n", smallInt)

	var smallFloat float32 = 355.9099993838883
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	// Default values and some aliases and checking the value of the variable without value assign
	var anotherVariable int
	fmt.Println("Default value of int", anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var noName string
	fmt.Println("Default value of string", noName)
	fmt.Printf("Variable is of type : %T \n", noName)

	var nobool bool
	fmt.Println("Default value of bool", nobool)
	fmt.Printf("Variable is of type : %T \n", nobool)

	var noFloat float32
	fmt.Println("Default value of float32", noFloat)
	fmt.Printf("Variable is of type : %T \n", noFloat)

	var noBigFloat float64
	fmt.Println("Default value of float64", noBigFloat)
	fmt.Printf("Variable is of type : %T \n", noBigFloat)

	// implicit variable - no need to put type , lexer will add type to it based on the value
	var company = "software"
	fmt.Println("company value is : ", company)
	fmt.Printf("type of the company variable : %T \n", company)

	// no var keyword or walurs operator
	// no need to specify var keyword and type for vaiable
	noOfUsers := 5
	fmt.Println("noOfUsers variable value (using walurs operator) : ", noOfUsers)
	fmt.Printf("type of the noOfUsers %T \n", noOfUsers)

	fmt.Println("Const variable PI value : ", pi)
}

//Result

// Variables
// WELECOME TO THE GO LANG WORLD :  Umesh
// Variable is of Type  string
// 255
// Variable is of type : uint
// 355.91
// Variable is of type : float32
// true
// Variable is of type : bool
// Default value of int 0
// Variable is of type : int
// Default value of string
// Variable is of type : string
// Default value of bool false
// Variable is of type : bool
// Default value of float32 0
// Variable is of type : float32
// Default value of float64 0
// Variable is of type : float64
