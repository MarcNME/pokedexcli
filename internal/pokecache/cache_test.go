package pokecache

import (
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	cache := NewCache(1 * time.Hour)

	cases := []struct {
		key string
		val string
	}{
		{
			key: "Test",
			val: "Hallo, Welt!",
		},
		{
			key: "1234",
			val: "Hallo, Welt!\n Etwas länger👀",
		},
	}

	for _, c := range cases {
		cache.Add(c.key, []byte(c.val))

		val, ok := cache.Get(c.key)

		if !ok {
			t.Error("Added key not found")
		}

		if string(val) != c.val {
			t.Errorf("Expected %s, got %s", c.val, val)
		}
	}

}

func TestCache_ReapLoop(t *testing.T) {
	cache := NewCache(1 * time.Second)

	cache.Add("Test", []byte("Hallo, Welt!"))

	time.Sleep(2 * time.Second)

	_, ok := cache.Get("Test")

	if ok {
		t.Error("Cache reap loop didn't work")
	}
}
