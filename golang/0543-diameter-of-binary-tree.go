package main

import "fmt"

func main() {
	fmt.Println(diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	L := dfs(root.Left, ans)
	R := dfs(root.Right, ans)
	*ans = max(*ans, L+R)
	fmt.Println(root.Val)
	fmt.Println(L)
	fmt.Println(R)
	fmt.Println(*ans)
	fmt.Println("===============")
	return max(L, R) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
