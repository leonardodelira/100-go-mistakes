package main

import (
	"fmt"
	"sync"
)

func main() {
	i := 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		i = 1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		i = 2
	}()

	wg.Wait()
	fmt.Print(i)
}
