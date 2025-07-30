package level1tasks

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu sync.Mutex
	m map[int]int 
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[int]int),
	}
}

func (c *ConcurrentMap) Set(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *ConcurrentMap) Get(key int) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func Task7() {
	m := NewConcurrentMap()
	wg := sync.WaitGroup{}

	// Запись
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Set(i, i*10) // записываем число в 10 раз больше
		}()
	}

	// Чтение
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if val, ok := m.Get(i); ok {
				fmt.Printf("Key is %d, value is %d\n", i, val)
			} else {
				fmt.Printf("Key %d doesn't have value yet\n", i)
			}
		}(i)
	}

	wg.Wait()
}