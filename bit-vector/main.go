package main

import (
	"fmt"

	"bit-vector/bits"
)

func main() {
	x := bits.NewBitVector()
	x.Add(1)
	x.Add(9)
	x.Add(143)
	x.Add(144)
	fmt.Println("x: ", x.String())
	x.Remove(144)
	x.Remove(200)
	fmt.Println("x: ", x.String())
	fmt.Println("x: ", x.Has(9), x.Has(123))

	x1 := x.Copy()
	fmt.Println("x1: ", x1.String())
	x2 := x.Copy()
	fmt.Println("x2: ", x2.String())
	x3 := x.Copy()
	fmt.Println("x3: ", x3.String())
	x4 := x.Copy()
	fmt.Println("x4: ", x4.String())

	y := bits.NewBitVector()
	y.AddAll(9, 42)
	fmt.Println("y: ", y.String())
	fmt.Printf("y=BitVector(%d): ", y.Len())
	for _, v := range(y.Elems()) {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")

	x1.UnionWith(y)
	fmt.Println("x|y: ", x1)

	x2.IntersectWith(y)
	fmt.Println("x&y: ", x2)

	x3.DifferenceWith(y)
	fmt.Println("x-y: ", x3)

	x4.SymmetricDifference(y)
	fmt.Println("x^y: ", x3)
}
