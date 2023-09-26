package main

import "testing"

func runSample(t *testing.T, n int, k int, P []int, S []int, expect int) {
	res := solve(n, k, P, S)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 4
	P := []int{1, 2, 1, 3}
	S := []int{6, 2, 1, 5, 7}
	expect := 54
	runSample(t, n, k, P, S, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 3
	P := []int{1, 2, 1, 3}
	S := []int{6, 6, 1, 4, 10}
	expect := 56
	runSample(t, n, k, P, S, expect)
}
