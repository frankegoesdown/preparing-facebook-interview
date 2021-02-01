package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isStrobogrammatic("6889"))
}

func isStrobogrammatic(num string) bool {
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		if !strings.Contains("00 11 88 69 96", string(num[i])+string(num[j])) {
			return false
		}
	}
	return true
}
