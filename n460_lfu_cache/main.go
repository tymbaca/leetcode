package main

import "container/list"

func New(capacity int) LFUCache {
	return LFUCache{
		cap: capacity,
	}
}

type LFUCache struct {
	cap     int
	buckets map[int]*list.List // map[freq]lru
}

func Constructor(capacity int) LFUCache {
	return New(capacity)
}

func (c *LFUCache) Get(key int) int {
}

func (c *LFUCache) Put(key int, value int) {
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
