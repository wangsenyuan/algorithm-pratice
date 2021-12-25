package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 3, 3}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 1, 2}
	expect := 10
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 2, 1}
	expect := 16
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 2, 3, 2, 3, 3, 1}
	expect := 54
	runSample(t, A, expect)
}
