//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type SecurityTag struct {
	name   string
	status bool
}

func activate(tag *SecurityTag) {
	tag.status = true
}

func deactivate(tag *SecurityTag) {
	tag.status = false
}

func checkout(basket []SecurityTag) {
	// iterate through basket and deactivate
	for i := range basket {
		deactivate(&basket[i])
	}
}

func main() {
	tShirt := SecurityTag{"t-shirt", true}
	jumper := SecurityTag{"jumper", true}
	jeans := SecurityTag{"jeans", true}
	socks := SecurityTag{"socks", true}

	basket := []SecurityTag{tShirt, jumper, socks, jeans}

	fmt.Println(basket)

	deactivate(&basket[1])

	fmt.Println(basket)

	checkout(basket)

	fmt.Println(basket)
}
