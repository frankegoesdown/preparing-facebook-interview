package main

func main() {

}

func largestBSTSubtree(root *TreeNode) (ans int) {
	if root==nil{
		return
	}
	dfs(root, &ans)
	return
}

func dfs(node *TreeNode, ans *int) (int, int, int) {
	if node==nil{
		return 0, math.MaxInt32, math.MinInt32
	}

	leftSize, leftLower, leftUpper:= dfs(node.Left, ans)
	rightSize, rightLower, rightUpper:= dfs(node.Right, ans)

	if leftSize==0 || rightSize==0 || node.Val <= leftUpper || node.Val >= rightLower{
		return 0, 0, 0
	}

	size:= leftSize+rightSize+1

	*ans = max(*ans, size)

	return size, min(leftLower, node.Val), max(rightUpper, node.Val)

}


func min(i, j int) int{
	if i< j{
		return i
	}

	return j
}

func max(i, j int) int{
	if i> j{
		return i
	}

	return j
}
