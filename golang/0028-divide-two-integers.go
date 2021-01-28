package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
}

func divide(dividend int, divisor int) int {
	sign := dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0
	dividend, divisor = abs(dividend), abs(divisor)
	currentDivisor, finalQuotient, partialQuotient := divisor, 0, 1
	for dividend-currentDivisor >= 0 {
		finalQuotient, dividend = finalQuotient+partialQuotient, dividend-currentDivisor
		fmt.Printf("currentDivisor %d \n", currentDivisor)
		fmt.Printf("partialQuotient %d \n", partialQuotient)
		partialQuotient, currentDivisor = partialQuotient<<1, currentDivisor<<1
		fmt.Printf("partialQuotient %d \n", partialQuotient)
		fmt.Printf("currentDivisor %d \n", currentDivisor)
		if dividend-currentDivisor < 0 {
			partialQuotient, currentDivisor = 1, divisor
		}
		fmt.Printf("dividend %d \n", dividend)
		fmt.Printf("partialQuotient %d \n", partialQuotient)
		fmt.Printf("currentDivisor %d \n", currentDivisor)
		fmt.Printf("finalQuotient %d \n", finalQuotient)

		fmt.Println("--------")
	}
	fmt.Println("=============")
	fmt.Println(dividend)
	fmt.Println(currentDivisor)
	fmt.Println(partialQuotient)
	if sign {
		finalQuotient = -finalQuotient
	}
	if finalQuotient > int(math.Pow(2, 31)-1) {
		finalQuotient = int(math.Pow(2, 31) - 1)
	}
	if finalQuotient < -int(math.Pow(2, 31)) {
		finalQuotient = -int(math.Pow(2, 31))
	}
	return finalQuotient
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
