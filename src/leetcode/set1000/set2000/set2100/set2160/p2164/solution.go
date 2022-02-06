package p2164

import "sort"

func sortEvenOdd(nums []int) []int {
	n := len(nums)
	even := make([]int, (n+1)/2)
	odd := make([]int, n-len(even))
	for i := 0; i < n; i += 2 {
		even[i/2] = nums[i]
		if i+1 < n {
			odd[i/2] = nums[i+1]
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	reverse(odd)
	for i := 0; i < n; i += 2 {
		nums[i] = even[i/2]
		if i+1 < n {
			nums[i+1] = odd[i/2]
		}
	}
	return nums
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
