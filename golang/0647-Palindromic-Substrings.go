package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaabc"))
}

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		helper(s, i, i, &count)   // odd
		helper(s, i, i+1, &count) // even
	}
	return count
}

func helper(s string, left, right int, c *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		*c++
		left--
		right++
		fmt.Println(s[left])
		fmt.Println(s[right])
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(*c)
		fmt.Println("==========")
	}
}
