//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts

//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

// define constants
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

// define interface
type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

// * Vehicles have a model name in addition to the vehicle type:
//   - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

// PickLift function definition for all types
func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

//   - Write a single function to handle all of the vehicles
//     that the shop works on.
func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Volkswagen Golf")
	truck := Truck("Chevrolet Silverado")
	motorcycle := Motorcycle("Yamaha 125 YBR")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}
