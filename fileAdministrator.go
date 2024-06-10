package main

import (
	"bufio"
	"fmt"
	"inventoryApp/structs"
	"os"
	"strconv"
	"strings"
)

const filepath = "data/products.txt"

func readProducts() ([]structs.Product, error) {
	file, err := os.Open(filepath)
	printError(err)

	defer file.Close()

	var products []structs.Product
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}
		idProduct, err := strconv.ParseUint(fields[0], 10, 16)
		printError(err)

		name := fields[1]
		price, err := strconv.ParseFloat(fields[2], 64)
		printError(err)

		quantity, err := strconv.ParseUint(fields[3], 10, 16)
		printError(err)

		product := structs.Product{
			Id:       idProduct,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		products = append(products, product)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func lastProductIndex() int {
	products, err := readProducts()
	printError(err)
	lastProduct := products[len(products)-1]
	return int(lastProduct.Id)
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func showProductsOnScreen() {
	products, err := readProducts()
	fmt.Println("-- Products Uploaded --")
	for _, products := range products {
		fmt.Printf("Id: %d, Product: %s, Price: %0.2f, Quantity: %d\n", products.Id, products.Name, products.Price, products.Quantity)
	}
	printError(err)
}

func eliminateStock() {
	fmt.Println("Please enter the ID product to delete: ")
	var idToDelete string
	fmt.Scanf("%s", &idToDelete)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file... ", err)
		return
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file... ", err)
		return
	}

	var newLines []string

	for _, line := range lines {
		parts := strings.SplitN(line, ",", 2)
		if parts[0] != idToDelete {
			newLines = append(newLines, line)
		}
	}

	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file to modify... ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range newLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing file... ", err)
			return
		}
	}

	writer.Flush()

	fmt.Println("Product ID:", idToDelete, " eliminated.")
}

func addNewProduct() {
	showProductsOnScreen()
	fmt.Println("\n --- Enter a new Product ---")
	fmt.Println("Please insert the name of the new Product")

	reader := bufio.NewReader(os.Stdin)
	productName, err := reader.ReadString('\n')
	printError(err)

	productName = strings.TrimSpace(productName)

	fmt.Println("Please insert the price of the new Product")
	var price float32
	fmt.Scanf("%f\n", &price)

	fmt.Println("Please insert the quantity of the new Product")
	var quantity int
	fmt.Scanf("%d\n", &quantity)

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	printError(err)

	defer file.Close()

	id := lastProductIndex() + 1

	newProduct := fmt.Sprintf("%d,%s,%.2f,%d\n", id, productName, price, quantity)
	_, writeError := file.WriteString(newProduct)
	if writeError != nil {
		fmt.Println("error writing product: ", writeError)
	} else {
		fmt.Printf("Product %s added!\n", productName)
	}
}
