package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println("pick 3: ", c.Pick(3))
	fmt.Println("pick 1: ", c.Pick(1))

}

type Solution struct {
	numz []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	total := 0
	res := 0
	for i := 0; i < len(this.numz); i++ {
		fmt.Println(i)
		if this.numz[i] == target {
			total++
			fmt.Printf("total: %d\n", total)
			fmt.Printf("rand.Intn(total): %d\n", rand.Intn(total))

			if rand.Intn(total) == 0 {
				res = i
			}
			fmt.Printf("res: %d\n", res)
		}
		fmt.Println("---")
	}
	return res
}
