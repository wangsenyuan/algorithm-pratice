package p1829

func getMaximumXor(nums []int, maximumBit int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	x := 1<<maximumBit - 1
	n := len(nums)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ans[n-1-i] = xor ^ x
		xor ^= nums[i]
	}

	return ans
}
