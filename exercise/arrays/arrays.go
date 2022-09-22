//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func printStats(list [4]Product) {
	var (
		cost       float64
		totalItems int
	)

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems++
		}
	}

	fmt.Println("The last item on the list is", list[totalItems-1])
	fmt.Println("The total nuymber of items on the list is", totalItems)
	fmt.Println("The total cost is", cost)
}

func main() {
	shoppingList := [4]Product{
		{name: "milk", price: 1.69},
		{name: "bread", price: 3.25},
		{name: "eggs", price: 2.50},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{"banana", 0.79}

	printStats(shoppingList)
}
