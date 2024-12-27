package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := "01100010"
	b := "00110"
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "1010111110"
	b := "0110"
	expect := 4
	runSample(t, a, b, expect)
}
