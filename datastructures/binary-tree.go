package datastructures

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{
				Key: k,
			}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{
				Key: k,
			}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func MainBinaryTree() {
	tree := &Node{
		Key: 100,
	}
	fmt.Println(tree)
	tree.Insert(200)
	tree.Insert(400)
	fmt.Println(tree.Search(600))
	fmt.Println(tree)

}
