package main

import (
	"fmt"
	goset "go-set/set"
)

func main() {
	set := goset.NewSet()
	set.Add("a")
	set.Add("b")
	fmt.Println("set:", set.String())
	set.Add("a")
	fmt.Println("set:", set.String())
	set.Remove("a")
	fmt.Println("set:", set.String())
}
