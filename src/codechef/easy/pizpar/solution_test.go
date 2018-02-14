package main

import "testing"

func runSample(t *testing.T, n int, A []int, m int, expect int64) {
	res := solve(n, A, m)

	if res != expect {
		t.Errorf("sample %d %v %d, expect %d, but got %d", n, A, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 10
	A := []int{1, 2, 3, 4, 5}
	runSample(t, n, A, m, 31)
}
