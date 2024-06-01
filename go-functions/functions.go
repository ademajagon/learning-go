// A function can take zero or more arguments
// In this example , add function takes two parameters of type int
// Type comes after the variable name

package main

import (
	"fmt"
)

// Takes two parameters, type comes after the variable name
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type
// You can omit the type from all but the last
// So, we shortened
func subtract(x, y int) int {
	return x - y
}

// A function can return any number of results
// The swap function returns two strings
func swap(x, y string) (string, string) {
	return x, y
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the func
// These names should be used to define the meaning of the return value
// This is known as a "naked" return
// Naked return statements should only be used in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments return the named return values.

	//var a = 30
	//var b = 20
	//return a, b
	return
}

func main() {
	fmt.Println(add(3, 5))
	fmt.Println(subtract(8, 2))

	a, b := swap("hello", "world")
	// := This is a short variable declaration in Go
	// This operator is a shorthand that declares new variables and assigns them values
	// In this case a and b are being declared and assigned the values that are returned by the func

	fmt.Println(a, b)

	fmt.Println(split(17))
}
