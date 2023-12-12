package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 4, 5}
	k := 2
	expect := 3
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{7, 8, 9, 10}
	k := 3
	expect := 13
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 10, 20, 21}
	k := 0
	expect := 1
	runSample(t, A, k, expect)
}
