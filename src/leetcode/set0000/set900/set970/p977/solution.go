package p977

import "sort"

func sortedSquares(A []int) []int {
	n := len(A)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[i] = A[i] * A[i]
	}
	sort.Ints(B)
	return B
}
