package cache

import (
	"time"
)

type Entry struct {
	Val interface{}
	Exp int64
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
	c.Cache[key] = &Entry{
		Val: val,
		Exp: time.Now().Unix() + int64(c.TTL),
	}
	return nil
}

func (c *TTLCache) Delete(key interface{}) error {
	delete(c.Cache, key)
	return nil
}

func (c *TTLCache) clean() {
	tick := time.Tick(time.Second)
	for range tick {
		now := time.Now().Unix()
		for k, v := range c.Cache {
			if v.Exp <= now {
				delete(c.Cache, k)
			}
		}
	}
}
