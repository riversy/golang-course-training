package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Ray",
		lastName:  "Britten",
		favoriteIceCreamFlavors: []string{
			"Gelato",
			"Mochi Ice Cream",
		},
	}

	p2 := person{
		firstName: "Corbett",
		lastName:  "Winley",
		favoriteIceCreamFlavors: []string{
			"Sorbet",
		},
	}

	for _, p := range []person{p1, p2} {
		fmt.Printf(">>> %v %v <<<\n", p.firstName, p.lastName)
		fmt.Printf("Favourite Ice Cream Flavors:\n")
		for _, flavour := range p.favoriteIceCreamFlavors {
			fmt.Printf("\t- %v\n", flavour)
		}
	}
}
