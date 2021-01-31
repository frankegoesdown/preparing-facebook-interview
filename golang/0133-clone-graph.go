package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	first := &Node{
		Val:       1,
		Neighbors: nil,
	}
	second := &Node{
		Val:       2,
		Neighbors: nil,
	}
	third := &Node{
		Val:       3,
		Neighbors: nil,
	}
	fourth := &Node{
		Val:       4,
		Neighbors: nil,
	}
	first.Neighbors = append(first.Neighbors, second)
	first.Neighbors = append(first.Neighbors, fourth)

	second.Neighbors = append(second.Neighbors, first)
	second.Neighbors = append(second.Neighbors, third)

	third.Neighbors = append(third.Neighbors, first)
	third.Neighbors = append(third.Neighbors, fourth)

	fourth.Neighbors = append(fourth.Neighbors, first)
	fourth.Neighbors = append(fourth.Neighbors, second)
	fmt.Println(cloneGraph(first))
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return walk(node, visited)
}

func walk(node *Node, visited map[*Node]*Node) *Node {
	fmt.Println(node.Val)
	fmt.Println(visited)
	if n, ok := visited[node]; ok {
		fmt.Println("visited:", n.Val)
		return n
	}
	n := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	visited[node] = n
	for i, _ := range node.Neighbors {
		n.Neighbors[i] = walk(node.Neighbors[i], visited)
	}
	return n
}
