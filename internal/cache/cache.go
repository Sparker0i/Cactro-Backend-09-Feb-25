package cache

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu      sync.RWMutex
	data    map[string]string
	maxSize int
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.data[key]; exists {
		c.data[key] = value
		return nil
	}

	if len(c.data) >= c.maxSize {
		return fmt.Errorf("cache is full: maximum of %d item size reached", c.maxSize)
	}

	c.data[key] = value
	return nil
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
