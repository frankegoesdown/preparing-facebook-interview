package main

import "fmt"

func main() {
	fmt.Println(calculate("22-3*10/2"))
}

func calculate(s string) int {
	ans, cur, last := 0, 0, 0
	op := '+'
	for i, ch := range s {
		if ch-'0' >= 0 && ch-'0' <= 9 {
			fmt.Println("cur before: ", cur)
			fmt.Println("ch: ", string(ch))
			cur = cur*10 + int(ch-'0')
			fmt.Println("cur after: ", cur)
		}
		fmt.Println(string(ch))
		fmt.Println(cur)
		fmt.Println("=======")
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			switch op {
			case '+':
				ans += last
				last = cur
			case '-':
				ans += last
				last = -cur
			case '*':
				last = last * cur
			case '/':
				last = last / cur
			}
			op = ch
			cur = 0
		}
	}
	ans += last
	return ans
}
