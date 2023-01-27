package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "11"
	B := "00"
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "10101"
	B := "10101"
	expect := 32
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "101"
	B := "011"
	expect := 4
	runSample(t, A, B, expect)
}
