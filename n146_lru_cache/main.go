package main

import "container/list"

func main() {
}

func New(capacity int) LRUCache {
	c := LRUCache{
		cap:  capacity,
		list: list.New(),
		hmap: make(map[int]*list.Element),
	}

	return c
}

type LRUCache struct {
	cap  int
	list *list.List // holds [node] as values
	hmap map[int]*list.Element
}

type node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return New(capacity)
}

func (c *LRUCache) Get(key int) int {
	el, ok := c.hmap[key]
	if !ok {
		return -1
	}

	c.list.MoveToBack(el)

	return el.Value.(node).val
}

func (c *LRUCache) Put(key int, value int) {
	el, ok := c.hmap[key]
	if ok {
		c.list.Remove(el)
	} else {
		if c.list.Len() == c.cap {
			c.evict()
		}
	}

	el = c.list.PushBack(node{
		key: key,
		val: value,
	})
	c.hmap[key] = el
}

func (c *LRUCache) evict() {
	el := c.list.Front()
	if el == nil {
		return
	}

	c.list.Remove(el)
	delete(c.hmap, el.Value.(node).key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
