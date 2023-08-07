package main

import "testing"

func runSample(t *testing.T, d int, expect int) {
	res := solve(d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 1
	expect := 6
	runSample(t, d, expect)
}

func TestSample2(t *testing.T) {
	d := 2
	expect := 15
	runSample(t, d, expect)
}

func TestSample3(t *testing.T) {
	d := 381
	expect := 294527
	runSample(t, d, expect)
}
