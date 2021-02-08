package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	c := Constructor()
	fmt.Println(c.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
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
	}))
	fmt.Println(c.deserialize("1,2,NIL,NIL,3,4,NIL,NIL,5,NIL,NIL"))
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var s []string
	this.serializeHelper(root, &s)
	return strings.Join(s, ",")
}
func (this *Codec) serializeHelper(node *TreeNode, s *[]string) {
	if node == nil {
		*s = append(*s, "NIL")
		return
	}
	*s = append(*s, strconv.Itoa(node.Val))
	this.serializeHelper(node.Left, s)
	this.serializeHelper(node.Right, s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(s string) *TreeNode {
	if s == "NIL" {
		return nil
	}

	strs := strings.Split(s, ",")
	nodes := make([]*TreeNode, len(strs))
	for i := 0; i < len(strs); i++ {
		if strs[i] != "NIL" {
			n, _ := strconv.Atoi(strs[i])
			nodes[i] = &TreeNode{
				Val: n,
			}
		}
	}

	return this.deserializeHelper(&nodes)
}

func (this *Codec) deserializeHelper(nodes *[]*TreeNode) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	node := (*nodes)[0]
	*nodes = (*nodes)[1:len(*nodes)]
	if node == nil {
		return node
	}
	node.Left = this.deserializeHelper(nodes)
	node.Right = this.deserializeHelper(nodes)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
