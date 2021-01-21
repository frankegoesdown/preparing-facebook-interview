package main

func main() {
	
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L , R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right,L,R)
}

func rangeSumBST2(root *TreeNode, L int, R int) int {

	if root == nil {
		return 0
	}

	var ans int

	dfs(root, L, R, &ans)
	return ans
}


func dfs(root *TreeNode, L int, R int, ans *int) {
	if root == nil {
		return
	}
	if L <= root.Val && root.Val <= R {
		*ans += root.Val
	}
	if L < root.Val {
		dfs(root.Left, L, R, ans)
	}
	if R > root.Val {
		dfs(root.Right, L, R, ans)
	}
}
