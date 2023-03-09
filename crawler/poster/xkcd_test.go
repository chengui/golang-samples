package xkcd

import (
	"fmt"
)

func Example() {
	x := NewXkcd()
	x.Load(500, 600)
	fmt.Println(x.Find("car"))
}
