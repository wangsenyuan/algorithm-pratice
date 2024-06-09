package p3176

func findWinningPlayer(skills []int, k int) int {
	// k比较小时，进行bruteforce
	// 否则肯定是最大的那个
	n := len(skills)
	var cnt int
	for i := 0; i < n; {
		j := i
		i++
		for i < n && skills[j] > skills[i] {
			i++
			cnt++
		}
		if cnt >= k || i == n {
			return j
		}
		cnt = 1
	}
	return -1
}
