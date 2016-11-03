package main

import (
	"fmt"
)

type Node struct {
	data int
	left *Node
	right *Node
}

func insert(node *Node, data int) *Node {
	if node == nil {
		return &Node {data: data}
	}
	if data < node.data  {
		node.left = insert(node.left, data)
		return node
	}
	
	node.right = insert(node.right, data)
	return node
}

func find(node *Node, data int) int {
	if node == nil {
		return 0
	}
	
	if node.data == data {
		return node.data
	}
	if data < node.data {
		fmt.Println("go left")
		return find(node.left, data)
	}

	fmt.Println("go right")
	return find(node.right, data)
}


func inorder(node *Node) {
	if node.left != nil {
		inorder(node.left)
	}
	fmt.Println(node.data)
	if node.right != nil {
		inorder(node.right)
	}
}

func main() {
	root := insert(nil, 1)
	
	insert(root, 10)
	insert(root, 15)
	insert(root, 20)
	insert(root, 2)

	inorder(root)
}
