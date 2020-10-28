package main

import "fmt"

/*
i: 0
j: 4
i: 1
j: 3
[98 98] b,b
true
*/
func main() {
	fmt.Println(validPalindrome("abbca"))
}

func isPal(s []byte) bool {
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func validPalindrome(t string) bool {
	s := []byte(t)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %d \n", i)
		fmt.Printf("j: %d \n", j)
		if s[i] != s[j] {
			return isPal(s[i:j]) || isPal(s[i+1:j+1])
		}
	}
	return true
}
