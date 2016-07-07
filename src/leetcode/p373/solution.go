package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	result := kSmallestPairs(nums1, nums2, 8)
	for _, x := range result {
		fmt.Printf("%v\n", x)
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0, k)

	i, j := 0, 0
	for len(result) < k {
		result = append(result, []int{nums1[i], nums2[j]})

		c := nums1[i] + nums2[j]

		a, b := i, j
		if a < len(nums1)-1 {
			a++
		}

		if b < len(nums2)-1 {
			b++
		}

		for j > 0 && nums1[a]+nums2[j-1] >= c {
			j--
		}

		for i > 0 && nums1[i-1]+nums2[b] >= c {
			i--
		}

		if nums1[a]+nums2[j] <= nums1[i]+nums2[b] {
			i = a
		} else {
			j = b
		}
	}

	return result
}
