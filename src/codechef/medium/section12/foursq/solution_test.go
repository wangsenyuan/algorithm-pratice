package main

import "testing"

func runSample(t *testing.T, n int, A []int, p int64, Q [][]int) {
	solver := NewSolver(n, A, p)
	var j int
	for i := 0; i < len(Q); i++ {
		cur := Q[i]
		if cur[0] == 1 {
			p := cur[1]
			v := cur[2]
			solver.Update(p, v)
			A[p-1] = v
		} else {
			l, r := cur[1], cur[2]
			res := solver.Query(l, r)

			sum := ((res[0]*res[0])%p + (res[1]*res[1])%p + (res[2]*res[2])%p + (res[3]*res[3])%p) % p

			var prod int64 = 1

			for i := l - 1; i < r; i++ {
				prod = (prod * int64(A[i])) % p
			}

			if prod != sum {
				t.Errorf("Sample result %v, not correct", res)
				return
			}

			j++
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	var p int64 = 100
	A := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{2, 1, 5},
		{2, 1, 3},
		{1, 1, 4},
		{2, 1, 3},
	}

	runSample(t, n, A, p, Q)
}
