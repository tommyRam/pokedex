package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntryData map[string]cacheEntry
	mux            *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.cacheEntryData[key]; ok {
		return false, fmt.Errorf("Key: %s already exist", key)
	}

	c.cacheEntryData[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	return true, nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.cacheEntryData[key]; !ok {
		return nil, false
	}

	cacheEntryData, _ := c.cacheEntryData[key]
	return cacheEntryData.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer c.mux.Unlock()
		for {
			now := <-ticker.C
			
			for key, value := range(c.cacheEntryData) {
				if curr_interval := now.Sub(value.createdAt); curr_interval >= interval {
				c.mux.Lock()
				delete(c.cacheEntryData, key)
			}
		}
	}}()
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cacheEntryData: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	newCache.reapLoop(interval)
	return newCache
}