package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPal("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if !isChar(s[l]) {
			l++
			continue
		}
		if !isChar(s[r]) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isChar(c byte) (res bool) {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		res = true
	}
	return
}
