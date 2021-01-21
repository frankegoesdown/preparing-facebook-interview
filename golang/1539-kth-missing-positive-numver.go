package main

import "fmt"
func main() {
	fmt.Println(findKthPositive([]int{2,3,4,7,11}, 5))
}

func findKthPositive(arr []int, k int) int {
	l := 0
	r := len(arr) - 1

	//binary search
	//arr[idx] - idx - 1 is the number of missing integers of position idx
	for l <= r {
		mid := r - (r - l) / 2
		fmt.Println(arr[mid])
		fmt.Println(mid)
		if arr[mid] - mid - 1 < k {
			l = mid + 1
		} else {
			r = mid -1
		}
		fmt.Println(l)
		fmt.Println("==============")
	}

	return k + l
}
