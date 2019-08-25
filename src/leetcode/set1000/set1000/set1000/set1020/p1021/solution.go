package p1021

func maxScoreSightseeingPair(A []int) int {
	// A[i] + A[j] + i - j => (A[i] + i) + (A[j] - j)
	n := len(A)
	// find the least at right of i
	max := A[n-1] - (n - 1)
	best := -(1 << 30)
	for i := n - 2; i >= 0; i-- {
		cur := A[i] + i + max
		if cur > best {
			best = cur
		}
		if A[i]-i > max {
			max = A[i] - i
		}
	}

	return best
}
