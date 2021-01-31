package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(str2tree("4(2(3)(1))(6(5))"))
}

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	if s[0] == '(' {
		return str2tree(s[1 : len(s)-1])
	}

	i := 0
	for ; i < len(s) && (s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-'); i++ {
	}
	fmt.Println(i)
	num, _ := strconv.Atoi(s[:i])
	node := &TreeNode{
		Val: num,
	}
	if i == len(s) {
		return node
	}

	count := 1
	j := i + 1
	for ; j < len(s) && count > 0; j++ {
		if s[j] == '(' {
			count++
		} else if s[j] == ')' {
			count--
		}
	}

	node.Left = str2tree(s[i:j])
	node.Right = str2tree(s[j:])
	return node
}
