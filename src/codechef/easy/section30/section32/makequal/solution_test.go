package main

import "testing"

func runSample(t *testing.T, A, B, C int, expect bool) {
	res := solve(A, B, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := 4
	B := 4
	C := 3
	expect := true
	runSample(t, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	A := 1
	B := 6
	C := 7
	expect := true
	runSample(t, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	A := 3
	B := 4
	C := 3
	expect := false
	runSample(t, A, B, C, expect)
}

func TestSample4(t *testing.T) {
	A := 2
	B := 4
	C := 3
	expect := true
	runSample(t, A, B, C, expect)
}