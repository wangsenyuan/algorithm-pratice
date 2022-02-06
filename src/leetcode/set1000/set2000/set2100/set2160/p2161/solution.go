package p2161

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, n)
	var x int
	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			res[x] = nums[i]
			x++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] == pivot {
			res[x] = nums[i]
			x++
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > pivot {
			res[x] = nums[i]
			x++
		}
	}
	return res
}
