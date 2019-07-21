package p5133

import "sort"

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	n := len(arr1)
	A := make([]int, n)
	B := make([]int, n)
	C := make([]int, n)
	D := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = arr1[i] + arr2[i] + i
		B[i] = -arr1[i] + arr2[i] + i
		C[i] = arr1[i] - arr2[i] + i
		D[i] = -arr1[i] - arr2[i] + i
	}

	var ans int
	prev := make([]int, 4)

	for i := 1; i < n; i++ {
		a := A[i] - A[prev[0]]
		b := B[i] - B[prev[1]]
		c := C[i] - C[prev[2]]
		d := D[i] - D[prev[3]]

		ans = max5(a, b, c, d, ans)

		if A[i] < A[prev[0]] {
			prev[0] = i
		}
		if B[i] < B[prev[1]] {
			prev[1] = i
		}

		if C[i] < C[prev[2]] {
			prev[2] = i
		}

		if D[i] < D[prev[3]] {
			prev[3] = i
		}
	}

	return ans
}

func max5(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	sort.Ints(arr)
	return arr[4]
}
