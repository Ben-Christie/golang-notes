package main

import "fmt"

func main() {
	// 									 map[key]value
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	// count of bread will be 1 -> by default its 0 and then we'll add 1 to it with the += 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	// delete from a map
	delete(shoppingList, "milk")
	fmt.Println("Milk removed:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs") // 12, returns value

	fmt.Println()

	// check for the existence of something within the map
	// first var is the thing we're looking for and the found var is a bool that'll return true or false depending on if we
	// find something
	cereal, found := shoppingList["cereal"]
	fmt.Println("Do we need cereal?")
	if !found {
		fmt.Println("No")
	} else {
		fmt.Println("Yes we need", cereal, "boxes")
	}

	// quantity -> 13
	totalItems := 0

	// key, value -> here we're using _ to ignore the keys
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("We need to buy", totalItems, "things")
}
