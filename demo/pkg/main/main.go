package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

// this will be the package that launches the program

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}
