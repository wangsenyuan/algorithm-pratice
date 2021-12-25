package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect bool) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, k, A, expect, res)
	}
	if expect {
		for i := 1; i < n; i++ {
			if A[i] == A[i-1] {
				t.Fatalf("Sample %d %d result %v, not correct", n, k, A)
			}
		}
	}
} 

func TestSample1(t *testing.T) {
	n, k := 5, 3
	A := []int{-1, -1, -1, -1, -1}
	runSample(t, n, k, A, true)
}
func TestSample2(t *testing.T) {
	n, k := 4, 5
	A := []int{1, 2, 3, 5}
	runSample(t, n, k, A, true)
}
func TestSample3(t *testing.T) {
	n, k := 3, 2
	A := []int{1, -1, 2}
	runSample(t, n, k, A, false)
}

func TestSample4(t *testing.T) {
	n, k := 4, 2
	A := []int{1, -1, -1, 1}
	runSample(t, n, k, A, false)
}

func TestSample5(t *testing.T) {
	n, k := 6, 4
	A := []int{-1, -1, 4, -1, 2, -1}
	runSample(t, n, k, A, true)
}