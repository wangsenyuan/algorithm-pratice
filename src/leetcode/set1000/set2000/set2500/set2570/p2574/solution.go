package p2574

func leftRigthDifference(nums []int) []int {
	n := len(nums)
	R := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		R[i] = nums[i]
		if i+1 < n {
			R[i] += R[i+1]
		}
	}

	var L int

	res := make([]int, n)

	for i := 0; i < n; i++ {
		L += nums[i]
		res[i] = abs(L - R[i])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
