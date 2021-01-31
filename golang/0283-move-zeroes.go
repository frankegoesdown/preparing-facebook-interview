package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
		fmt.Println(i)
		fmt.Println(p)
		fmt.Println(nums)
		fmt.Println("========")
	}
}
