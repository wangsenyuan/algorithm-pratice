package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{3, 5, 7}
	runSample(t, n, A, 8)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{2, 4, 6}
	runSample(t, n, A, 10)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{5, 5}
	runSample(t, n, A, 3)
}
