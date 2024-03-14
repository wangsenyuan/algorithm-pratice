package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "+-+"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "--+---+"
	expect := 7
	runSample(t, s, expect)
}
