package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{3, 4, 5} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			c <- num
		}
		close(c)
	}()

	for num := range joinedChan(a, b, c) {
		fmt.Println(num)
	}
}

func joinedChan(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int) {
				defer wg.Done()
				for num := range ch {
					mergedCh <- num
				}
			}(ch)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}
