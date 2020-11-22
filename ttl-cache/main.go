package main

import (
	"fmt"
	"time"
	ttlcache "ttl-cache/cache"
)

func main() {
	cache := ttlcache.NewTTLCache(3)
	cache.Set("aa", 1)
	v1, ok := cache.Get("aa")
	fmt.Println("cache:aa=", v1, ok)
	time.Sleep(4 * time.Second)
	v2, ok := cache.Get("aa")
	fmt.Println("cache:aa=", v2, ok)
}
