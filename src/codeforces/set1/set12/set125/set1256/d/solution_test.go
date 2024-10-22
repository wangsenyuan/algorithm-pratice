package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(k, s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 5
	s := "11011010"
	expect := "01011110"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	k := 9
	s := "1111100"
	expect := "0101111"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	k := 11
	s := "1111100"
	expect := "0011111"
	runSample(t, s, k, expect)
}
