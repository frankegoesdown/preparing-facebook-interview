package main

import "fmt"

func main() {
	fmt.Println(minDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := minDepth(root.Left)
	rightValue := minDepth(root.Right)
	fmt.Println(leftValue)
	fmt.Println(rightValue)
	fmt.Println("==========")

	if leftValue == 0 || rightValue == 0 {
		return leftValue + rightValue + 1
	} else {
		return min(leftValue, rightValue) + 1
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
