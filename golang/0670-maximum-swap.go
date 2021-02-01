package main

import "fmt"

func main() {
	fmt.Println(maximumSwap(21736))
}

func maximumSwap(num int) int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	fmt.Println(digits)
	var maxDigit, maxIndex int
	index1, index2 := 0, 0
	for i, digit := range digits {
		if digit > maxDigit {
			maxDigit = digit
			maxIndex = i
		} else if digit < maxDigit {
			index1 = i
			index2 = maxIndex
		}
		fmt.Println(i)
		fmt.Println(digit)
		fmt.Println(index1)
		fmt.Println(index2)
		fmt.Println(maxDigit)
		fmt.Println(maxIndex)
		fmt.Println("==============")
	}

	digits[index1], digits[index2] = digits[index2], digits[index1]

	var res int
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return res
}
