package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoubleLinkedList) Insert(val int) {
	/*head -tail*/

	/* node <> head */
	node := &Node{val: val}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
}

func (list *DoubleLinkedList) Print() {
	node := list.head

	for node != nil {
		fmt.Printf("%d\n", node.val)
		node = node.next
	}

}

func main() {
	list := &DoubleLinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Print()
}
