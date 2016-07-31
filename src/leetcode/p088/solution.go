package main

import "fmt"

func main() {
	nums1 := make([]int, 4)
	nums1[0] = 1
	nums1[1] = 3
	nums2 := []int{2, 4}
	merge(nums1, 2, nums2, 2)
	fmt.Printf("%v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}

	for n > 0 {
		nums1[i] = nums2[n-1]
		i--
		n--
	}
}
