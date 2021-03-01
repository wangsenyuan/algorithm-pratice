package p1775

import "sort"

func minOperations(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	if m > n*6 || m*6 < n {
		return -1
	}
	var sum1 int
	for _, num := range nums1 {
		sum1 += num
	}
	var sum2 int
	for _, num := range nums2 {
		sum2 += num
	}
	if sum1 < sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}
	// sum1 >= sum2
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 如果要调整，就应该调整最大和最小的数
	i := len(nums1) - 1
	j := 0
	var res int
	for sum1 > sum2 && i >= 0 && j < len(nums2) {
		diff := sum1 - sum2
		a := nums1[i]
		b := nums2[j]
		// 减少a到1或者增加b到6
		if diff >= a-1 && diff >= 6-b {
			//任何一个都可以
			res++
			if a-1 > 6-b {
				sum1 -= a - 1
				i--
			} else {
				sum2 += 6 - b
				j++
			}
		} else if a-1 > diff {
			res++
			sum1 = sum2
			i--
		} else {
			res++
			sum2 = sum1
			j++
		}
	}
	if sum1 != sum2 {
		for i >= 0 && sum1 > sum2 {
			res++
			sum1 -= nums1[i] - 1
			i--
		}
		for j < len(nums2) && sum1 > sum2 {
			res++
			sum2 += 6 - nums2[j]
			j++
		}
	}
	return res
}
