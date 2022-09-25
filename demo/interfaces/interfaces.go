package main

import "fmt"

// naming convention for interfaces is usually a doer (ends in er)
type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (a Salad) PrepareDish() {
	fmt.Println("dress salad")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare Dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken Wings"),
		Salad("Cesar Salad"),
	}

	fmt.Println(dishes)
}
