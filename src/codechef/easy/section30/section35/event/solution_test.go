package main

import "testing"

func runSample(t *testing.T, start string, end string, l int, r int, expect string) {
	res := solve(start, end, l, r)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "saturday", "sunday", 2, 4, "2")
}

func TestSample2(t *testing.T) {
	runSample(t, "monday", "wednesday", 1, 20, "many")
}

func TestSample3(t *testing.T) {
	runSample(t, "saturday", "sunday", 3, 5, "impossible")
}
