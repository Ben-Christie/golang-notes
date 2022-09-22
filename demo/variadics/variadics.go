package main

import "fmt"

// variadics allow us to write fuunctions with an undefined number of parameters

// 3 dots indicates that we have a variadic, we can have any number of integers listed
// nums in this case is a slice of ints
func sum(nums ...int) int {
	sum := 0

	for _, number := range nums {
		sum += number
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	// the below is the same as this: answer = sum(1, 2, 3, 4, 5, 6)
	answer := sum(all...)
	fmt.Println(answer)
}
