package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// standard function
func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// the same function as above but as a receiver function
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// receiver function for when a space becomes vacant
func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println("Initial:", lot)

	// using receiver function
	lot.occupySpace(1)

	// using normal function
	occupySpace(&lot, 2)

	fmt.Println("After occupied:", lot)

	lot.vacateSpace(2)

	fmt.Println("After vacate:", lot)
}
