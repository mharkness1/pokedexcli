package pokecache

import (
	"sync"
	"time"
)

// Definitions of entry and map that takes url -> GET response.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

// Creates a new cache needs use.
func NewCache(interval time.Duration) (Cache, error) {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	return c, nil
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().Local(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}
