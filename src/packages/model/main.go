package model

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func (g gopher) Jump() string {

	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"

}

type horse struct {
	name   string
	weight int
}

func (h horse) Jump() string {

	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	}
	return "I will jump, neigh!!!"

}

// Jumper : Adds calculated jump field using Jump() method
type Jumper interface {
	Jump() string
}

// GetList : Exports list of our gophers and horses to consume on another function
func GetList() []Jumper {

	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 30}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []Jumper{phil, noodles, barbaro}

	return list

}
