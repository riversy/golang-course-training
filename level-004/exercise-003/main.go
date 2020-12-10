package main

import "fmt"

func main() {
	x := [][]int{
		{42, 43, 44, 45, 46},
		{47, 48, 49, 50, 51},
		{44, 45, 46, 47, 48},
		{43, 44, 45, 46, 47},
	}
	for i, v := range x {
		fmt.Println(i, v, len(v), cap(v))
	}

	fmt.Printf("%T\t\t%v\t%v\n", x, len(x), cap(x))
	fmt.Printf("%T\t\t%v\t%v\n", x, len(x), cap(x))
}
