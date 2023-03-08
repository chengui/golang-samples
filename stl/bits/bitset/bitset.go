package bitset

import (
	"bytes"
	"fmt"
)

const (
	Base = 8
)

type BitSet struct {
	bits []uint8
}

func NewBitSet() *BitSet {
	return &BitSet{}
}

func (v *BitSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, w := range v.bits {
		if w == 0 {
			continue
		}
		for j := 0; j < Base; j++ {
			if w&(1<<j) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", Base*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (v *BitSet) Has(x int) bool {
	w, b := x/Base, x%Base
	return w < len(v.bits) && v.bits[w]&(1<<b) != 0
}

func (v *BitSet) Add(x int) {
	w, b := x/Base, x%Base
	for len(v.bits) <= w {
		v.bits = append(v.bits, 0)
	}
	v.bits[w] |= (1 << b)
}

func (v *BitSet) Remove(x int) {
	w, b := x/Base, x%Base
	if w <= len(v.bits) {
		v.bits[w] &= (0xFE << b)
	}
}

func (v *BitSet) Clear() {
	v.bits = make([]uint8, len(v.bits))
}

func (v *BitSet) Copy() *BitSet {
	u := new(BitSet)
	u.bits = make([]uint8, len(v.bits))
	copy(u.bits, v.bits)
	return u
}

func (v *BitSet) Len() int {
	cnt := 0
	for _, w := range v.bits {
		u := w
		for u != 0 {
			u &= (u - 1)
			cnt++
		}
	}
	return cnt
}

func (v *BitSet) Elems() []int {
	o := make([]int, 0)
	for i, w := range v.bits {
		if w == 0 {
			continue
		}
		for j := 0; j < Base; j++ {
			if w&(1<<j) != 0 {
				o = append(o, Base*i+j)
			}
		}
	}
	return o
}

func (v *BitSet) AddAll(xs ...int) {
	for _, x := range xs {
		v.Add(x)
	}
}

func (v *BitSet) UnionWith(u *BitSet) {
	for i := range u.bits {
		if i < len(v.bits) {
			v.bits[i] |= u.bits[i]
		} else {
			v.bits = append(v.bits, u.bits[i])
		}
	}
}

func (v *BitSet) IntersectWith(u *BitSet) {
	for i := range v.bits {
		if i < len(u.bits) {
			v.bits[i] &= u.bits[i]
		} else {
			v.bits[i] = 0
		}
	}
}

func (v *BitSet) DifferenceWith(u *BitSet) {
	for i := range u.bits {
		if i < len(v.bits) {
			v.bits[i] &= ^u.bits[i]
		} else {
			v.bits = append(v.bits, 0)
		}
	}
}

func (v *BitSet) SymmetricDifference(u *BitSet) {
	for i := range u.bits {
		if i < len(v.bits) {
			v.bits[i] = v.bits[i] ^ u.bits[i]
		} else {
			v.bits = append(v.bits, u.bits[i])
		}
	}
}
