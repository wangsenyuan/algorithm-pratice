package main

import "testing"

func runSample(t *testing.T, A, B, C int, expect int) {
	res := solve(A, B, C)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B, C := 1, 2, 3
	expect := 4
	runSample(t, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	A, B, C := 0, 0, 2
	expect := 1
	runSample(t, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	A, B, C := 87, 986, 15
	expect := 64
	runSample(t, A, B, C, expect)
}
