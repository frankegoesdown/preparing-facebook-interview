package main

import "fmt"

func main() {
	fmt.Println(numFriendRequests([]int{16, 17, 18}))
}

func numFriendRequests(ages []int) int {
	count := make([]int, 121)

	for _ , age := range ages {
		count[age]++
	}
	var res int
	fmt.Println(count)
	for ageA := 0; ageA <= 120; ageA++ {
		countA := count[ageA]
		fmt.Println(countA)
		for ageB := 0; ageB <= 120; ageB++ {
			countB := count[ageB]
			if ageA/2+7 >= ageB {
				continue
			}
			if ageA < ageB {
				continue
			}
			if ageA < 100 && 100 < ageB {
				continue
			}

			res += countA * countB

			if ageA == ageB {
				res -= countA
			}
		}
	}
	return res
}
