package cache

import (
	"fmt"
	"container/list"
)

type Node struct {
	Key interface{}
	Val interface{}
}

type LRUCache struct {
	Index map[interface{}]*list.Element
	Cache *list.List
	Size int
}

func NewLRUCache(size int) *LRUCache {
	index := make(map[interface{}]*list.Element)
	cache := list.New()
	return &LRUCache{
		Index: index,
		Cache: cache,
		Size: size,
	}
}

func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	idx, ok := c.Index[key]
	if ok == false {
		return nil, false
	}
	c.Cache.MoveToFront(idx)
	node := idx.Value.(*Node)
	return node.Val, true
}

func (c *LRUCache) Set(key interface{}, val interface{}) error {
	idx, ok := c.Index[key]
	if ok {
		c.Cache.MoveToFront(idx)
		node := idx.Value.(*Node)
		node.Val = val
		return nil
	}
	if len(c.Index) >= c.Size {
		idel := c.Cache.Back()
		node := idel.Value.(*Node)
		delete(c.Index, node.Key)
		c.Cache.Remove(idel)
	}
	node := &Node{
		Key: key,
		Val: val,
	}
	c.Index[key] = c.Cache.PushFront(node)
	return nil
}

func (c *LRUCache) Delete(key interface{}) (interface{}, error) {
	idx, ok := c.Index[key]
	if ok == false {
		return nil, fmt.Errorf("item not found")
	}
	c.Cache.Remove(idx)
	node := idx.Value.(*Node)
	return node.Val, nil
}

func (c *LRUCache) Len() int {
	return c.Cache.Len()
}
