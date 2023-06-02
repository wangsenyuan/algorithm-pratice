package main

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := solve(A, K)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	A := []int{4, 1}
	expect := 12
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	k := 0
	A := []int{5, 6}
	expect := 30
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{7, 1, 3}
	expect := 61
	runSample(t, A, k, expect)
}
