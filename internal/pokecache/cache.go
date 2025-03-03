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
