package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

/*
i: 0
j: 3
left: 1
right: 4
res: [1 1 1 1]
i: 1
j: 2
left: 2
right: 12
res: [1 1 4 1]
i: 2
j: 1
left: 6
right: 24
res: [1 12 8 1]
i: 3
j: 0
left: 24
right: 24
res: [24 12 8 6]
[24 12 8 6]
*/
// Time complexity : O(N)
// Space complexity : O(N)
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}
	left := 1
	right := 1
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1 - i
		res[j] = res[j] * right
		right = right * nums[j]
		res[i] = res[i] * left
		left = left * nums[i]
		fmt.Printf("i: %d \n", i)
		fmt.Printf("j: %d \n", j)
		fmt.Printf("left: %d \n", left)
		fmt.Printf("right: %d \n", right)
		fmt.Printf("res: %d \n", res)
	}

	return res
}
