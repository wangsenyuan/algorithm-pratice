package main

import "testing"

func runSample(t *testing.T, n int, A []int, k int, expect bool) {
	res := solve(n, A, k)

	if res != expect {
		t.Errorf("sample %d %v %d, expect %t, but got %t", n, A, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 1
	A := []int{10, 10}
	runSample(t, n, A, k, true)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	A := []int{1, 2, 2}
	runSample(t, n, A, k, true)
}

func TestSample3(t *testing.T) {
	n, k := 3, 2
	A := []int{1, 2, 3}
	runSample(t, n, A, k, false)
}
