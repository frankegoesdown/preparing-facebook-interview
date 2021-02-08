package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	var ans []string
	for _, n := range nums {
		if n > lower {
			if n-1 > lower {
				ans = append(ans, strconv.Itoa(lower)+"->"+strconv.Itoa(n-1))
			} else {
				ans = append(ans, strconv.Itoa(lower))
			}
		}
		lower = n + 1
	}
	if lower <= upper {
		if lower < upper {
			ans = append(ans, strconv.Itoa(lower)+"->"+strconv.Itoa(upper))
		} else {
			ans = append(ans, strconv.Itoa(lower))
		}
	}
	return ans
}
