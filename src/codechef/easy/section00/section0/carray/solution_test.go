package main

import "testing"

func runSample(t *testing.T, n int, A []int, k int, b int, expect int) {
	res := solve(n, A, k, b)
	if res != expect {
		t.Errorf("sample %d %v %d %d, expect %d, but got %d", n, A, k, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	k := 4
	b := 1
	A := []int{100, 2, 4, 17, 8}
	runSample(t, n, A, k, b, 3)
}
