//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayStatus(servers map[string]int) {
	// how many servers are there?
	fmt.Println("\nThere are", len(servers), "servers")

	// how many servers under each state?
	// the key will be the corresponding value for each of the status codes and the valuew will be the number of values with
	// that status code
	states := make(map[int]int)

	// iterate over map
	for _, status := range servers {
		switch status {
		case Online:
			states[Online] += 1
		case Offline:
			states[Offline] += 1
		case Maintenance:
			states[Maintenance] += 1
		case Retired:
			states[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println(states[Online], "servers are online")
	fmt.Println(states[Offline], "servers are offline")
	fmt.Println(states[Maintenance], "servers are undergoing maintenance")
	fmt.Println(states[Retired], "servers are retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	// make empty map -> map[key]value
	serverStatus := make(map[string]int)

	// iterate over servers and set status to online
	for _, server := range servers {
		serverStatus[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	displayStatus(serverStatus)
	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	//  - call display server info function
	displayStatus(serverStatus)
	//  - change server status of all servers to `Maintenance`
	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	//  - call display server info function
	displayStatus(serverStatus)
}
