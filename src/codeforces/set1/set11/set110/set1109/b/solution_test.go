package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "nolon", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "kinnikkinnik", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "babab", 2)
}
