package main

import "testing"

func runSample(t *testing.T, n int, A, B string, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := "101"
	B := "010"
	expect := 2
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := "0000"
	B := "1111"
	expect := 0
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	A := "101001"
	B := "001010"
	expect := 4
	runSample(t, n, A, B, expect)
}
