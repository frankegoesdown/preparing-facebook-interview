package main

import (
	"fmt"
	"sort"
)

/*
m = map[97:99 98:100 99:101 100:102 101:103 102:104 103:105 104:97 105:106 106:107 107:108 108:98 109:109 110:110 111:111 112:112 113:113 114:114 115:115 116:116 117:117 118:118 119:119 120:120 121:121 122:122]
i =  0
s =  hello
words = [agbbo leetcode]
i =  1
s =  leetcode
words = [agbbo bggteofg]
true
*/

// Time complexity : O(N)
// Space complexity : O(1)
func main() {
	fmt.Println(isAlienSorted2([]string{"hello", "leetcode", "ololo"}, "hlabcdefgijkmnopqrstuvwxyz"))
}

func isAlienSorted(words []string, order string) bool {
	m := make(map[rune]rune)
	for i, b := range order {
		m[b] = rune('a' + byte(i))
	}
	fmt.Printf("m = %v \n", m)

	for i, s := range words {
		runes := []rune(s)
		for i, r := range runes {
			runes[i] = m[r]
		}
		words[i] = string(runes)
		fmt.Printf("i =  %d \n", i)
		fmt.Printf("s =  %s \n", s)
		fmt.Printf("words = %v \n", words)
	}
	return sort.StringsAreSorted(words)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isAlienSorted2(words []string, order string) bool {
	dic := make([]int, 26)
	for i := 0; i <= len(order)-1; i++ {
		fmt.Println(i)
		fmt.Println(order[i] - 'a')
		dic[order[i]-'a'] = i
	}
	fmt.Println(dic)
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			minLen := min(len(words[i]), len(words[j]))
			for k := 0; k < minLen; k++ {
				c1, c2 := words[i][k]-'a', words[j][k]-'a'
				if dic[c1] < dic[c2] {
					break
				} else if dic[c1] > dic[c2] {
					return false
				} else if k == minLen-1 && len(words[i]) > len(words[j]) {
					return false
				}
			}
		}
	}
	return true
}
