package main

import (
	"fmt"
	"sync"
)

// V will be a {V comparable, ptr *Node} type
type Cache[K comparable, V any] struct {
	cache    map[K]V
	size     int // bytes
	max_size int
	mu       sync.Mutex
}

func NewCache[K comparable, V any](max_size int) *Cache[K, V] {
	c := Cache[K, V]{cache: make(map[K]V), size: 0, max_size: max_size}

	return &c
}

func (c *Cache[K, V]) MaxSize() int {
	return c.max_size
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.cache[k]

	return v, ok
}

func (c *Cache[K, V]) Set(k K, v V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[k] = v
}

func main() {
	c := NewCache[int, int](64)
	fmt.Println(c.MaxSize())
	v, ok := c.Get(1)
	fmt.Println(v, ok)
	c.Set(1, 1)
	v, ok = c.Get(1)
	fmt.Println(v, ok)

}
