package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(leastInterval([]byte{'A','A','A','A','A','A','B','C','D','E','F','G'}, 2))
}

func leastInterval(tasks []byte, n int) int {
	tmp, p := make([]int, 26), 1
	for _, v := range tasks {
		tmp[v-'A']++
	}
	fmt.Println(tmp)
	sort.Slice(tmp, func(a, b int) bool { return tmp[a] > tmp[b] })
	fmt.Println(tmp)
	for i := 1; i < 26; i++ {
		if tmp[0] == tmp[i] {
			p++
		}
	}
	return max((n+1)*(tmp[0]-1)+p, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
