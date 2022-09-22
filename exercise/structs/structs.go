//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func height(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

func perimeter(width, height int) int {
	return (width * 2) + (height * 2)
}

func area(width, height int) int {
	return width * height
}

func doubleSize(rect Rectangle) Rectangle {
	rect.a.y += height(rect)
	rect.b.x += width(rect)

	return rect
}

func main() {
	rectangle := Rectangle{Coordinate{0, 7}, Coordinate{10, 0}}

	fmt.Println(rectangle)
	fmt.Println("The perimeter of the rectangle is:", perimeter(width(rectangle), height(rectangle)))
	fmt.Println("The area of the rectangle is:", area(width(rectangle), height(rectangle)))

	rectangle = doubleSize(rectangle)

	fmt.Println(rectangle)
	fmt.Println("The perimeter of the doubled rectangle is:", perimeter(width(rectangle), height(rectangle)))
	fmt.Println("The area of the doubled rectangle is:", area(width(rectangle), height(rectangle)))

}
