package set

import (
	"strings"
)

type Set struct {
	Map map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		Map: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(key interface{}) {
	s.Map[key] = struct{}{}
}

func (s *Set) Remove(key interface{}) {
	delete(s.Map, key)
}

func (s *Set) Contains(key interface{}) bool {
	_, ok := s.Map[key]
	return ok
}

func (s *Set) String() string {
	a := make([]string, 0)
	for k := range s.Map {
		a = append(a, k.(string))
	}
	r := strings.Join(a, ",")
	return r
}
