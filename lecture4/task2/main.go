package main

import (
	"fmt"
	"sync"
)

func main() {
	mp := sync.Map{}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mp.Store(id, 1)
		}(i)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			val, ok := mp.Load(id)
			if ok {
				fmt.Println(id, val)
			}
		}(i)
	}

	wg.Wait()
}

func main2() {
	mp := make(map[int]int)

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			mp[id] = 1
			mu.Unlock()
		}(i)
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val, ok := mp[id]
			if ok {
				fmt.Println(id, val)
			}
		}(i)
	}

	wg.Wait()
}
