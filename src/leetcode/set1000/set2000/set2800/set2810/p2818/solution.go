package p2818

func maxSum(nums []int) int {
	n := len(nums)
	ans := -1

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if getLargestDigit(nums[i]) == getLargestDigit(nums[j]) {
				ans = max(ans, nums[i]+nums[j])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getLargestDigit(num int) int {
	ans := num % 10
	for num > 0 {
		r := num % 10
		if r > ans {
			ans = r
		}
		num /= 10
	}
	return ans
}
