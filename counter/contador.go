package counter

import "sync"

type Contador struct {
	mu    sync.Mutex
	value int
}

func (c *Contador) Incrementing() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Contador) Value() int {
	return c.value
}
