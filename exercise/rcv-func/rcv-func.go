//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth int
	energy, maxEnergy int
}

func (player *Player) increaseHealth(amount int) {
	player.health += amount

	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println(player.name, "has the following health points remaining: ", player.health)
}

func (player *Player) decreaseHealth(amount int) {
	player.health -= amount

	if player.health <= 0 {
		fmt.Println(player.name, "has died!")
	} else {
		fmt.Println(player.name, "has the following health points remaining: ", player.health)
	}
}

func (player *Player) increaseEnergy(amount int) {
	player.energy += amount

	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}

	fmt.Println(player.name, "has the following energy points remaining: ", player.energy)
}

func (player *Player) decreaseEnergy(amount int) {
	player.energy -= amount

	if player.energy <= 0 {
		fmt.Println(player.name, "is exhausted!")
	} else {
		fmt.Println(player.name, "has the following energy points remaining: ", player.energy)
	}
}

func main() {
	player1 := Player{"Ben", 100, 100, 100, 100}

	player1.decreaseEnergy(30)
	player1.increaseEnergy(10)
	fmt.Println()
	player1.decreaseHealth(10)
	player1.increaseHealth(20)

}
