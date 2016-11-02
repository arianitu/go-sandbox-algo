package main

import (
	"fmt"
	"math"
)

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{nodes: make([]*Node, 0)}
}

func (q *Queue) Push(node *Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Pop() *Node {
	if q.IsEmpty() {
		return nil
	}

	r := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	return r
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

func findShortestPathToNode(root *Node, data int) []int {
	path := make([]int, 0)

	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		node := q.Pop()

		if node.data == data {
			for node.parent != nil {
				path = append(path, node.data)
				node = node.parent
			}

			// root
			path = append(path, node.data)
			break
		}
		if node.left != nil {
			node.left.parent = node
			q.Push(node.left)
		}
		if node.right != nil {
			node.right.parent = node
			q.Push(node.right)
		}
	}

	return reverse(path)
}

func reverse(nums []int) []int {
	reversed := make([]int, len(nums))
	i := len(nums) - 1
	n := 0
	for i >= 0 {
		reversed[n] = nums[i]
		i -= 1
		n += 1
	}

	return reversed
}

func height(n *Node) float64 {
	if n == nil {
		return -1
	}
	return 1 + math.Max(height(n.left), height(n.right))
}

func main() {
	tree := &Node{
		data: 1,
		left: &Node{
			data: 2,
			left: &Node{
				data: 3,
				left: &Node{
					data: 4,
					left: &Node{
						data: 5,
					},
				},
			},
		},
		right: &Node{
			data: 6,
			right: &Node{
				data: 7,
			},
		},
	}

	fmt.Println("height of tree is", height(tree))
	fmt.Println("path to 5 is", findShortestPathToNode(tree, 5))

}
