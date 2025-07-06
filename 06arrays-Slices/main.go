package main

import "fmt"

func main() {
	fmt.Println("Welecome to the go array class..")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "watermelon"
	fruitList[3] = "bannana"

	fmt.Println(fruitList)
	fmt.Println("len of the fruitList", len(fruitList))

	// slices

	vegList := []string{"Drumstick"}
	fmt.Println("Veglist ", vegList)
	ptr := &vegList
	fmt.Printf("Veglist variable memory address %p\n", ptr)

	vegList = append(vegList, "carrot")
	fmt.Println("Updated veglist : ", vegList)
	fmt.Printf("Updated veglist pointer %p\n", &vegList)
}
