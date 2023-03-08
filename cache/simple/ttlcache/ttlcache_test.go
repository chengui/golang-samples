package ttlcache

import (
	"fmt"
	"time"
)

func Example() {
	cache := New(3)
	cache.Set("aa", 1)
	v1, ok := cache.Get("aa")
	fmt.Println("cache:aa=", v1, ok)
	time.Sleep(4 * time.Second)
	v2, ok := cache.Get("aa")
	fmt.Println("cache:aa=", v2, ok)
	// Output:
}
