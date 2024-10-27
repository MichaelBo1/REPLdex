package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
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
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find the key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("mismatch in values")
				return
			}
		})
	}
}

func TestPurgeLoop(t *testing.T) {
	base := 5 * time.Millisecond
	wait := base + 5*time.Millisecond

	cache := NewCache(base)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find the key")
		return
	}

	time.Sleep(wait)

	// Should now be purged
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("value should have been purged")
		return
	}
}
