package main

func main() {

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

func nextPermutation3(nums []int) {
	n := len(nums)
	k := -1
	l := -1
	for k = n - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}
	if k == -1 {
		reverse(&nums, 0, n-1)
	} else {
		for l = n-1; l > k; l-- {
			if nums[l] > nums[k] {
				break
			}
		}
		swap(&nums, k, l)
		reverse(&nums, k+1, n-1)
	}
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
