package application

import (
	"fmt"
	"testing"
	"time"

	"github.com/chrxn1c/pokemon-repl/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.value)
			val := cache.Get(c.key)
			if val == nil {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	val := cache.Get("https://example.com")
	if val == nil {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	val = cache.Get("https://example.com")
	if val != nil {
		t.Errorf("expected to not find key")
		return
	}
}
