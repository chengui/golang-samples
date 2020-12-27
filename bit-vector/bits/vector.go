package bits

import (
	"bytes"
	"fmt"
)

const (
	Base = 8
)

type BitVector struct {
	Bits []uint8
}

func NewBitVector() *BitVector{
	return &BitVector{}
}

func (v *BitVector) Has(x int) bool {
	w, b := x / Base, x % Base
	return w < len(v.Bits) && v.Bits[w] & (1<<b) != 0
}

func (v *BitVector) Add(x int) {
	w, b := x / Base, x % Base
	for len(v.Bits) <= w {
		v.Bits = append(v.Bits, 0)
	}
	v.Bits[w] |= (1<<b)
}

func (v *BitVector) Remove(x int) {
	w, b := x / Base, x % Base
	if w <= len(v.Bits) {
		v.Bits[w] &= (0xFE<<b)
	}
}

func (v *BitVector) Clear() {
	v.Bits = make([]uint8, len(v.Bits))
}

func (v *BitVector) Copy() *BitVector {
	u := new(BitVector)
	u.Bits = make([]uint8, len(v.Bits))
	copy(u.Bits, v.Bits)
	return u
}

func (v *BitVector) Len() int {
	cnt := 0
	for _, w := range v.Bits {
		u := w
		for u != 0 {
			u &= (u - 1)
			cnt++
		}
	}
	return cnt
}

func (v *BitVector) Elems() []int {
	o := make([]int, 0)
	for i, w := range v.Bits {
		if w == 0 {
			continue
		}
		for j := 0; j < Base; j++ {
			if w & (1<<j) != 0 {
				o = append(o, Base * i + j)
			}
		}
	}
	return o
}

func (v *BitVector) AddAll(xs ...int) {
	for _, x := range xs {
		v.Add(x)
	}
}

func (v *BitVector) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, w := range v.Bits {
		if w == 0 {
			continue
		}
		for j := 0; j < Base; j++ {
			if w & (1<<j) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", Base * i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (v *BitVector) UnionWith(u *BitVector) {
	for i, _ := range u.Bits {
		if i < len(v.Bits) {
			v.Bits[i] |= u.Bits[i]
		} else {
			v.Bits = append(v.Bits, u.Bits[i])
		}
	}
}

func (v *BitVector) IntersectWith(u *BitVector) {
	for i, _ := range v.Bits {
		if i < len(u.Bits) {
			v.Bits[i] &= u.Bits[i]
		} else {
			v.Bits[i] = 0
		}
	}
}

func (v *BitVector) DifferenceWith(u *BitVector) {
	for i, _ := range u.Bits {
		if i < len(v.Bits) {
			v.Bits[i] &= ^u.Bits[i]
		} else {
			v.Bits = append(v.Bits, 0)
		}
	}
}

func (v *BitVector) SymmetricDifference(u *BitVector) {
	for i, _ := range u.Bits {
		if i < len(v.Bits) {
			v.Bits[i] = v.Bits[i]^u.Bits[i]
		} else {
			v.Bits = append(v.Bits, u.Bits[i])
		}
	}
}
