package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 5, 4, 7, 10, 12}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{536870912, -536870912}
	expect := 2
	runSample(t, A, expect)
}
