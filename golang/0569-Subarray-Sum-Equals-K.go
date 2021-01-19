package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{3, 4, 7, 2, -3, 1, 4, 2}, 7))
}

func subarraySum(nums []int, k int) int {

	count, sum := 0, 0
	sums := make(map[int]int)
	sums[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if val, ok := sums[sum-k]; ok {
			fmt.Println(sums)
			fmt.Println(sum)
			fmt.Println(val)
			count += val
		}
		sums[sum]++
		// map.put(sum, map.getOrDefault(sum, 0) + 1);
	}
	return count
}

// sums := make(map[int]int)
// res := 0
// sum := 0
// sums[0] = 1
// for i := 0; i < len(nums); i++ {
// 	sum += nums[i]
// 	res += sums[sum-k]
// 	fmt.Println(sum - k)
// 	fmt.Println(sums)
// 	sums[sum]++
// 	fmt.Println(sum)
// 	fmt.Println(res)
// 	fmt.Println(sums)
// 	fmt.Println(sums)
// }
// return res
// }
