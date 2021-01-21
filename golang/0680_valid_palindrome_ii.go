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


func validPalindrome2(s string) bool {

	l := 0
	r := len(s)-1

	for l < r {
		if s[l] != s[r] {
			return helper(s, l+1, r) || helper(s, l, r-1)
		} else {
			l++
			r--
		}
	}
	return true
}

func helper(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] { return false }
		l++
		r--
	}
	return true
}
