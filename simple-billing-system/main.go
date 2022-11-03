package main

import "fmt"

var m = make(map[string]float64)

func main() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your adress:")
	var adress string
	fmt.Scanln(&adress)

	fmt.Println("Enter total items:")
	var items int
	fmt.Scanln(&items)

	var total float64
	var product string
	var price float64
	for i := 0; i <= items-1; i++ {
		fmt.Println("Enter the product: ")
		fmt.Scanln(&product)
		fmt.Printf("Enter the price $ of %v:\n", product)
		fmt.Scanln(&price)

		total += price

		if items == i {
			break
		}
		m[product] = price
	}
	fmt.Println()
	fmt.Println("		=====Bill Information=====		")
	fmt.Println()
	fmt.Printf("Customer's name: %v\nCustomer's adress: %v\n", name, adress)
	fmt.Println("----------------------------")
	for k, v := range m {
		fmt.Println("Product:", k, "| price:", v)
	}
	fmt.Println("----------------------------")
	fmt.Printf("Total price is %v$\n", total)
	fmt.Println("Thank you for visiting!")
}
