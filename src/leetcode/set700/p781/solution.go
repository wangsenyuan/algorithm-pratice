package p781

import "sort"

func numRabbits(answers []int) int {
	sort.Ints(answers)

	var res int
	n := len(answers)
	for i, j := 0, 0; i <= n; i++ {
		if i < n && answers[i] == answers[j] && j+answers[j] >= i {
			continue
		}
		res += max(i-j, answers[j]+1)
		j = i
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
