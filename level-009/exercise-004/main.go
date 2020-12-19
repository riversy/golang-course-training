package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("Counter", counter)
			mu.Unlock()

			fmt.Println("Goroutines", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
}
