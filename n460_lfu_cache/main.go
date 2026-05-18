package main

import "container/list"

func New(capacity int) LFUCache {
	return LFUCache{
		len:     0,
		cap:     capacity,
		minFreq: 0,
		buckets: make(map[int]*list.List),
		hmap:    make(map[int]element),
	}
}

type LFUCache struct {
	len     int
	cap     int
	minFreq int
	buckets map[int]*list.List // map[freq]lru
	hmap    map[int]element
}

type element struct {
	freq int
	elem *list.Element
}

type elementValue struct {
	key, val int
}

func Constructor(capacity int) LFUCache {
	return New(capacity)
}

func (c *LFUCache) Get(key int) int {
	el, ok := c.hmap[key]
	if !ok {
		return -1
	}

	oldBucket := c.buckets[el.freq]
	oldBucket.Remove(el.elem)

	if el.freq == c.minFreq && oldBucket.Len() == 0 {
		c.minFreq++
	}

	el.freq++

	newBucket, ok := c.buckets[el.freq]
	if !ok {
		newBucket = list.New()
	}

	el.elem = newBucket.PushBack(el.elem.Value)
	c.buckets[el.freq] = newBucket
	c.hmap[key] = el

	return el.elem.Value.(elementValue).val
}

func (c *LFUCache) Put(key int, value int) {
	el, ok := c.hmap[key]
	if ok {
		oldBucket := c.buckets[el.freq]
		oldBucket.Remove(el.elem)
		if el.freq == c.minFreq && oldBucket.Len() == 0 {
			c.minFreq++
		}

		el.freq++
	} else {
		if c.len == c.cap {
			c.evict()
		}

		c.minFreq = 1
		el.freq = 1
		c.len++
	}

	newBucket, ok := c.buckets[el.freq]
	if !ok {
		newBucket = list.New()
	}

	el.elem = newBucket.PushBack(elementValue{key: key, val: value})
	c.buckets[el.freq] = newBucket
	c.hmap[key] = el
}

func (c *LFUCache) evict() {
	buck, ok := c.buckets[c.minFreq]
	if !ok {
		return
	}

	el := buck.Front()
	if el == nil {
		return
	}

	buck.Remove(el)
	delete(c.hmap, el.Value.(elementValue).key)
	c.len--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
