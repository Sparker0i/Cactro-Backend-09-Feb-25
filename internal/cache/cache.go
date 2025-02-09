package cache

import (
	"fmt"
	"sync"
)

// Cache represents an in-memory key-value store with a fixed maximum capacity.
type Cache struct {
	mu      sync.RWMutex
	data    map[string]string
	maxSize int
}

// New creates and returns a new Cache instance with the given maximum size.
func New(maxSize int) *Cache {
	return &Cache{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

// Set inserts or updates a key-value pair in the cache.
// If the key is new and the cache is already full, it returns an error.
func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Update value if key exists.
	if _, exists := c.data[key]; exists {
		c.data[key] = value
		return nil
	}

	// Check if adding a new key would exceed max size.
	if len(c.data) >= c.maxSize {
		return fmt.Errorf("cache is full: maximum of %d items reached", c.maxSize)
	}

	c.data[key] = value
	return nil
}

// Get retrieves the value for a given key.
// Returns the value and true if found; otherwise, an empty string and false.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

// Delete removes the specified key from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
