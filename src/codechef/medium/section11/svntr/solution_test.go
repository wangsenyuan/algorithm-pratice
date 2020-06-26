package main

import "testing"

func runSample(t *testing.T, N, M, s int, A [][]int, expect int64) {
	res := solve(N, M, s, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, s := 2, 3, 7
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expect := int64(11)
	runSample(t, N, M, s, A, expect)
}

func TestSample2(t *testing.T) {
	N, M, s := 3, 3, 12
	A := [][]int{
		{3, 3, 3},
		{4, 4, 4},
		{3, 3, 3},
	}
	expect := int64(27)
	runSample(t, N, M, s, A, expect)
}
