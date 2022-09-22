package main

// pass by value -> whenever you call a function, a copy of each argument is made, regardless of how large it is
// this can make things very slow when dealing with large amounts of data
// pointers allow us to change this default pass by value method by giving direct access instead of creating a new copy

import "fmt"

type Counter struct {
	hits int
}

// takes a pointer to a counter
func increment(counter *Counter) {
	// don't need to use * infront of counter as its a structure and uses the (.) notation which will auto destructure
	counter.hits += 1
	fmt.Println("Current Count:", counter)
}

// replaces a string with something else
func replace(old *string, new string, counter *Counter) {
	// use * to access old string
	*old = new
	increment(counter)
}

func main() {
	// left empty to default to 0 hits
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	// the pointer to the hello var, the new string and the pointer to the counter
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}
