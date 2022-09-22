package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

// only works on an array of length 4
func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "has been cleaned")
		} else {
			fmt.Println(room.name, "has not been cleaned")
		}
	}
}

func main() {
	// using the ... we don't need to specify a length, the computer will calculate how many values we've created and set that to the length, good if we want to make changes later (alternatively we could specify a number)
	rooms := [...]Room{
		{name: "dining room"},
		{name: "living room"},
		{name: "bedroom"},
		{name: "bathroom"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)
}
