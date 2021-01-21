package main

import "github.com/golang/go/src/pkg/fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(averageOfLevels(&TreeNode{
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

func averageOfLevels(root *TreeNode) (result []float64) {
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	fmt.Printf("%+v\n", queue)
	for len(queue) > 0 {
		var newQueue []*TreeNode
		level := len(result)
		result = append(result, 0)
		fmt.Println(level)
		fmt.Println(result)
		for _, node := range queue {
			result[level] += float64(node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		fmt.Println(newQueue)
		fmt.Println(result)
		result[level] /= float64(len(queue))
		queue = newQueue
		fmt.Println(result)
		fmt.Println("==============")
	}
	return result
}
