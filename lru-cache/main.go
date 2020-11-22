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
	v, err := cache.Get("aa")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("cache:aa=", v)
	cache.Set("dd", "dd")
	v2, err := cache.Get("bb")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("cache:bb=", v2)
	v3, err := cache.Get("cc")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("cache:cc=", v3)
}
