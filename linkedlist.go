package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	root *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Push(data string) {
	if l.root == nil {
		l.root = &Node{data: data}
		return
	}

	node := l.root
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{data: data}
}


func (l *LinkedList) Remove(data string) {
	if l.root == nil {
		return
	}
	if l.root.data == data {
		l.root = l.root.next
		return
	}

	node := l.root
	for node.next != nil {
		if node.next.data == data {
			node.next = node.next.next;
			return
		}
		node = node.next
	}
}

func (l *LinkedList) Print() {
	node := l.root
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	list := NewLinkedList()
	list.Push("x")
	list.Push("b")
	list.Push("c")
	
	list.Remove("x")
	list.Remove("c")
	list.Print()
}
