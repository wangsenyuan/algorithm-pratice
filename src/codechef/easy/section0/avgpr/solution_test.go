package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []int{2, 2}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, []int{2, 1, 3}, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, []int{4, 2, 5, 1, 3, 5}, 7)
}
