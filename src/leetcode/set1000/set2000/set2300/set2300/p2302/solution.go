package p2302

func countSubarrays(nums []int, k int64) int64 {
	// sum(A) * len(A) < k
	n := len(nums)
	var sum int64
	var l int64
	var res int64
	for i, j := 0, 0; i < n; i++ {
		sum += int64(nums[i])
		l++
		for sum*l >= k || sum*l < 0 {
			sum -= int64(nums[j])
			l--
			j++
		}
		res += int64(i - j + 1)
	}
	return res
}
