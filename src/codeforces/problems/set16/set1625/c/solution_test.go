package main

import "testing"

func runSample(t *testing.T, N int, L int, K int, X []int, S []int, expect int) {
	res := int(solve(L, N, K, X, S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, L, K := 4, 10, 0
	X := []int{0, 3, 4, 8}
	A := []int{5, 8, 3, 6}
	expect := 47
	runSample(t, N, L, K, X, A, expect)
}


func TestSample2(t *testing.T) {
	N, L, K := 4, 10, 2
	X := []int{0, 3, 4, 8}
	A := []int{5, 8, 3, 6}
	expect := 38
	runSample(t, N, L, K, X, A, expect)
}
