package main

import (
	"fmt"
	"math" // package
)

func main() {
	// fmt.Println(math.pi)
	// When importing a package you can refer only to its exported names
	// Any unexported names are not accessible from outside the package.

	// To fix the error rename math.pi to math.Pi
	fmt.Println(math.Pi) // exported name
}
