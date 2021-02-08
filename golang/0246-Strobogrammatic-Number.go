package main

import (
	"fmt"
)

func main() {
	fmt.Println(isStrobogrammatic("6889"))
}
func isStrobogrammatic(num string) bool {
	m := map[byte]byte{
		'8': '8',
		'1': '1',
		'6': '9',
		'9': '6',
		'0': '0',
	}

	l, h := 0, len(num)-1

	for l <= h {
		if _, ok := m[num[l]]; !ok {
			return false
		}

		if m[num[l]] != num[h] {
			return false
		}
		l++
		h--
	}
	return true
}
