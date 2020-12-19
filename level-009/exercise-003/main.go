package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var counter int
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter", counter)
			wg.Done()
		}()
	}
	wg.Wait()
}
