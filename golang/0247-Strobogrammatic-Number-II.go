package main

import "fmt"

func main() {
	fmt.Println(findStrobogrammatic(2))
}

func findStrobogrammatic(n int) []string {
	return helper(n, n)
}

func helper(n int, m int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}

	list := helper(n-2, m)

	res := make([]string, 0)

	for i := 0; i < len(list); i++ {
		s := list[i]

		if n != m {
			res = append(res, "0"+s+"0")
		}

		res = append(res, "1"+s+"1")
		res = append(res, "6"+s+"9")
		res = append(res, "8"+s+"8")
		res = append(res, "9"+s+"6")
	}

	return res
}
