package simple_cache

import (
	"testing"
)

var cache = NewSimpleCache()

func Test_cache(t *testing.T) {
	cache.Set("key", "value")

	val, ok := cache.Get("key")

	if !ok {
		t.Errorf("Expected ok to be true, got false")
	}

	if val != "value" {
		t.Errorf("Expected value to be 'value', got %s", val)
	}
}
