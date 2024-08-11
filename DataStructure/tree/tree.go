package tree

import "fmt"

type Node struct {
	Data  int
	left  *Node
	right *Node
}

func (t *Node) Insert(data int) {

	if t.Data < data { // Insert Node at left

		if t.left == nil {
			t.left = &Node{Data: data}
		} else {
			t.left.Insert(data)
		}

	} else {
		if t.right == nil {
			t.right = &Node{Data: data}
		} else {
			t.right.Insert(data)
		}
	}

}

func (n *Node) TraverseNode() {
	if n == nil {
		return
	}
	n.left.TraverseNode()
	fmt.Println("Print ", n.Data)
	n.right.TraverseNode()
}
