package main

import "fmt"

func main() {
	fmt.Println(multiply([][]int{{1, 0, 0}, {-1, 0, 3}}, [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}

func multiply(A [][]int, B [][]int) [][]int {
	// edge case
	if len(A) == 0 {
		return [][]int{}
	}

	// row is frist matrix
	ans := make([][]int, len(A))
	for i := range ans {
		// column is from second matrix
		ans[i] = make([]int, len(B[0]))
	}

	AFullZeroRows := make([]bool, len(A))
ALoop:
	for idx, row := range A {
		for _, val := range row {
			if val != 0 {
				continue ALoop
			}
		}
		AFullZeroRows[idx] = true
	}

	BFullZeroCol := make([]bool, len(B[0]))
BLoop:
	for col := 0; col < len(B[0]); col++ {
		for row := 0; row < len(B); row++ {
			if B[row][col] != 0 {
				continue BLoop
			}
		}
		BFullZeroCol[col] = true
	}

	for rowIdx, aRow := range A {
		// skip this row completely
		if AFullZeroRows[rowIdx] {
			continue
		}
		for colIdx := 0; colIdx < len(B[0]); colIdx++ {
			// skip this column for this row
			if BFullZeroCol[colIdx] {
				continue
			}
			// calculate result for this postion
			sum := 0
			for i := 0; i < len(A[0]); i++ {
				sum = sum + aRow[i]*B[i][colIdx]
			}
			ans[rowIdx][colIdx] = sum
		}
	}

	return ans
}
