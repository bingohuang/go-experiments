package main

import (
	"fmt"
)

type set struct {
	m map[string]struct{}
}

func NewSet()  *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(v string) {
	s.m[v] = struct{}{}
}

func (s *set) Contains(v string) bool {
	_, c := s.m[v]
	return c
}

func (s *set) Remove(v string)  {
	delete(s.m, v)
}

func main() {
	s := NewSet()

	s.Add("Bingo")
	s.Add("Huang")

	fmt.Println(s.Contains("Bingo"))
	fmt.Println(s.Contains("Huang"))

	s.Remove("Bingo")
	fmt.Println(s.Contains("Bingo"))
}
