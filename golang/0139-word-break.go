package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if i-len(w) >= 0 && dp[i-len(w)] && s[i-len(w):i] == w {
				fmt.Println(i)
				fmt.Println(i-len(w))
				fmt.Println(s[i-len(w):i])
				fmt.Println(w)
				dp[i] = true
				break
			}
			fmt.Println(dp)
		}
		fmt.Println("==========")
	}
	fmt.Println(dp)
	return dp[len(s)]
}
