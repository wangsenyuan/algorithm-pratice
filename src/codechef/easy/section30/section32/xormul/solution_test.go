package main

import "testing"

func runSample(t *testing.T, n int, A, B int, X int) {
	expect := int64(A^X) * int64(B^X)

	x := solve(n, A, B)

	res := int64(A^x) * int64(B^x)

	if expect != res {
		t.Errorf("Sample result %d, gives wrong result %d, expect %d", x, res, expect)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0, 0, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 6, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2, 1, 0)
}
