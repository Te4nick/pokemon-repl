package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	mutex                 sync.Mutex
	content               map[string]cacheEntry
	destroyValueTimestamp time.Duration
}

func (cache *Cache) NewCache(destroyValueTimestamp time.Duration) {
	cache.destroyValueTimestamp = destroyValueTimestamp
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mutex.Lock()

	defer cache.destroyRottenValue(key)
	defer cache.mutex.Unlock()

	cache.content[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value}
}

func (cache *Cache) Get(key string) (value []byte, isFound bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	valueByKey, isFound := cache.content[key]

	if !isFound {
		return []byte{}, false
	}

	return valueByKey.value, true
}

func (cache *Cache) destroyRottenValue(key string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	timer := time.NewTimer(cache.destroyValueTimestamp)
	go func() {
		<-timer.C
		delete(cache.content, key)
	}()
}
