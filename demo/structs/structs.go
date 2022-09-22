package main

import "fmt"

// structures are similar to classes in other programming languages

// creating a structure
type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println("Casey's Ticket\n", casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2} // boarded value not specified so set to default i.e. false
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println("Bill's Ticket\n", bill, "\nElla's Ticket\n", ella)

	// uninitialised variable
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println("Heidi's Ticket\n", heidi)

	// update fields
	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}

	if casey.Boarded {
		fmt.Println("Casey has boarded the bus")
	}

	heidi.Boarded = true

	// create new bus
	bus := Bus{heidi}
	fmt.Println("Bus\n", bus)

	fmt.Println(bus.FrontSeat.Name, "is in the front seat of the bus")
}
