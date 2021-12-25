package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	expect := 499122178
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 831870296
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, 2}
	expect := 665496237
	runSample(t, A, expect)
}
