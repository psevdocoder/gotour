package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    *sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	fmt.Println(c.count)
}

func main() {
	var wg sync.WaitGroup

	c := counter{mu: new(sync.Mutex)}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.inc()
			wg.Done()
		}()
	}

	wg.Wait()
}
