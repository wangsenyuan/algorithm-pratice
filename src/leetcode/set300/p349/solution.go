package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersection(nums1, nums2)
	fmt.Printf("%v\n", result)
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	//fmt.Printf("after sort nums1 = %v\n", nums1)
	//fmt.Printf("after sort nums2 = %v\n", nums2)
	res := make([]int, 0, len(nums1))

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if i > 0 && nums1[i] == nums1[i-1] {
			i++
			continue
		}

		if j > 0 && nums2[j] == nums2[j-1] {
			j++
			continue
		}

		a, b := nums1[i], nums2[j]

		if a == b {
			res = append(res, a)
			i, j = i+1, j+1
			continue
		}

		if a > b {
			j++
			continue
		}
		i++
	}

	return res
}

func intersection1(nums1 []int, nums2 []int) []int {
	cnt := make(map[int]bool)

	for _, num := range nums1 {
		cnt[num] = true
	}

	res := make([]int, 0, len(nums2))

	for _, num := range nums2 {
		if cnt[num] {
			res = append(res, num)
			cnt[num] = false
		}
	}

	return res
}
