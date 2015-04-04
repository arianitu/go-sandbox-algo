package bfs

import (
	"container/list"
)

type Node struct {
	data string
	left *Node
	right *Node
}

func Bfs(n *Node, visit func (n *Node)) {
	if n == nil {
		panic("Root node is null")
	}
	
	queue := list.New()
	queue.PushBack(n)

	elem := queue.Front()
	for elem != nil {
		node, ok := queue.Remove(elem).(*Node)
		if ! ok {
			panic("Node is not of type *Node")
		}

		visit(node)
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
		elem = queue.Front()
	}
}

