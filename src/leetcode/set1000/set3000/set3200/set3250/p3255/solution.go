package p3255

func resultsArray(nums []int, k int) []int {
	n := len(nums)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
	}

	for i := 0; i < n; {
		j := i
		for i < n && nums[i] == nums[j]+i-j {
			if i-j+1 >= k {
				res[i] = nums[i]
			}
			i++
		}
	}

	return res[k-1:]
}
