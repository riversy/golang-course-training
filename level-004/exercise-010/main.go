package main

import "fmt"

func main() {
	favors := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	favors[`yoda`] = []string{
		`Padawan`,
		`Jedi`,
	}

	if _, ok := favors[`no_dr`]; ok {
		delete(favors, `no_dr`)
	}

	for key, personFavors := range favors {
		fmt.Printf(">>> Favs of the %v <<<\n", key)
		for i, fav := range personFavors {
			fmt.Printf("\t- %v. %v\n", i+1, fav)
		}
	}
}
