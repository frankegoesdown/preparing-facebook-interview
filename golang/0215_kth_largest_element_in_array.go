package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
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
