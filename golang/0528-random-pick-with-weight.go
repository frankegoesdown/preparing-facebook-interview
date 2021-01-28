package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	w := []int{1, 3}
	sol := Constructor528(w)
	fmt.Printf("1.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("2.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("3.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("4.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("5.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("6.PickIndex = %v\n", sol.PickIndex())
}

type Solution struct {
	sums []int
}

func Constructor528(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	sums := make([]int, len(w))
	sums[0] = w[0]
	for i := 1; i < len(w); i++ {
		sums[i] = w[i] + sums[i-1]
	}
	fmt.Println(sums)
	return Solution{sums}
}

func (this *Solution) PickIndex() int {
	//fmt.Println(this.sums[len(this.sums)-1])
	pick := rand.Intn(this.sums[len(this.sums)-1])
	fmt.Println(pick)
	l, r := 0, len(this.sums)-1
	for l <= r {
		mid := l + (r-l)/2
		if this.sums[mid] > pick {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
