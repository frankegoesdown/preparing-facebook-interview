package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	mismatchs := make([]int, 26)
	for i := 0; i < len(p); i++ {
		mismatchs[int(p[i]-'a')]++
	}
	fmt.Println(mismatchs)

	for i := 0; i < len(p); i++ {
		mismatchs[int(s[i]-'a')]--
	}
	fmt.Println(mismatchs)
	var ans []int
	if check(mismatchs) {
		ans = append(ans, 0)
	}

	for i := 1; i+len(p) <= len(s); i++ {
		end := i + len(p) - 1
		mismatchs[s[end]-'a']--
		mismatchs[s[i-1]-'a']++
		fmt.Println(string(s[end]))
		fmt.Println(string(s[i-1]))
		if check(mismatchs) {
			ans = append(ans, i)
		}
		fmt.Println(ans)
		fmt.Println(mismatchs)
	}

	return ans
}

func check(mismatchs []int) bool {
	for _, mismatch := range mismatchs {
		if mismatch != 0 {
			return false
		}
	}

	return true
}
