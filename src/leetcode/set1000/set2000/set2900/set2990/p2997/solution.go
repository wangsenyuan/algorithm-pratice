package p2997

func minOperations(nums []int, k int) int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}

	xor ^= k

	var res int

	for xor > 0 {
		res++
		xor -= xor & -xor
	}
	return res
}
