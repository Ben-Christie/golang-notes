//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func message() string {
	return "This is a message"
}

func add(a, b, c int) int {
	return a + b + c
}

func number() int {
	return 2
}

func twoNumbers() (int, int) {
	return 2, 3
}

func main() {
	fmt.Println(greet("Ben"))

	fmt.Println(message())

	answer := add(1, 2, 3)
	fmt.Println(answer)

	fmt.Println(number())

	fmt.Println(twoNumbers())

	x, y := twoNumbers()

	fmt.Println(add(x, y, 7))
}
