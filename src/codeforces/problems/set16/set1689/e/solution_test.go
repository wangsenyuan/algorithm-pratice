package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	} else {
		if !connected(A) {
			t.Errorf("Sample result %v, not correct", A)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 12}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 0, 0, 0}
	expect := 3
	runSample(t, A, expect)
}
