package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(closestValue(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}, float64(3.714286)))
}

func closestValue(root *TreeNode, target float64) int {
	ans := root.Val
	for root != nil {
		fmt.Println(root.Val)
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(ans)-target) {
			ans = root.Val
		}
		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
		//fmt.Println(root.Val)
	}
	return ans
}
