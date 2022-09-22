package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is", sum)

	// standard for loop
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	// while loop
	x := 0

	for x < 10 {
		fmt.Print(x)
		x++
	}

	//infinite loop
	// for {
	// 	//write code here
	// }
}
