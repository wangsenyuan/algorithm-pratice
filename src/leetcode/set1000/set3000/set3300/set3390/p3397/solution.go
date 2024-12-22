package p3397

import "sort"

func maxDistinctElements(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)
	arr := make([]int, n)
	copy(arr, nums)

	arr[0] -= k

	for i := 1; i < n; i++ {
		arr[i] = max(nums[i]-k, arr[i-1]+1)
		if arr[i] > nums[i]+k {
			arr[i] = nums[i] + k
		}
	}
	var res int
	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			res++
		}
	}
	return res
}
