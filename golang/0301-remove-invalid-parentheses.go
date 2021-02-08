package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("(a)())()"))
}

func removeInvalidParentheses(s string) []string {
	var ans = make([]string, 0)
	l, r := minRemoveCount(s)
	fmt.Println(l)
	fmt.Println(r)
	fmt.Println(ans)
	fmt.Println("===========")
	dfs(s, 0, l, r, 0, &ans)
	return ans
}

func dfs(s string, idx, l, r int, count int, ans *[]string) {
	fmt.Println(s)
	fmt.Println(idx)
	fmt.Println(l)
	fmt.Println(r)
	fmt.Println(count)
	fmt.Println("-------")
	if count < 0 {
		return
	}
	if idx == len(s) {
		if count == 0 {
			*ans = append(*ans, s)
		}
		return
	}
	var prevB byte
	for i := idx; i < len(s) && count >= 0; i++ {
		switch s[i] {
		case '(':
			if l > 0 && prevB != '(' {
				fmt.Println("first")
				fmt.Println(s)
				ss := s[:i] + s[i+1:]
				fmt.Println(ss)
				fmt.Println("======")
				dfs(ss, i, l-1, r, count, ans)
			}
			count++
		case ')':
			if r > 0 && prevB != ')' {
				fmt.Println("second")
				fmt.Println(i)
				fmt.Println(l)
				fmt.Println(r)
				ss := s[:i] + s[i+1:]
				fmt.Println(s)
				fmt.Println(ss)
				fmt.Println("=======")
				dfs(ss, i, l, r-1, count, ans)
			}
			count--
		}
		prevB = s[i]
	}

	if count == 0 {
		*ans = append(*ans, s)
	}
	return
}

func minRemoveCount(s string) (int, int) {
	l := 0
	r := 0
	for _, b := range s {
		switch b {
		case '(':
			l++
		case ')':
			l--
			if l < 0 {
				r++
				l = 0
			}
		}
	}
	return l, r
}
