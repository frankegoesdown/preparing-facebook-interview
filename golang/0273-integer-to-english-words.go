package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(123))
}

var lessThan21 = []string{
	"",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
	"Twenty",
}

var ten = []string{
	"",
	"",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var thousand = []string{
	"",
	"Thousand",
	"Million",
	"Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""
	i := 0

	for num > 0 {
		if num%1000 != 0 {
			fmt.Println("mod:", num%1000)
			res = lessK(num%1000) + thousand[i] + " " + res
		}

		num /= 1000
		i++
	}

	return strings.TrimRight(res, " ")
}

func lessK(num int) string {
	fmt.Println(num)
	if num == 0 {
		return ""
	}

	if num <= 20 {
		return lessThan21[num] + " "
	}

	if num < 100 {
		return ten[num/10] + " " + lessK(num%10)
	}
	fmt.Println("===")
	return lessThan21[num/100] + " Hundred " + lessK(num%100)
}
