package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, []int{1, 2, 3}, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, []int{2, 1, 2}, 1)
}
