package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "110"
	expect := alice
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "00"
	expect := bob
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1010"
	expect := alice
	runSample(t, s, expect)
}
