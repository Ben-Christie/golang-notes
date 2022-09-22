// allows us to have the main func below
package main

// fmt is the format package and allows us to print out information
import "fmt"

func main() {
	var myName = "Ben"
	// println automatically adds a space
	fmt.Println("my name is", myName)

	// this will just print out the myName variable twice with a spce inbetween
	fmt.Println("my name is", myName, myName)

	// Type Anotation (define the type)

	var name string = "Alex"
	fmt.Println("name =", name)

	// Create and Assign (:=)

	username := "admin"
	fmt.Println("username =", username)

	// unitialised variable

	var sum int
	fmt.Println("sum =", sum)

	// multi assignment

	part1, part2 := 2, 4
	fmt.Println("sum =", part1+part2) // 2 + 4 = 6

	// comma ok idiom - allows us to recreate a value when we're using commas
	// below we'll recreate part2 even though we've already created it above
	part2, part3 := 1, 3
	fmt.Println("sum =", part2+part3) // 1 + 3 = 4

	// block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("Today we're studying", lessonName, "we completed a", lessonType)

	// exclamation mark will be ignored (denoted by underscore)
	word1, word2, _ := "hello", "world", "!"

	fmt.Println(word1, word2)
}
