package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// we need to mark every tree node's offset from root node
	type node struct {
		*TreeNode     // we extend all members from TreeNode
		offset    int // zero for root, left is negative, right is positive
	}

	queue := []node{{root, 0}}
	var visited []node
	var leftMost, rightMost int
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		visited = append(visited, n)
		// next step, if we go left, offset -1, go right, offset +1
		if n.Left != nil {
			queue = append(queue, node{n.Left, n.offset - 1})
		}
		if n.Right != nil {
			queue = append(queue, node{n.Right, n.offset + 1})
		}
		// record left most or right most
		if n.offset > rightMost {
			rightMost = n.offset
		}
		if n.offset < leftMost {
			leftMost = n.offset
		}
	}
	// result slice length
	l := rightMost - leftMost + 1
	result := make([][]int, l)
	for _, e := range visited {
		// turn offset(ie: -1, 0, 1) to idx(ie: 0, 1, 2) by minus smallest offset
		result[e.offset-leftMost] = append(result[e.offset-leftMost], e.Val)
	}
	return result
}
