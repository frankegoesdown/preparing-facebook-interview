package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(62))
}

func hammingWeight(num uint32) int {
	var sum int
	for num != 0 {
		sum++
		num &= num - 1
		fmt.Println(num)
	}
	return sum
}
