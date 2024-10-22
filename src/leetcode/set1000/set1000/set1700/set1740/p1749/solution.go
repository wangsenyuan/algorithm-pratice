package p1749

func maxAbsoluteSum(nums []int) int {
	var sum int
	var mn, mx int
	var res int
	for _, x := range nums {
		sum += x
		res = max(res, sum-mn, abs(sum-mx))
		mn = min(mn, sum)
		mx = max(mx, sum)
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}

func maxAbsoluteSum1(nums []int) int {
	a := maxSum(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] = -nums[i]
	}
	b := maxSum(nums)

	return max(a, b)
}

func maxSum(arr []int) int {
	var best int
	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > best {
			best = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return best
}
