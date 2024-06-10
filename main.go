package main

import (
	"fmt"
	"inventoryApp/utils"
)

func main() {

	for {
		utils.ClearConsole()
		userChoise, err := utils.Menu()
		if err != nil {
			fmt.Println("Invalid option, enter a correct option.", err)
		}
		switch userChoise {
		case 1:
			utils.ClearConsole()
			showProductsOnScreen()
			pressAnyButtonToContinue()
		case 2:
			utils.ClearConsole()
			addNewProduct()
			pressAnyButtonToContinue()
		case 3:
			showProductsOnScreen()
			eliminateStock()
			pressAnyButtonToContinue()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, enter a correct option.")
		}
	}

}

func pressAnyButtonToContinue() {
	fmt.Println("Press any button to get back to menu...")
	fmt.Scanln()
	utils.ClearConsole()
}
