package level1tasks

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu sync.Mutex
}

func (c *counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func Task18() {
	c := counter{}
	wg := sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}