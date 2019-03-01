package main

import "fmt"

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func (g gopher) jump() string {

	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"

}

type horse struct {
	name   string
	weight int
}

func (h horse) jump() string {

	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	}
	return "I will jump, neigh!!!"

}

type jumper interface {
	jump() string
}

func getList() []jumper {

	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 30}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}

	return list

}

func main() {

	jumperList := getList()

	for _, jumper := range jumperList {

		fmt.Println(jumper.jump())

	}

}
