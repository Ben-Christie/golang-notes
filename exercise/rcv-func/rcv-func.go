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

func (player *Player) increaseHealth() {
	player.health += 10

	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println(player.name, "has the following health points remaining: ", player.health)
}

func (player *Player) decreaseHealth() {
	player.health -= 10

	if player.health <= 0 {
		fmt.Println(player.name, "has died!")
	} else {
		fmt.Println(player.name, "has the following health points remaining: ", player.health)
	}
}

func (player *Player) increaseEnergy() {
	player.energy += 10

	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}

	fmt.Println(player.name, "has the following energy points remaining: ", player.energy)
}

func (player *Player) decreaseEnergy() {
	player.energy -= 10

	if player.energy <= 0 {
		fmt.Println(player.name, "is exhausted!")
	} else {
		fmt.Println(player.name, "has the following energy points remaining: ", player.energy)
	}
}

func main() {
	player1 := Player{"Ben", 100, 100, 100, 100}

	player1.decreaseEnergy()
	player1.decreaseEnergy()
	player1.decreaseEnergy()
	player1.increaseEnergy()
	fmt.Println()
	player1.decreaseHealth()
	player1.decreaseHealth()
	player1.increaseHealth()

}
