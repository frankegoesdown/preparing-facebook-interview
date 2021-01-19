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
	fmt.Println(m)
	fmt.Println("ADOBECODEBANC")
	start, end, head, counter := 0, 0, 0, 0
	minLen := math.MaxInt32
	for end < len(s) {

		if m[s[end]] > 0 {
			counter++
		}
		m[s[end]]--
		fmt.Printf("string(s[end]) %s \n", string(s[end]))
		fmt.Printf("m[s[end]] %d \n", m[s[end]])
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
		fmt.Printf("start %d \n", start)
		fmt.Printf("end %d \n", end)
		fmt.Printf("counter %d \n", counter)
		fmt.Printf("head %d \n", head)
		fmt.Printf("minLen %d \n", minLen)
		fmt.Println(m)
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[head : head+minLen]
}
