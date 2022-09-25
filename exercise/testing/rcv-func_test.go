//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//

package main

import "testing"

// create player to be used in all the tests
func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
}

// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
func TestMaximums(t *testing.T) {
	player := newPlayer()

	player.increaseHealth(20)

	if player.health > player.maxHealth {
		t.Errorf("health should not be greater than the max health, got %v, expected %v", player.health, player.maxHealth)
	}

	player.increaseEnergy(20)

	if player.energy > player.maxEnergy {
		t.Errorf("energy should not be greater than the max energy, got %v, expected %v", player.energy, player.maxEnergy)
	}
}

// - Health & energy can not go below 0
func TestMinimums(t *testing.T) {
	player := newPlayer()

	player.decreaseHealth((player.health + 1))

	if player.health < 0 {
		t.Errorf("players health should not drop below 0, got %v", player.health)
	}

	player.decreaseEnergy((player.energy + 1))

	if player.energy < 0 {
		t.Errorf("players energy should not drop below 0, got %v", player.energy)
	}

}

//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
