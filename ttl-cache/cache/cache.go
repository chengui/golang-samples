package cache

import (
	"fmt"
	"time"
)

type Entry struct {
	Val interface{}
	TTL int
}

type TTLCache struct {
	Cache map[interface{}]*Entry
	TTL int
}

func NewTTLCache(ttl int) *TTLCache {
	cache := TTLCache{
		Cache: make(map[interface{}]*Entry),
		TTL: ttl,
	}
	go cache.clean()
	return &cache
}

func (c *TTLCache) Get(key interface{}) (interface{}, bool) {
	entry, ok := c.Cache[key]
	if ok == false {
		return nil, false
	}
	return entry.Val, true
}

func (c *TTLCache) Set(key interface{}, val interface{}) error {
	c.Cache[key] = &Entry{Val: val, TTL: 0}
	return nil
}

func (c *TTLCache) Delete(key interface{}) error {
	return nil
}

func (c *TTLCache) clean() {
	tick := time.Tick(time.Second)
	for range tick {
		for k, v := range c.Cache {
			v.TTL += 1
			fmt.Println("clean: ", v, v.TTL >= c.TTL)
			if v.TTL >= c.TTL {
				delete(c.Cache, k)
			}
		}
	}
}
