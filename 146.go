package main

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Get(key int) int
	Put(key int, value int)
	Size() int
}

type LRUCache struct {
	cap        int
	keyMap     map[int]*list.Element
	linkedList *list.List
}

type node struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:        capacity,
		keyMap:     make(map[int]*list.Element, capacity),
		linkedList: new(list.List),
	}
}

func (c *LRUCache) geNodeFromElement(elem *list.Element) *node {
	switch v := elem.Value.(type) {
	case *node:
		return v
	default:
		panic("unreachable")
	}
}
func (c *LRUCache) Get(key int) int {
	if link, ok := c.keyMap[key]; ok {
		c.linkedList.MoveToFront(link)
		return c.geNodeFromElement(link).value

	}
	return -1
}

func (c *LRUCache) extractLatest() {
	del := c.linkedList.Back()
	c.linkedList.Remove(del)
	delete(c.keyMap, c.geNodeFromElement(del).key)
}

func (c *LRUCache) Put(key int, value int) {
	if link, ok := c.keyMap[key]; ok {
		n := c.geNodeFromElement(link)
		n.value = value
		c.linkedList.MoveToFront(link)
		return
	}

	if c.Size() == c.cap {
		c.extractLatest()
	}

	n := &node{
		key:   key,
		value: value,
	}
	c.keyMap[key] = c.linkedList.PushFront(n)
}

func (c *LRUCache) Size() int {
	return len(c.keyMap)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	capacity := 3
	obj := Constructor(capacity)
	fmt.Println(obj)
	key := 1
	value := 100
	obj.Put(key, value)
	fmt.Println(obj)
	param1 := obj.Get(key)
	fmt.Println(param1)
}
