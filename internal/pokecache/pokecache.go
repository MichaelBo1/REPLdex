package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheElement struct {
	created time.Time
	value   []byte
}

// Mutex needed as maps are not concurrency-safe by default.
type Cache struct {
	cache map[string]cacheElement
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheElement{},
		mu:    &sync.Mutex{},
	}

	go c.purgeLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	fmt.Printf("Adding key: %s", key)
	c.mu.Lock()
	c.cache[key] = cacheElement{
		created: time.Now().UTC(),
		value:   value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("Getting key: %s", key)
	c.mu.Lock()
	elem, exists := c.cache[key]
	c.mu.Unlock()
	return elem.value, exists
}

// Tick checks every discardInterval, anything existing beyond that time (i.e., created before) is discarded
func (c *Cache) purgeLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.purge(time.Now().UTC(), interval)
	}
}

func (c *Cache) purge(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.created.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
