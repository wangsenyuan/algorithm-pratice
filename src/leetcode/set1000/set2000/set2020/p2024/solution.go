package p2024

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)

	T := make([]int, n+1)

	var best int

	for i, j := 0, 0; i < n; i++ {
		T[i+1] = T[i]
		if answerKey[i] == 'T' {
			T[i+1]++
		}

		for j < i && T[i+1]-T[j]+k < i+1-j {
			j++
		}
		best = max(best, i+1-j)
	}

	for i, j := 0, 0; i < n; i++ {
		T[i+1] = T[i]
		if answerKey[i] == 'F' {
			T[i+1]++
		}
		for j < i && T[i+1]-T[j]+k < i+1-j {
			j++
		}
		best = max(best, i+1-j)
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
