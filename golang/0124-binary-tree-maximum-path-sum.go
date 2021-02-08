package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	//var dfs func(*TreeNode) int

	dfs(root, &maxSum)

	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := max(0, dfs(root.Left, maxSum))
	right := max(0, dfs(root.Right, maxSum))
	sum := left + root.Val + right
	if *maxSum < sum {
		*maxSum = sum
	}

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
