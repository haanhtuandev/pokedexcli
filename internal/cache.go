package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (*Cache, error) {
	new_cache := &Cache{
		CacheMap: make(map[string]cacheEntry),
		interval: interval,
	}
	go new_cache.reapLoop()
	return new_cache, nil
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

	return nil
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.CacheMap[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		for k, entry := range c.CacheMap {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.CacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}
