package main

import "testing"

func runSample(t *testing.T, n int, A []int64, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int64{4, 5, 6, 4}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int64{121, 122, 121, 126}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	A := []int64{1, 1, 1, 1, 1, 1}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []int64{0, 0, 0, 0}
	expect := 8
	runSample(t, n, A, expect)
}
