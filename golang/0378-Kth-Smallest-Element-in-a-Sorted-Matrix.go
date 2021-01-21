package main

import (
	"container/heap"
	"fmt"
)
func main() {
	fmt.Println(kthSmallest2([][]int{[]int{1,5,9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8))
}

func kthSmallest2(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	min := matrix[0][0]
	max := matrix[n - 1][n - 1]

	for min < max {
		mid := (min + max) / 2
		if checkPos(matrix, mid, k) {
			max = mid //if mid position >= k
		} else {
			min = mid + 1 // + 1 is necessary to make sure l == r in the end
		}

	}
	return max
}

func checkPos(matrix [][]int, mid, k int) bool {
	n := len(matrix)
	count := 0
	i := n - 1 //start from matrix's left bottom corner
	for j := 0; j < n; j++ {
		for i >= 0 && matrix[i][j] > mid {
			i--
		}
		count += i + 1
	}
	return count >= k
}



func kthSmallest(matrix [][]int, k int) int {
	mlen, nlen := len(matrix), len(matrix[0])
	var pq Elems
	for i := 0; i < mlen; i++ {
		pq = append(pq, Elem{m: i, n: 0, val: matrix[i][0]})
	}
	fmt.Printf("%+v\n", pq)
	heap.Init(&pq)
	var res int
	for i := 0; i < k; i++ {
		fmt.Println(i)
		e, _ := heap.Pop(&pq).(Elem)
		if e.n < nlen-1 {
			heap.Push(&pq, Elem{m: e.m, n: e.n + 1, val: matrix[e.m][e.n+1]})
		}
		res = e.val
		fmt.Println(res)
		fmt.Printf("%+v\n", pq)
	}
	return res
}

// Below code is Golang boiler plate to suffice priority queue interface.
// Refer https://golang.org/pkg/container/heap/
type Elem struct {
	m   int
	n   int
	val int
}

type Elems []Elem

func (pq Elems) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq Elems) Len() int {
	return len(pq)
}

func (pq Elems) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq *Elems) Pop() interface{} {
	vpq := *pq
	x := vpq[len(vpq)-1]
	*pq = vpq[:len(vpq)-1]
	return x
}

func (pq *Elems) Push(num interface{}) {
	numElm, _ := num.(Elem)
	*pq = append(*pq, numElm)
}
