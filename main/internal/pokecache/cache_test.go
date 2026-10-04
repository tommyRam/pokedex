package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		data []byte
	} {
		{
			key: "https://example.com",
			data: []byte("Test data"),
		}, 
		{
			key: "https://example.com/path",
			data: []byte("more test data"),
		},
	}

	for i, c := range cases {
		t.Run(
		fmt.Sprintf("Test case %v", i),
		func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.data)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.data) {
				t.Errorf("expected to find the same value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("example.com", []byte("Test data"))
	
	_, ok := cache.Get("example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("example.com")
	if ok {
		t.Errorf("expected to be not found")
		return
	}
}