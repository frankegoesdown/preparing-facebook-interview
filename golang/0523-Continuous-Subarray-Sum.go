package main

import "fmt"

func main() {
	fmt.Println(checkSubarraySum([]int{2, 5, 33, 6, 7, 25, 15}, 13))
}

func checkSubarraySum(nums []int, k int) bool {
	tmp := make(map[int]int)
	tmp[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		fmt.Println(sum)
		sum += nums[i]
		if k != 0 {
			sum = sum%k
		}
		fmt.Println(sum)
		if prevIndex, ok := tmp[sum]; ok {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			tmp[sum] = i
		}
		fmt.Println(tmp)
		fmt.Println("===========")
	}
	return false
}
