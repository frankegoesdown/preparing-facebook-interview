package main

import (
	"strconv"
	"strings"
)

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
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := []string{}
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
