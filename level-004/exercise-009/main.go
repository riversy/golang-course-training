package main

import "fmt"

func main() {
	favs := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	favs[`yoda`] = []string{
		`Padawan`,
		`Jedi`,
	}

	for key, personFavs := range favs {
		fmt.Printf(">>> Favs of the %v <<<\n", key)
		for i, fav := range personFavs {
			fmt.Printf("\t- %v. %v\n", i+1, fav)
		}
	}
}
