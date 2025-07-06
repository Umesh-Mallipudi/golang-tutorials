package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcomeMsg := "Hello Developers :)"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your Experience : ")

	input, _ := reader.ReadString('\n')

	score, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		walletBalance := score * 100
		fmt.Println("YOUR Wallet Balance : ", walletBalance)
	}
}
