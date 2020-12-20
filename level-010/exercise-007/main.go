package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go func() {
		var wg sync.WaitGroup

		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				for i := 1; i <= 10; i++ {
					c <- i
				}
				wg.Done()
			}()
		}

		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
