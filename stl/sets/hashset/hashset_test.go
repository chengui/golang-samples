package hashset

import (
	"fmt"
)

func Example() {
	s := New()
	s.Add("a")
	s.Add("b")
	fmt.Println("s:", s.String())
	s.Add("a")
	fmt.Println("s:", s.String())
	s.Remove("a")
	fmt.Println("s:", s.String())
	s.Add("c")
	s.Range(func(k interface{}) bool {
		fmt.Println(k)
		return true
	})
	// Output:
	// s: a,b
	// s: a,b
	// s: b
	// c
	// b
}
