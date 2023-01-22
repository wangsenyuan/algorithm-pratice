package p2541

import "sort"

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	// 在某个i上使得 nums1[i] = nums2[j] 那么必然需要在另外一个j上+/-
	// k is fixed
	diff := make([]int, n)
	var sum int64
	for i := 0; i < n; i++ {
		diff[i] = nums1[i] - nums2[i]
		if k == 0 {
			if diff[i] != 0 {
				return -1
			}
		} else if abs(diff[i])%k != 0 {
			return -1
		}
		sum += int64(diff[i])
	}
	if k == 0 {
		return 0
	}
	if sum != 0 {
		return -1
	}

	sort.Ints(diff)
	var res int64
	for i, j := 0, n-1; i < j; {
		for i+1 < j && diff[i] < 0 && diff[j] > 0 && diff[i]+diff[j] != 0 {
			if abs(diff[i]) > abs(diff[j]) {
				res += int64(diff[j] / k)
				diff[i] += diff[j]
				diff[j] = 0
				j--
			} else {
				res -= int64(diff[i] / k)
				diff[j] += diff[i]
				diff[i] = 0
				i++
			}
		}
		if diff[i]+diff[j] != 0 {
			return -1
		}
		res += int64(diff[j] / k)
		i++
		j--
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
