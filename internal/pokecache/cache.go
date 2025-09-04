package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := make(map[string]cacheEntry)
	cacheNew := Cache{cache: cache, interval: interval}

	go cacheNew.reapLoop()

	return &cacheNew
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{time.Now(), value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.cache[key]
	if exists {
		return entry.val, true
	} else {
		fmt.Println("Key does not exist")
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mu.Lock()

		timeCurrent := time.Now()
		for key, entry := range c.cache {
			expire := entry.createdAt.Add(c.interval)
			if timeCurrent.After(expire) {
				delete(c.cache, key)
			} else {
				continue
			}

		}
		c.mu.Unlock()
	}

}
