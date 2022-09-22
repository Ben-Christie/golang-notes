package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("----", title, "----")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Supermarket", "Butcher", "Bakery"}

	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	// start from index 2 to the end of the slice
	route = route[2:]
	printSlice("Remaining locations", route)
}
