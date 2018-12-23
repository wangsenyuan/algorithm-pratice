package p961

func repeatedNTimes(A []int) int {
	n := len(A)

	i := 1

	for i < n {
		if A[i] == A[i-1] {
			return A[i]
		}
		if i > 1 && A[i] == A[i-2] {
			return A[i]
		}

		if i > 2 && A[i] == A[i-3] {
			return A[i]
		}

		i++
	}
	return 0
}
