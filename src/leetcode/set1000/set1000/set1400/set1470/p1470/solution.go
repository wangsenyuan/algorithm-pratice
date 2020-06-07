package p1470

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)

	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}

	return res
}
