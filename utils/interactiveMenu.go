package utils

import (
	"fmt"
)

func Menu() (int, error) {

	ClearConsole()
	fmt.Println("================================")
	fmt.Println("	  Welcome to inventory")
	fmt.Println("================================")
	fmt.Println("Choose your option.")
	fmt.Println()

	fmt.Println(" 1 - See all available products")
	fmt.Println(" 2 - Add new product ")
	fmt.Println(" 3 - Eliminate existing product")
	fmt.Println()
	fmt.Println(" 0 - Exit program ")
	fmt.Println()

	fmt.Print("Enter a value between 1 to 3: ")
	var menuChoise int
	_, err := fmt.Scanln(&menuChoise)

	return menuChoise, err
}
