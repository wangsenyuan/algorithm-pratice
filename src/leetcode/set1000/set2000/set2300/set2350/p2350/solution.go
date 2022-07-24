package p2350

func shortestSequence(rolls []int, k int) int {
	// 长度为l的子序列的，都出现的投递的长度为
	// dp[1] = k
	// dp[2] = k * k
	// dp[l] = k ** l
	// 首先l是很小的

	l := 1
	n := len(rolls)

	for i := 0; i < n; {
		num := make(map[int]bool)
		for len(num) < k && i < n {
			num[rolls[i]] = true
			i++
		}
		if len(num) == k {
			l++
		}
	}

	return l
}
