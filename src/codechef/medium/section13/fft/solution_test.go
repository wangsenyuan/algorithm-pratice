package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res  != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{6, 6 ,6 ,6 ,6 ,6}
	expect := 120
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{1, 1, 1, 1, 1, 1}
	expect := 1
	runSample(t, n, A, expect)
}