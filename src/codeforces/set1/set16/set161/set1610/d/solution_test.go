package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2, 4, 7}
	expect := 10
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{12391240, 103904, 1000000000, 4142834, 12039, 142035823, 1032840, 49932183, 230194823, 984293123}
	expect := 996
	runSample(t, A, expect)
}
