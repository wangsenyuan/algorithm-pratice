package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int) {
	res := solve(n, k, A)
	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	A := []int{1, 3}
	expect := 1
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 7, 3
	A := []int{3, 7, 5, 4, 6, 2, 1}
	expect := 0
	runSample(t, n, k, A, expect)
}
