package dfs

import "testing"

var tree = &Node{
	data: "LOL",
	left: &Node {
		data: "DOG",
		left: &Node {
			data: "MEEP",
		},
		right: &Node {
			data: "KEEP",
		},
	},
	right: &Node {
		data: "YO",
	},
}

func TestPreorder(t *testing.T) {
	p := []string{"LOL", "DOG", "MEEP", "KEEP", "YO"}
	i := 0
	visit := func(n *Node) {
		if (n.data != p[i]) {
			t.Fatalf("At index %d, expected %s, got %s", i, p[i], n.data);
		}
		i++
	}
	Preorder(tree, visit)
}

func TestInorder(t *testing.T) {
	p := []string{"MEEP", "DOG", "KEEP", "LOL", "YO"}
	i := 0
	visit := func(n *Node) {
		if (n.data != p[i]) {
			t.Fatalf("At index %d, expected %s, got %s", i, p[i], n.data);
		}
		i++
	}
	Inorder(tree, visit)
}

func TestPostorder(t *testing.T) {
	p := []string{"MEEP", "KEEP", "DOG", "YO", "LOL"}
	i := 0
	visit := func(n *Node) {
		if (n.data != p[i]) {
			t.Fatalf("At index %d, expected %s, got %s", i, p[i], n.data);
		}
		i++
	}
	Postorder(tree, visit)
}

