package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Operating System\t%v\n", runtime.GOOS)
	fmt.Printf("Architecture\t\t%v\n", runtime.GOARCH)
}
