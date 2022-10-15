package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 4, 6, 1}
	expect := 19
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{108, 24, 18}
	expect := 258
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{93179, 93187, 93199, 93229, 93239, 93241, 93251, 93253, 93257, 93263}
	expect := 392286903
	runSample(t, A, expect)
}
