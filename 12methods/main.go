package main

import "fmt"

// func main() {

// 	fmt.Println("Welecome to the Golang world with Structs and methods :)")

// 	p := Person{"John", 25, "Single", 1234567890}
// 	p.PrintPerson()
// }

// type Person struct {
// 	Name     string
// 	Age      int
// 	Status   string
// 	PhoneNum int
// }

// func (p Person) PrintPerson() {
// 	fmt.Println("Name:", p.Name)
// 	fmt.Println("Age:", p.Age)
// 	fmt.Println("Status:", p.Status)
// 	fmt.Println("Phone Number:", p.PhoneNum)
// }

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) Describe() {
	fmt.Println("Book Info: ")
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.Pages)
}

func (b *Book) AddPages(pages int) {
	b.Pages += pages
	fmt.Println("Book Pages:", b.Pages)
}
func main() {
	myBook := Book{"Golang", "John", 100}
	myBook.Describe()
	myBook.AddPages(10099)
	myBook.Describe()

}
