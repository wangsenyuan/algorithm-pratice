package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 4}
	runSample(t, n, A, 2)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{2, 3, 5}
	runSample(t, n, A, 3)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{2, 3, 3, 3, 3, 1}
	runSample(t, n, A, 2)
}
