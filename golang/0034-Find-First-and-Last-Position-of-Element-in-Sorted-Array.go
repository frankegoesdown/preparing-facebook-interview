package main

func main() {

}

func searchRange(nums []int, target int) []int {
	l := extremeInsertionIndex(nums, target, true)
	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}
	r := extremeInsertionIndex(nums, target, false) -1
	return []int{l, r}
}

func extremeInsertionIndex(nums []int, target int, left bool) int {
	l := 0
	h := len(nums)
	for l < h {
		mid := (h-l) /2 + l
		if nums[mid] > target || left && nums[mid] == target {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
