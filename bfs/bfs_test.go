package bfs

import (
	"testing"
)

var tree = &Node{
	data: "F",
	left: &Node {
		data: "B",
		left: &Node {
			data: "A",
		},
		right: &Node {
			data: "D",
			left: &Node {
				data: "C",
			},
			right: &Node {
				data: "E",
			},
		},
	},
	right: &Node {
		data: "G",
		right: &Node {
			data: "I",
			left: &Node {
				data: "H",
			},
		},
	},
}

func TestBfs(t *testing.T) {
	p := []string{"F", "B", "G", "A", "D", "I", "C", "E", "H"}
	i := 0
	Bfs(tree, func (n *Node) {
		if n.data != p[i] {
			t.Fatalf("Expected %s, got %s, at index %d", p[i], n.data, i)
		}
		i++
	});
}
