package main

import "testing"

func runSample(t *testing.T, n int, A []int, m int, B []int, expect int64) {
	res := solve(n, A, m, B)
	if res != expect {
		t.Errorf("sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample(t *testing.T) {
	n, m := 8, 3
	A := []int{3, 7, 5, 4, 9, 6, 1, 3}
	B := []int{6, 3, 1}
	runSample(t, n, A, m, B, 20)
}
