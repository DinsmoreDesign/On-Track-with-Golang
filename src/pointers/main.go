package main

import "fmt"

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func validateAge(g *gopher) { // "*" indicates a pointer to the original version

	g.isAdult = g.age >= 21

}

func main() {

	phil := &gopher{name: "Phil", age: 35} // Not using "&" creates a copy, not a reference, so this will not actually change our original data structure

	validateAge(phil)

	fmt.Println(phil) // Prints &{Phil 30 true}

}
