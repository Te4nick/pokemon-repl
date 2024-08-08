package cache

import (
	"context"
	"sync"
	"time"
)

const CacheDefaultTTL time.Duration = time.Minute * 5

type Cache struct {
	mutex   *sync.RWMutex
	content map[string][]byte
	ttl     time.Duration
}

func NewCache(ttl time.Duration) (cache *Cache) {
	return &Cache{
		mutex:   &sync.RWMutex{},
		content: map[string][]byte{},
		ttl:     ttl,
	}
}

func (cache *Cache) Set(key string, value []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), cache.ttl)
		defer cancel()
		<-ctx.Done()
		cache.mutex.Lock()
		delete(cache.content, key)
		cache.mutex.Unlock()
	}()

	cache.content[key] = value
}

func (cache *Cache) Get(key string) []byte {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	value, ok := cache.content[key]
	if !ok {
		return nil
	}

	return value
}
