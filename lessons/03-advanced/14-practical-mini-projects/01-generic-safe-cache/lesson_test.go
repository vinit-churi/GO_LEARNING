package main

import "testing"

func TestCacheSetGet(t *testing.T) {
	cache := NewCache[string, int]()
	cache.Set("jobs", 3)
	got, ok := cache.Get("jobs")
	if !ok || got != 3 {
		t.Fatalf("Get() = %d, %v", got, ok)
	}
}

func TestCacheGetMissing(t *testing.T) {
	cache := NewCache[string, int]()
	if _, ok := cache.Get("missing"); ok {
		t.Fatal("Get(missing) = true")
	}
}

func TestCacheSize(t *testing.T) {
	cache := NewCache[string, int]()
	cache.Set("jobs", 3)
	cache.Set("workers", 2)
	if got := cache.Size(); got != 2 {
		t.Fatalf("Size() = %d", got)
	}
}
