package p2364

func countBadPairs(nums []int) int64 {
	n := len(nums)

	tot := int64(n) * int64(n-1) / 2

	cnt := make(map[int]int)

	var good int64

	for i := 0; i < n; i++ {
		x := nums[i] - i
		good += int64(cnt[x])
		cnt[x]++
	}

	return tot - good
}
