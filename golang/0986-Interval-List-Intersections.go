package main

import "fmt"

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
}
func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		start, end := getInterval(A[i], B[j])
		fmt.Println(start)
		fmt.Println(end)
		if start <= end {
			result = append(result, []int{start, end})
		}
		if A[i][1] == end {
			i++
		}
		if B[j][1] == end {
			j++
		}
		fmt.Println(result)
	}
	return result
}

func getInterval(A []int, B []int) (int, int) {
	start := max(A[0], B[0])
	end := min(A[1], B[1])
	return start, end
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
