package main

import (
	"fmt"
	s "strset"
	i "int8set"
	f "float32set"
	b "byteset"
)

func main() {
	s1 := s.Make("red", "green", "blue", "yellow")
	s2 := s.Make("yellow", "green", "white")
	fmt.Println(s1.Intersection(s2))

	i1 := i.Make(1, 2, 3, 4)
	i2 := i.Make(4, 2, 0)
	fmt.Println(i1.Intersection(i2))

	f1 := f.Make(.1, .2, .3, .4)
	f2 := f.Make(.4, .2, .0)
	fmt.Println(f1.Intersection(f2))

	b1 := b.Make('1', '2', '3', '4')
	b2 := b.Make('4', '2', '0')
	fmt.Println(b1.Intersection(b2))
}
