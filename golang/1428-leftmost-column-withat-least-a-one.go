package main

func main() {

}

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	m, n := dimensions[0], dimensions[1]
	i, j := 0, n-1
	for i < m && j >= 0 {
		if binaryMatrix.Get(i, j) == 0 {
			i++
		} else {
			j--
		}
	}
	if j == n-1 {
		return -1
	}
	return j + 1
}
