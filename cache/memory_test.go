package cache

import (
	"sync"
	"testing"
	"time"
)

func TestMemoryCache_SetGet(t *testing.T) {
	cache := NewMemoryCache()
	key, value := "test_key", "test_value"

	// Test Set and Get
	cache.Set(key, value, time.Hour)
	if v, ok := cache.Get(key); !ok || v != value {
		t.Errorf("Get failed for key %s, expected %s, got %v, ok %v", key, value, v, ok)
	}

	// Test non-existing key
	if _, ok := cache.Get("non_existing_key"); ok {
		t.Error("Get returned true for non-existing key")
	}
}

func TestMemoryCache_Expiration(t *testing.T) {
	cache := NewMemoryCache()
	key, value := "expiring_key", "expiring_value"

	// Set with a very short duration to test expiration
	cache.Set(key, value, time.Millisecond) // Should expire very quickly
	time.Sleep(10 * time.Millisecond)       // Wait a bit to ensure expiration

	if _, ok := cache.Get(key); ok {
		t.Error("Item did not expire as expected")
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache()
	key, value := "delete_key", "delete_value"

	// Set and then delete
	cache.Set(key, value, time.Hour)
	cache.Delete(key)

	if _, ok := cache.Get(key); ok {
		t.Error("Delete did not remove the item from cache")
	}
}

func TestMemoryCache_Clear(t *testing.T) {
	cache := NewMemoryCache()
	keys := []string{"key1", "key2", "key3"}
	values := []string{"value1", "value2", "value3"}
	for i, k := range keys {
		cache.Set(k, values[i], time.Hour)
	}

	cache.Clear()

	for _, k := range keys {
		if _, ok := cache.Get(k); ok {
			t.Errorf("Clear did not remove item with key %s from cache", k)
		}
	}
}

func TestMemoryCache_ConcurrentAccess(t *testing.T) {
	cache := NewMemoryCache()
	key, value := "thread_safe_key", "thread_safe_value"

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(key, value, time.Hour)
			if v, ok := cache.Get(key); !ok || v != value {
				t.Errorf("Concurrent get failed: expected %s, got %v, ok %v", value, v, ok)
			}
		}()
	}
	wg.Wait()

	// Check if the last set value is still there after concurrent operations
	if v, ok := cache.Get(key); !ok || v != value {
		t.Errorf("Final get after concurrent operations failed: expected %s, got %v, ok %v", value, v, ok)
	}
}
