package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt++
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
