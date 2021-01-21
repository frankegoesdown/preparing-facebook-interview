package main

import (
	"fmt"
)

func main() {
	n := &Node{
		Val: 3,
		Next: &Node{
			Val: 4,
			Next: &Node{
				Val: 1,
				Next: nil,
			},
		},
	}
	fmt.Println(insert(n, 2))
}

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		n := &Node{Val: x}
		n.Next = n
		return n
	}

	var insertAfter *Node
	biggestNode := aNode
	if x >= aNode.Val {
		insertAfter = aNode
	}

	for n := aNode.Next; n != aNode; n = n.Next {
		if x >= n.Val && (insertAfter == nil || insertAfter.Val <= n.Val) {
			insertAfter = n
		}

		if biggestNode.Val < n.Val {
			biggestNode = n
		}
	}

	if insertAfter == nil {
		biggestNode.Next = &Node{Val: x, Next: biggestNode.Next}
	} else {
		insertAfter.Next = &Node{Val: x, Next: insertAfter.Next}
	}

	return aNode
}
