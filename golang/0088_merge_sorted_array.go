package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m+n && j < n {
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
			continue
		}
		i++
	}
	if j < n {
		copy(nums1[m+j:], nums2[j:])
	}
}
