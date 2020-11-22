package main

import (
	"fmt"
	lrucache "lru-cache/cache"
)

func main() {
	cache := lrucache.NewLRUCache(3)
	cache.Set("aa", "aa")
	cache.Set("bb", "bb")
	cache.Set("cc", "cc")
	v, ok := cache.Get("aa")
	fmt.Println("cache:aa=", v, ok)
	cache.Set("dd", "dd")
	v2, ok := cache.Get("bb")
	fmt.Println("cache:bb=", v2, ok)
	v3, ok := cache.Get("cc")
	fmt.Println("cache:cc=", v3, ok)
}
