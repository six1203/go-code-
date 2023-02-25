package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type ByAge []Person

func (p ByAge) Len() int {
	return len(p)
}

func (p ByAge) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p ByAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	family := []Person{
		{"David", 2},
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(family))
	fmt.Println(family)
}
