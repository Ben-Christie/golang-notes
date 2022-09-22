package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz 1 was higher than quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz 1 was lower than quiz 2")
	} else {
		fmt.Println("quiz 1 was equal to quiz 2")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("unacceptable grades")
	}
}
