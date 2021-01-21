package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest2(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

func findKthLargest3(nums []int, k int) int {
	fmt.Println(nums)
	fmt.Println(k)
	k = k - 1
	pivotValue := nums[k]
	fmt.Println(pivotValue)
	fmt.Println("===========")
	nums[len(nums)-1], nums[k] = nums[k], nums[len(nums)-1]
	idx := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= pivotValue {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}

	nums[idx], nums[len(nums)-1] = nums[len(nums)-1], nums[idx]
	if idx == k {
		return nums[idx]
	} else if idx > k {
		return findKthLargest(nums[:idx], k+1)
	} else {
		return findKthLargest(nums[idx+1:], k-idx)
	}
}
