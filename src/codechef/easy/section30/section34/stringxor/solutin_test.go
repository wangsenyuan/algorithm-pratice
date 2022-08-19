package main

import "testing"

func runSample(t *testing.T, A, B string, expect bool) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "00"
	B := "01"
	expect := false
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "010"
	B := "000"
	expect := true
	runSample(t, A, B, expect)
}
