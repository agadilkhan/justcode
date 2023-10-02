package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	m := make(map[int]int)

	mu := sync.RWMutex{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			m[id] = 1
			mu.Unlock()
		}(i)
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			val := m[id]
			mu.Unlock()
			fmt.Printf("key %d value %d\n", id, val)
		}(i)
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration)
}

func main1() {
	start := time.Now()
	m := sync.Map{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			m.Store(id, 1)

		}(i)
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			val, ok := m.Load(id)
			if ok {
				fmt.Printf("key %d value %d\n", id, val)
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration)
}
