package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	isLessThanFive := true

	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}

	i = 0

	for {
		fmt.Println(i)
		if i >= 5 {
			break
		}
		i++
	}

}
