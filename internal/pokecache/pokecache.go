package pokecache

import (
	"context"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	mutex                 *sync.RWMutex
	content               map[string]cacheEntry
	destroyValueTimestamp time.Duration
}

func NewCache(destroyValueTimestamp time.Duration) (cache *Cache) {
	return &Cache{
		mutex:                 &sync.RWMutex{},
		content:               map[string]cacheEntry{},
		destroyValueTimestamp: destroyValueTimestamp,
	}
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), cache.destroyValueTimestamp)
		defer cancel()
		<-ctx.Done()
		cache.mutex.Lock()
		delete(cache.content, key)
		cache.mutex.Unlock()
	}()

	cache.content[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (cache *Cache) Get(key string) (value []byte) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	valueByKey, ok := cache.content[key]
	if !ok {
		return nil
	}

	return valueByKey.value
}
