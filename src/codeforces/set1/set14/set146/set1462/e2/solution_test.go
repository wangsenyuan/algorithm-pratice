package main

import "testing"

func runSample(t *testing.T, A []int, m int, k int, expect int) {
	res := solve(A, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, k := 3, 2
	A := []int{1, 2, 4, 3}
	expect := 2
	runSample(t, A, m, k, expect)
}

func TestSample2(t *testing.T) {
	m, k := 2, 1
	A := []int{1, 1, 1, 1}
	expect := 6
	runSample(t, A, m, k, expect)
}

func TestSample3(t *testing.T) {
	m, k := 4, 3
	A := []int{5, 6, 1, 3, 2, 9, 8, 1, 2, 4}
	expect := 20
	runSample(t, A, m, k, expect)
}
