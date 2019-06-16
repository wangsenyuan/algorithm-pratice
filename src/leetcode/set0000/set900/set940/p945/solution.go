package p945

import "sort"

func minIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)
	n := len(A)

	var res int

	var baseLine = A[0]
	var j int
	for i := 0; i < n; i++ {
		if A[i] <= baseLine+j {
			res += baseLine + j - A[i]
			j++
		} else {
			baseLine = A[i]
			j = 1
		}
	}

	return res
}
