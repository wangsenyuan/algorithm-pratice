package p1982

import (
	"sort"
)

func recoverArray(n int, S []int) []int {

	A := make([]int, n)

	sort.Ints(S)

	for i := n; i >= 2; i-- {
		d := S[1] - S[0]
		var l, r int
		X := make([]int, 0, len(S)/2)
		Y := make([]int, 0, len(S)/2)
		used := make([]bool, 1<<i)
		for {
			for l < 1<<i && used[l] {
				l++
			}
			if l == 1<<i {
				break
			}
			X = append(X, S[l])
			used[l] = true
			for used[r] || S[r] != S[l]+d {
				r++
			}
			Y = append(Y, S[r])
			used[r] = true
		}
		j := sort.SearchInts(X, 0)
		if j < len(X) && X[j] == 0 {
			S = X
			A[i-1] = d
		} else {
			A[i-1] = -d
			S = Y
		}
	}

	A[0] = S[0] + S[1]

	return A
}

func getAllSubSum(A []int, X []int) {
	l := len(X)
	n := len(A)
	for state := 0; state < l; state++ {
		X[state] = 0
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				X[state] += A[i]
			}
		}
	}

}
