package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1111"
	expect := "1000"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "010010"
	expect := "010000"
	runSample(t, s, expect)
}
