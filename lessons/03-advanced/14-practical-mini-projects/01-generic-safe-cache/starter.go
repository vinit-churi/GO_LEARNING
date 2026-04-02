//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"sync"
)

type Cache[K comparable, V any] struct {
	mu     sync.Mutex
	values map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{values: map[K]V{}}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.values[key]
	return value, ok
}

func (c *Cache[K, V]) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.values)
}

func main() {
	cache := NewCache[string, int]()
	cache.Set("jobs", 3)
	fmt.Println(cache.Size())
}
