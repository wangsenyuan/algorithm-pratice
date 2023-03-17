package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 20, 1, 25, 30}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 10, 15}
	expect := 4
	runSample(t, A, expect)
}
