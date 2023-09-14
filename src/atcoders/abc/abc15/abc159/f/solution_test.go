package main

import "testing"

func runSample(t *testing.T, n int, s int, A []int, expect int) {
	res := solve(n, s, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, s := 3, 4
	A := []int{2, 2, 4}
	expect := 5
	runSample(t, n, s, A, expect)
}

func TestSample2(t *testing.T) {
	n, s := 10, 10
	A := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	expect := 152
	runSample(t, n, s, A, expect)
}
