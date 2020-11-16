package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addStrings("10", "2"))
}

func addStrings(num1 string, num2 string) string {
	var res strings.Builder
	carry := byte(0)
	i := len(num1) - 1
	j := len(num2) - 1
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += num1[i] - '0'
			i--
		}
		if j >= 0 {
			sum += num2[j] - '0'
			j--
		}
		carry = sum / 10
		sum = sum % 10
		res.WriteString(string(sum + '0'))
	}
	if carry != 0 {
		res.WriteString(string(carry + '0'))
	}
	reversed := res.String()
	res.Reset()
	for i := len(reversed) - 1; i >= 0; i-- {
		res.WriteString(string(reversed[i]))
	}
	return res.String()
}
