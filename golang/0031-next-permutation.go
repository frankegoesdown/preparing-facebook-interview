package main

import "fmt"

func main() {
	fmt.Println(nextPermutation3([]int{1, 5, 8, 4, 7, 6, 5, 3, 1}))
}

func reverseFrom(slice []int, from int) []int {
	L := len(slice)

	half := from + (L-from)/2
	j := 0
	for i := from; i < half; i++ {
		slice[i], slice[L-1-j] = slice[L-1-j], slice[i]
		j++
	}
	return slice
}

func nextPermutation(slice []int) []int {
	L := len(slice)

	var m int
	var n int

	found := false
	for i := L - 1; i > 0; i-- {
		if slice[i-1] < slice[i] {
			m = i - 1
			found = true
			break
		}
	}

	if !found {
		return reverseFrom(slice, 0)
	}

	for i := L - 1; i > 0; i-- {
		if slice[i] > slice[m] {
			n = i
			break
		}
	}

	// Swap
	slice[m], slice[n] = slice[n], slice[m]
	// Reverse m+1...
	return reverseFrom(slice, m+1)
}

func nextPermutation2(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		val := nums[i]
		j := i + 1
		for ; j < l; j++ {
			if nums[j] > nums[i] {
				nums[i] = nums[j]
				nums[j] = val
				return
			}
		}
		if j == l {
			for k := i; k < l-1; k++ {
				nums[k] = nums[k+1]
			}
			nums[l-1] = val
		}
	}
}

func nextPermutation3(nums []int) []int {
	fmt.Println(nums)
	n := len(nums)
	k := 0
	l := 0
	for k = n - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			fmt.Println(nums[k])
			fmt.Println(nums[k+1])
			fmt.Println("-------")
			break
		}
	}
	fmt.Println(k)
	fmt.Println(n)
	fmt.Println(l)
	if k >= 0 {
		for l = n-1; l > k; l-- {
			if nums[l] > nums[k] {
				break
			}
		}
		swap(&nums, k, l)
	}
		fmt.Println("------------")
	fmt.Printf("k: %d\n", k)
	fmt.Printf("l: %d\n", l)
	fmt.Println(nums)
	reverse(&nums, k+1, n-1)

	return nums
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	fmt.Println("====")
	fmt.Println((*nums)[i])
	fmt.Println((*nums)[j])
	fmt.Println("====")
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
