package dfs

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func visit(n *Node) {
	fmt.Println(n.data)
}

func Preorder(n *Node, visit func(n *Node)) {
	if n == nil {
		return
	}
	visit(n)
	Preorder(n.left, visit)
	Preorder(n.right, visit)
}

func Inorder(n *Node, visit func(n *Node)) {
	if n == nil {
		return
	}
	Inorder(n.left, visit)
	visit(n)
	Inorder(n.right, visit)
}

func Postorder(n *Node, visit func(n *Node)) {
	if n == nil {
		return
	}
	Postorder(n.left, visit)
	Postorder(n.right, visit)
	visit(n)
}

func main() {
	n := &Node{
		data: "LOL",
		left: &Node{
			data: "DOG",
			left: &Node{
				data: "MEEP",
			},
			right: &Node{
				data: "KEEP",
			},
		},
		right: &Node{
			data: "YO",
		},
	}
	Preorder(n, func(n *Node) {
		fmt.Println(n.data)
	})
}
