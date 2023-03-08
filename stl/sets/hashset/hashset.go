package hashset

import (
	"strings"
)

type HashSet struct {
	m map[interface{}]struct{}
}

func New() *HashSet {
	return &HashSet{
		m: make(map[interface{}]struct{}),
	}
}

func (s *HashSet) Add(key interface{}) {
	s.m[key] = struct{}{}
}

func (s *HashSet) Remove(key interface{}) {
	delete(s.m, key)
}

func (s *HashSet) Contains(key interface{}) bool {
	_, ok := s.m[key]
	return ok
}

func (s *HashSet) Empty() bool {
	return s.Size() == 0
}

func (s *HashSet) Size() int {
	return len(s.m)
}

func (s *HashSet) Clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *HashSet) Values() []interface{} {
	v := make([]interface{}, 0)
	for k := range s.m {
		v = append(v, k)
	}
	return v
}

func (s *HashSet) Range(f func(k interface{}) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}

func (s *HashSet) String() string {
	a := make([]string, 0)
	for k := range s.m {
		a = append(a, k.(string))
	}
	r := strings.Join(a, ",")
	return r
}
