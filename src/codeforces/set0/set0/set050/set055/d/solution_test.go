package main

import "testing"

func runSample(t *testing.T, l, r string, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1", "9", 9)
}

func TestSample2(t *testing.T) {
	runSample(t, "12", "15", 2)
}
