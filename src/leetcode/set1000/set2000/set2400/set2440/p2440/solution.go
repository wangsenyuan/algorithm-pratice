package p2440

func findMaxK(nums []int) int {
	cnt := make(map[int64]bool)

	for _, num := range nums {
		cnt[int64(num)] = true
	}
	ans := -1
	for _, num := range nums {
		if num > ans {
			if cnt[-int64(num)] {
				ans = num
			}
		}
	}

	return ans
}
