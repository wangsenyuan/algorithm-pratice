package p976

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)

	n := len(A)

	for i := n - 3; i >= 0; i-- {
		a := A[i]
		b := A[i+1]
		c := A[i+2]
		if a+b > c {
			return a + b + c
		}
	}

	return 0
}
