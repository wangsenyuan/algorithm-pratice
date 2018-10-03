package p915

func partitionDisjoint(A []int) int {
	n := len(A)
	right := make([]int, n)

	right[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = A[i]
		if right[i] > right[i+1] {
			right[i] = right[i+1]
		}
	}

	var maxSoFar int

	for i := 0; i < n-1; i++ {
		if A[i] > maxSoFar {
			maxSoFar = A[i]
		}
		if maxSoFar <= right[i+1] {
			return i + 1
		}
	}

	return -1
}
