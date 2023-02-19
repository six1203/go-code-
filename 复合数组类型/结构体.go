package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

type Tree struct {
	Val         int
	left, right *Tree
}

func add(tree *Tree, val int) *Tree {
	if val < tree.Val {
		tree.left = add(tree.left, val)
	} else {
		tree.right = add(tree.right, val)
	}
	return tree
}

func Sort(values []int) {
	var root *Tree

	for _, v := range values {
		root = add(root, v)
	}
	AppendValues(values[:0], root)

}

func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.Val)
		values = AppendValues(values, t.right)
	}
	return values
}
func main() {

	//fmt.Println(dilbert.Salary)
	//dilbert.Salary -= 50000
	//fmt.Println(dilbert.Salary)
	//
	//position := &dilbert.Position
	//
	//*position = "Senior " + *position
	//
	//fmt.Println(dilbert.Position, dilbert.DoB)

	a := &dilbert
	fmt.Println(*a)
}
