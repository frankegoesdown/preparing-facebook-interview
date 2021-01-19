package main

import "fmt"

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}

func minRemoveToMakeValid(s string) string {
	stack := []byte{}
	open := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if open > 0 {
				open--
				stack = append(stack, s[i])
			}
		} else if s[i] == '(' {
			open++
			stack = append(stack, s[i])
		} else {
			stack = append(stack, s[i])
		}
		fmt.Printf("open: %d \n", open)
		fmt.Printf("stack: %v \n", stack)

	}
	fmt.Println("================")

	res := make([]byte, len(stack)-open)
	p := len(res) - 1
	fmt.Println(res)
	fmt.Println(p)
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == '(' {
			if open > 0 {
				open--
				continue
			}
			res[p] = '('
			p--
		} else {
			res[p] = stack[i]
			p--
		}
		fmt.Printf("i: %d \n", i)
		fmt.Printf("p: %d \n", p)
		fmt.Printf("open: %d \n", open)
		fmt.Printf("res: %v \n", res)
	}
	return string(res)
}
