package p2874

func minOperations(nums []int, k int) int {
	// k <= 50
	var flag uint64

	var target uint64 = (1 << k) - 1

	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		if x <= k {
			flag |= (1 << (x - 1))
		}
		if flag == target {
			return n - i
		}
	}
	return -1
}
