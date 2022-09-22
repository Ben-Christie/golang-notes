//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(title string, line []Part) {
	fmt.Println("----", title, "----")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	fmt.Println()
}

func main() {
	// to make an empty slice with a specified length use the make function
	// assemblyLine := make([]Part, 3) -> creates a slice of parts of length 3
	assemblyLine := []Part{"Chassis", "Door", "Axles"}

	printAssemblyLine("Initial Assembly Line", assemblyLine)

	assemblyLine = append(assemblyLine, "Bonnet", "Exhaust")

	printAssemblyLine("Updated Assembly Line with 5 items", assemblyLine)

	assemblyLine = assemblyLine[3:]

	printAssemblyLine("Updated Assembly Line with 2 items", assemblyLine)

}
