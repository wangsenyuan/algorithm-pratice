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

	nums1 = removeDup(nums1)
	nums2 = removeDup(nums2)
	//fmt.Printf("after sort nums1 = %v\n", nums1)
	//fmt.Printf("after sort nums2 = %v\n", nums2)
	result := make([]int, 0, len(nums1)+len(nums2))

	for i, j := 0, 0; i < len(nums1); i++ {
		x := nums1[i]

		for j < len(nums2) && nums2[j] < x {
			j++
		}

		if j == len(nums2) {
			break
		}

		if nums2[j] > x {
			continue
		}
		result = append(result, x)
	}

	return result
}

func removeDup(nums []int) []int {
	result := make([]int, 0, len(nums))

	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			continue
		}
		result = append(result, x)
	}

	return result
}
