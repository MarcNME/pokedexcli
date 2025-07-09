package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mutex:   &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	val, ok := cache.entries[key]

	return val.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	for {
		cache.mutex.Lock()
		for k, v := range cache.entries {
			removeEntry := v.createdAt.Before(time.Now().Add(interval))
			if removeEntry {
				delete(cache.entries, k)
			}
		}
		cache.mutex.Unlock()
		time.Sleep(interval)
	}
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
