package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	m := [128]int{}
	for _, v := range t {
		m[v]++
	}
	start, end, head, counter := 0, 0, 0, 0
	minLen := math.MaxInt32
	for end < len(s) {
		if m[s[end]] > 0 {
			counter++
		}
		m[s[end]]--
		end++
		for counter == len(t) {
			if end-start < minLen {
				minLen = end - start
				head = start
			}
			m[s[start]]++
			if m[s[start]] == 1 {
				counter--
			}
			start++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[head : head+minLen]
}
