package p2038

func winnerOfGame(colors string) bool {
	n := len(colors)
	cnt := make([]int, 2)

	for i := 0; i < n; {
		j := i
		for i < n && colors[i] == colors[j] {
			i++
		}
		cnt[int(colors[j]-'A')] += max(0, i-j-2)
	}
	return cnt[0] > cnt[1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
