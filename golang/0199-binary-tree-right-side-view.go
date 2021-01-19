package main

import (
	"fmt"
)

func main() {
	n := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 5, Left: nil, Right: nil}}, Right: &TreeNode{Val: 3, Left: nil, Right: &TreeNode{Val: 4, Left: nil, Right: nil}}}
	fmt.Println(rightSideView(&n))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	rightView(root, &result, 0)
	return result
}

func rightView(curr *TreeNode, result *[]int, currDepth int) {
	if curr == nil {
		return
	}
	if currDepth == len(*result) {
		*result = append(*result, curr.Val)
	}

	rightView(curr.Right, result, currDepth+1)
	rightView(curr.Left, result, currDepth+1)
}
