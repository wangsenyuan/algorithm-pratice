package main

import "testing"

func runSample(t *testing.T, n int, B []int) {
	A := make([]int, n)
	copy(A, B)

	solve(n, A)

	// B is A now

	for i := 0; i < n; i++ {
		j := B[i]
		j--
		if A[j]%A[i] != 0 {
			t.Fatalf("Sample result %v, not correct", B)
		}
		for k := j + 1; k < n; k++ {
			if A[k]%A[i] == 0 {
				t.Fatalf("Sample result %v, not correct", B)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	B := []int{5, 2, 3, 4, 5}
	runSample(t, n, B)
}

func TestSample2(t *testing.T) {
	n := 4
	B := []int{4, 4, 4, 4}
	runSample(t, n, B)
}
