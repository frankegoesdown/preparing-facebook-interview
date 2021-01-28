package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// [4,2,5,1,3]
func main() {
	fmt.Printf("%+v\n", treeToDoublyList(&Node{
		Val: 4,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Node{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}))
}

func buildDll(node *Node, head **Node, tail **Node) {
	if node != nil {
		buildDll(node.Left, head, tail)
		if *head == nil {
			*head = node
			*tail = node
		} else {
			(*tail).Right = node
			node.Left = *tail
			*tail = node
		}
		fmt.Println(node.Val)
		fmt.Println((**head).Val)
		fmt.Println((**(tail)).Val)
		fmt.Println("========")
		buildDll(node.Right, head, tail)
	}
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	var head *Node
	var tail *Node
	buildDll(root, &head, &tail)
	head.Left = tail
	tail.Right = head

	return head
}
