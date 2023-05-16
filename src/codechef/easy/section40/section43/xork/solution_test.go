package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 2, 8, 1}
	k := 15
	expect := 4
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 4, 2, 8, 1}
	k := 7
	expect := 1
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 4, 2, 8, 1}
	k := 20
	expect := -1
	runSample(t, A, k, expect)
}
