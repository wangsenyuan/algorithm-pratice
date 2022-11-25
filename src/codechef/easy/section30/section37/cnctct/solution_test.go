package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 4, 4, 4}
	expect := 12
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 6}
	expect := 6
	runSample(t, A, expect)
}
