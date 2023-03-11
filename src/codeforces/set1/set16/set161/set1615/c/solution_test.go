package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "100010111"
	B := "101101100"
	expect := 3
	runSample(t, A, B, expect)
}
func TestSample2(t *testing.T) {
	A := "001011011"
	B := "011010101"
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "01001100"
	B := "00111101"
	expect := -1
	runSample(t, A, B, expect)
}
