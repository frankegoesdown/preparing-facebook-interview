package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(verticalTraversal(&TreeNode{
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

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	cache := make(map[int][][2]int)
	minCol, maxCol := 0, 0
	dfs(root, cache, 0, 0, &minCol, &maxCol)
	fmt.Println(cache)
	fmt.Println(minCol)
	fmt.Println(maxCol)
	for i := minCol; i <= maxCol; i++ {
		sort.Slice(cache[i], func(a, b int) bool {
			return cache[i][a][0] < cache[i][b][0] || (cache[i][a][0] == cache[i][b][0] && cache[i][a][1] < cache[i][b][1])
		})
		var tmp []int
		for _, v := range cache[i] {
			tmp = append(tmp, v[1])
		}
		res = append(res, tmp)
	}
	return res
}

func dfs(root *TreeNode, cache map[int][][2]int, r, c int, minCol, maxCol *int) {
	if root == nil {
		return
	}
	if _, ok := cache[c]; !ok {
		cache[c] = [][2]int{[2]int{r, root.Val}}
	} else {
		cache[c] = append(cache[c], [2]int{r, root.Val})
	}
	*minCol = min(*minCol, c)
	*maxCol = max(*maxCol, c)
	dfs(root.Left, cache, r+1, c-1, minCol, maxCol)
	dfs(root.Right, cache, r+1, c+1, minCol, maxCol)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
