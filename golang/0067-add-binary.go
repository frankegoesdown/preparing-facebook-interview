package main

import "fmt"

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
func addBinary(a string, b string) string {
	i, j, res := len(a) - 1, len(b) - 1, []byte{}
	var carry byte
	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 { carry += a[i] - '0'; i-- }
		if j >= 0 { carry += b[j] - '0'; j-- }
		res = append(res, carry % 2 + '0')
		carry = carry / 2
	}
	return string(reverse(res))
}

func reverse(s []byte) []byte {
	for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 { s[l], s[r] = s[r], s[l] }
	return s
}
