package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
func findAnagrams(s string, p string) []int {
	sArr, pArr := [26]int{}, [26]int{}
	var ans = []int
	for _, v := range p {
		pArr[v-'a']++
	}
	for i, v := range s {
		if i >= len(p) {
			sArr[s[i-len(p)]-'a']--
		}
		sArr[v-'a']++
		if sArr == pArr {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
