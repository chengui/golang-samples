package cache

import (
	"fmt"
	"container/list"
)

type Node struct {
	Key string
	Value string
}

type LRUCache struct {
	Index map[string]*list.Element
	Cache *list.List
	Size int
}

func NewLRUCache(size int) *LRUCache {
	cache := list.New()
	index := make(map[string]*list.Element, size)
	return &LRUCache{
		Cache: cache,
		Index: index,
		Size: size,
	}
}

func (c *LRUCache) Get(key string) (string, error) {
	idx, ok := c.Index[key]
	if ok != true {
		return "", fmt.Errorf("item not found")
	}
	c.Cache.MoveToFront(idx)
	elem := c.Cache.Front()
	node := elem.Value.(Node)
	return node.Value, nil
}

func (c *LRUCache) Set(key string, val string) error {
	idx, ok := c.Index[key]
	if ok {
		node := idx.Value.(Node)
		node.Value = val
		c.Cache.MoveToFront(idx)
		return nil
	}
	if len(c.Index) >= c.Size {
		idel := c.Cache.Back()
		node := idel.Value.(Node)
		kdel := node.Key
		c.Cache.Remove(idel)
		delete(c.Index, kdel)
	}
	node := Node{
		Key: key,
		Value: val,
	}
	c.Index[key] = c.Cache.PushFront(node)
	return nil
}

func (c *LRUCache) Delete(key string) error {
	idx, ok := c.Index[key]
	if ok != true {
		return fmt.Errorf("item not found")
	}
	c.Cache.Remove(idx)
	return nil
}

func (c *LRUCache) Len() int {
	return c.Cache.Len()
}
