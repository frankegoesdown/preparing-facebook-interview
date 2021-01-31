package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   11,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   13,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}
	c := Constructor(t)
	fmt.Printf("val: %d",c.Next())
	fmt.Printf("val: %d",c.Next())
	fmt.Printf("val: %d",c.Next())
	fmt.Println(c.HasNext())
	fmt.Printf("val: %d",c.Next())
	fmt.Println(c.HasNext())
	fmt.Printf("val: %d",c.Next())
	fmt.Printf("val: %d",c.Next())
	fmt.Printf("val: %d",c.Next())
	fmt.Println(c.HasNext())
	fmt.Printf("val: %d",c.Next())
	fmt.Println(c.HasNext())
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{
		stack: stack,
	}
	res.push(root)
	return res
}

// Next returns the next smallest number
func (it *BSTIterator) Next() int {
	size := len(it.stack)
	var top *TreeNode
	it.stack, top = it.stack[:size-1], it.stack[size-1]
	it.push(top.Right)
	return top.Val
}

// HasNext returns whether we have a next smallest number
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func (it *BSTIterator) push(root *TreeNode) {
	for root != nil {
		//fmt.Println(root.Val)
		it.stack = append(it.stack, root)
		root = root.Left
	}
	fmt.Printf("%+v \n", it.stack)
	for _, v := range it.stack{
		fmt.Println(v.Val)
	}
	fmt.Println("===========")
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
