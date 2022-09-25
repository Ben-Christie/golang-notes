package display

// can't use a main func in here

import "fmt"

// function with capital letter means its going to be exported
// public and private in Go is based on capitalisation, capital letter = public, lower case letter = private
func Display(msg string) {
	fmt.Println(msg)
}
