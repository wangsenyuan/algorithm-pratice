package main

import "testing"

func runSample(t *testing.T, N, M int, A []int, K int, L int, expect int) {
	res := solve(N, M, A, K, L)
	if res != expect {
		t.Errorf("Sample %d %d %v %d %d, expect %d, but got %d", N, M, A, K, L, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, K, L := 6, 5, 19, 3
	A := []int{4, 8, 16, 12, 14, 18}
	expect := 12
	runSample(t, N, M, A, K, L, expect)
}

func TestSample2(t *testing.T) {
	N, M, K, L := 1, 10, 20, 3
	A := []int{3}
	expect := 16
	runSample(t, N, M, A, K, L, expect)
}

func TestSample3(t *testing.T) {
	N, M, K, L := 5, 2, 6, 10
	A := []int{1, 2, 5, 3, 4}
	expect := 29
	runSample(t, N, M, A, K, L, expect)
}

func TestSample4(t *testing.T) {
	N, M, K, L := 1, 1, 9, 5
	A := []int{8}
	expect := 2
	runSample(t, N, M, A, K, L, expect)
}
