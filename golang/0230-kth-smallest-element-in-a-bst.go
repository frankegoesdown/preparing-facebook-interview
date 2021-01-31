package main

import "fmt"

func main() {
	fmt.Println(kthSmallest(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var vals []int

	dfs(root.Left, &vals)
	vals = append(vals, root.Val)
	if len(vals) >= k {
		return vals[k-1]
	}
	dfs(root.Right, &vals)

	return vals[k-1]
}

func dfs(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, vals)

	*vals = append(*vals, node.Val)
	dfs(node.Right, vals)
}
