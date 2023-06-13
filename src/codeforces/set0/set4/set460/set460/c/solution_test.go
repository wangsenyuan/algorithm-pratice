package main

import "testing"

func runSample(t *testing.T, A []int, w int, m int, expect int) {
	res := solve(A, w, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2, 2, 2, 1, 1}
	m, w := 2, 3
	expect := 2
	runSample(t, A, w, m, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 8}
	m, w := 5, 1
	expect := 9
	runSample(t, A, w, m, expect)
}
