package p2808

func minimumSeconds(nums []int) int {
	pos := make(map[int]int)
	ans := make(map[int]int)
	n := len(nums)
	for i := 0; i < 2*n; i++ {
		x := nums[i%n]
		if j, ok := pos[x]; ok {
			if v, has := ans[x]; !has || v < i-j {
				ans[x] = i - j
			}
		}
		pos[x] = i
	}
	best := n
	for _, v := range ans {
		best = min(best, v)
	}

	return best / 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
