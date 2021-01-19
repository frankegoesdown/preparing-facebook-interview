package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kClosest2([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		return points[a][0]*points[a][0]+points[a][1]*points[a][1] <
			points[b][0]*points[b][0]+points[b][1]*points[b][1]
	})
	return points[:K]
}

// ========================================

type Point struct {
	distance int
	value    []int
}

type PQueue []Point

func (p PQueue) Len() int {
	return len(p)
}

func (p PQueue) Less(a, b int) bool {
	return p[a].distance > p[b].distance
}

func (p PQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQueue) Push(x interface{}) {
	point := x.(Point)
	*p = append(*p, point)
}

func (p *PQueue) Pop() interface{} {
	old := *p
	n := len(old)
	point := old[n-1]
	//old[n-1] = nil
	*p = old[0 : n-1]
	return point
}

func kClosest2(points [][]int, k int) [][]int {
	pq := PQueue{}
	for _, point := range points {
		p := Point{
			distance: (point[0] * point[0]) + (point[1] * point[1]),
			value:    point,
		}
		heap.Push(&pq, p)
		fmt.Println(&pq)
		if pq.Len() > k {
			heap.Pop(&pq)
		}
		fmt.Println(&pq)
		fmt.Println("==========")
	}

	res := make([][]int, k)

	for i := 0; i < len(pq); i++ {
		res[i] = pq[i].value
	}
	return res
}
