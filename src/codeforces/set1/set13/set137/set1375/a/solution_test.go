package main

import "testing"

func runSample(t *testing.T, n int, A []int) {
	solve(n, A)

	var a, b int

	for i := 1; i < n; i++ {
		if A[i-1] < A[i] {
			a++
		} else if A[i-1] > A[i] {
			b++
		} else {
			a++
			b++
		}
	}

	if a < n/2 || b < n/2 {
		t.Errorf("Sample result %v not correct", A)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-2, 4, 3}
	runSample(t, len(A), A)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	runSample(t, len(A), A)
}

func TestSample3(t *testing.T) {
	A := []int{-2, 4, 7, -6, 4}
	runSample(t, len(A), A)
}

func TestSample4(t *testing.T) {
	A := []int{9, 7, -4, -2, 1, -3, 9, -4, -5}
	runSample(t, len(A), A)
}

func TestSample5(t *testing.T) {
	A := []int{4, -1, -9, -4, -8, -9, -5, -1, 9}
	runSample(t, len(A), A)
}
