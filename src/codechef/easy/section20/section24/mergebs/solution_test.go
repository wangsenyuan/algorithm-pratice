package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(len(A), A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "1010"
	B := "0101"
	expect := 6
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "10010"
	B := "11001"
	expect := 12
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "00001"
	B := "11110"
	expect := 4
	runSample(t, A, B, expect)
}
