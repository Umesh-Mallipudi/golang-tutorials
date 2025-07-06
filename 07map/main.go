package main

import "fmt"

func main() {
	fmt.Println("Map's in GOlang")

	var citys = make(map[string]string)
	citys["Ap"] = "AndhraPradesh"
	citys["hyd"] = "Hyderabad"
	citys["bgl"] = "Benguluru"

	fmt.Println("List of citys : ", citys)
	fmt.Printf("AP stands for %v \n", citys["Ap"])
}
