package main

import "fmt"

func main() {
	states := []string{
		`Alabama`,
		`Alaska`,
		`Arizona`,
		`Arkansas`,
		`California`,
		`Colorado`,
		`Connecticut`,
		`Delaware`,
		`Florida`,
		`Georgia`,
		`Hawaii`,
		`Idaho`,
		`Illinois`,
		`Indiana`,
		`Iowa`,
		`Kansas`,
		`Kentucky`,
		`Louisiana`,
		`Maine`,
		`Maryland`,
		`Massachusetts`,
		`Michigan`,
		`Minnesota`,
		`Mississippi`,
		`Missouri`,
		`Montana`,
		`Nebraska`,
		`Nevada`,
		`New Hampshire`,
		`New Jersey`,
		`New Mexico`,
		`New York`,
		`North Carolina`,
		`North Dakota`,
		`Ohio`,
		`Oklahoma`,
		`Oregon`,
		`Pennsylvania`,
		`Rhode Island`,
		`South Carolina`,
		`South Dakota`,
		`Tennessee`,
		`Texas`,
		`Utah`,
		`Vermont`,
		`Virginia`,
		`Washington`,
		`West Virginia`,
		`Wisconsin`,
		`Wyoming`,
	}

	statesQty := len(states)
	statesCap := len(states)
	fmt.Println("len -", statesQty, "cap -", statesCap)

	for i := 0; i < statesQty; i++ {
		fmt.Println(i, states[i])
	}
}
