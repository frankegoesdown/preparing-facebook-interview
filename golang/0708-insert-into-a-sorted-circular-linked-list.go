package main

import (
	"fmt"
)

func main() {
	n1 := &Node{
		Val:  3,
		Next: nil,
	}
	n2 := &Node{
		Val:  4,
		Next: nil,
	}
	n3 := &Node{
		Val:  1,
		Next: nil,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	n := insert(n1, 2)
	fmt.Printf("%+v\n", n)
	fmt.Printf("%+v\n", n.Next)
	fmt.Printf("%+v\n", n.Next.Next)
	fmt.Printf("%+v\n", n.Next.Next.Next)

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
		fmt.Println(n.Val)
		fmt.Println(x)
		fmt.Println(n.Next.Val)
		fmt.Printf("insertAfter: %+v\n", insertAfter)
		fmt.Println(biggestNode.Val)
		fmt.Println("----------")
	}

	if insertAfter == nil {
		biggestNode.Next = &Node{Val: x, Next: biggestNode.Next}
	} else {
		insertAfter.Next = &Node{Val: x, Next: insertAfter.Next}
	}
	fmt.Printf("================")
	return aNode
}
