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
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
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
