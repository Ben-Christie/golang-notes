package main

import "fmt"

func main() {
	// kind of an easier way of making a for loop with a counter

	slice := []string{"Hello", "World", "!"}

	// i starts at 0 in this case and element refers to the thing itself
	for i, element := range slice {
		fmt.Println(i, element)

		// allows us to iterate over individual letters in each string
		// underscore to ignore counter
		for _, letter := range element {
			// %q to print rune
			fmt.Printf("%q\n", letter)
		}
	}
}
