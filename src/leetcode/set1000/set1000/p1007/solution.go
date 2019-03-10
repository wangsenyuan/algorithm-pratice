package p1007

import "sort"

func minDominoRotations(A []int, B []int) int {
	n := len(A)
	if n <= 1 {
		return 0
	}
	res := make([]int, 4)
	// try make all A equal A[0]
	for i := 1; i < n && res[0] >= 0; i++ {
		if A[i] == A[0] {
			// do nothing
		} else if B[i] == A[0] {
			// swap
			res[0]++
		} else {
			// wrong
			res[0] = -1
		}
	}
	// try make all A equal B[0]
	res[1] = 1
	for i := 1; i < n && res[1] >= 0; i++ {
		if A[i] == B[0] {
			// do nothing
		} else if B[i] == B[0] {
			// swap
			res[1]++
		} else {
			res[1] = -1
		}
	}
	// try make all B equal B[0]
	for i := 1; i < n && res[2] >= 0; i++ {
		if B[i] == B[0] {

		} else if A[i] == B[0] {
			res[2]++
		} else {
			res[2] = -1
		}
	}
	// try make all B equal A[0]
	res[3] = 1
	for i := 1; i < n && res[3] >= 0; i++ {
		if B[i] == A[0] {

		} else if A[i] == A[0] {
			res[3]++
		} else {
			res[3] = -1
		}
	}
	sort.Ints(res)
	for i := 0; i < 4; i++ {
		if res[i] >= 0 {
			return res[i]
		}
	}
	return -1
}
