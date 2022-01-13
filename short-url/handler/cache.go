package handler

import (
	"sync"
)

type Cache struct {
	Memo *sync.Map
}

func NewCache() *Cache {
	return &Cache{
		Memo: &sync.Map{},
	}
}

func (c *Cache) Set(k, v string) error {
	c.Memo.Store(k, v)
	return nil
}

func (c *Cache) Get(k string) (string, bool) {
	if v, ok := c.Memo.Load(k); !ok {
		return "", ok
	} else {
		return v.(string), ok
	}
}
