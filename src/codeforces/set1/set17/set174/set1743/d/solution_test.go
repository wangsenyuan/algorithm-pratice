package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "11010"
	expect := "11111"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1110010"
	expect := "1111110"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0000"
	expect := "0"
	runSample(t, s, expect)
}
