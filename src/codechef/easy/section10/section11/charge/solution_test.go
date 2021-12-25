package main

import "testing"

func runSample(t *testing.T, n int, A []int, T []int) {
	res := solve(n, A, T)

	if len(res) > 2*n {
		t.Errorf("Sample result %v exceeds 2 * n", res)
	} else {
		X := make([]int, n)
		for _, cur := range res {
			i := cur[0] - 1
			L, R := cur[1], cur[2]
			if R > T[i] {
				t.Errorf("Sample result %v, not correct", cur)
			}
			X[i] += R - L
		}

	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{31, 32, 6, 13, 7}
	T := []int{14, 50, 34, 4, 31}
	runSample(t, n, A, T)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{43, 26, 22, 11, 30}
	T := []int{26, 4, 41, 46, 49}
	runSample(t, n, A, T)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{36, 40, 49, 19, 37}
	T := []int{18, 11, 48, 15, 33}
	runSample(t, n, A, T)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{16, 3, 24, 21, 21}
	T := []int{24, 31, 36, 49, 50}
	runSample(t, n, A, T)
}
