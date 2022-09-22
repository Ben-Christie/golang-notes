package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println(dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println(bakersDozen)

	// nested functions
	numTentacles := add(double(2), 4)
	fmt.Println("An octopus has", numTentacles, "tentacles")
}
