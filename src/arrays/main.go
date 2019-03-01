package main

import "fmt"

func main() {

	// Defines an array of Strings with max length of 3
	var langs [3]string

	langs[0] = "Go"
	langs[1] = "Ruby"
	langs[2] = "JavaScript"
	// langs[3] = "LOLcode" // Throws an error because index of 3 can't exist according to our definition above

	fmt.Println(langs)

	// Defines an array of strings with an undetermined length
	var slices []string

	slices = append(slices, "Go")
	slices = append(slices, "Ruby")
	slices = append(slices, "JavaScript")
	slices = append(slices, "LOLcode")

	fmt.Println(slices)

	sliceLiteral := []string{"Go", "Ruby", "JavaScript"}

	for _, element := range sliceLiteral {

		fmt.Println(element)

	}

	fmt.Println(sliceLiteral)

}
