package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	fmt.Println(dp)
	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] = min(dp[x], dp[x-coin]+1)
			fmt.Println(x)
			fmt.Println(coin)
			fmt.Println(dp)
			fmt.Println("===================")
		}
	}

	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	} else {
		return dp[len(dp)-1]
	}
}
