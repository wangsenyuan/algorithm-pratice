package main

import "testing"

func runSample(t *testing.T, x []string, expect int) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []string{
		"ssh",
		"hs",
		"s",
		"hhhs",
	}
	expect := 18
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := []string{
		"h", "s",
	}
	expect := 1
	runSample(t, x, expect)
}
