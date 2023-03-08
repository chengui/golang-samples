package bitset

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
}

func TestRemove(t *testing.T) {

}

func TestCopy(t *testing.T) {

}

func TestAddAll(t *testing.T) {
}

func TestUnionWith(t *testing.T) {

}

func TestIntersectWith(t *testing.T) {

}

func TestDifferenceWith(t *testing.T) {

}

func TestSymmetricDifference(t *testing.T) {

}

// -------- Examples ----------
func ExampleBitSet_Add() {
	x := NewBitSet()
	x.Add(1)
	x.Add(9)
	x.Add(143)
	x.Add(144)
	fmt.Println("x:", x.String())
	x.Remove(144)
	x.Remove(200)
	fmt.Println("x:", x.String())
	fmt.Println("x:", x.Has(9), x.Has(123))
	// Output:
	// x: {1 9 143 144}
	// x: {1 9 143}
	// x: true false
}

func ExampleBitSet_Copy() {
	x := NewBitSet()
	x.Add(1)
	x.Add(9)
	y := x.Copy()
	fmt.Println("y:", y.String())
	// Output:
	// y: {1 9}
}

func ExampleBitSet_AddAll() {
	x := NewBitSet()
	x.AddAll(9, 42)
	fmt.Printf("len(x)=%d\n", x.Len())
	fmt.Println("x:", x.String())
	for _, v := range x.Elems() {
		fmt.Printf("%d ", v)
	}
	// Output:
	// len(x)=2
	// x: {9 42}
	// 9 42
}

func ExampleBitSet() {
	x, y := NewBitSet(), NewBitSet()
	x.AddAll(1, 2)
	y.AddAll(1, 3)

	x1 := x.Copy()
	x1.UnionWith(y)
	fmt.Println("x|y:", x1)

	x2 := x.Copy()
	x2.IntersectWith(y)
	fmt.Println("x&y:", x2)

	x3 := x.Copy()
	x3.DifferenceWith(y)
	fmt.Println("x-y:", x3)

	x4 := x.Copy()
	x4.SymmetricDifference(y)
	fmt.Println("x^y:", x3)
	// Output:
	// x|y: {1 2 3}
	// x&y: {1}
	// x-y: {2}
	// x^y: {2}
}
