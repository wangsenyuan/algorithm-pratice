package main

import "testing"

func runSample(t *testing.T, b string, expect string) {
	res := solve(b)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abbaac", "abac")
}

func TestSample2(t *testing.T) {
	runSample(t, "ac", "ac")
}

func TestSample3(t *testing.T) {
	runSample(t, "bccddaaf", "bcdaf")
}

func TestSample4(t *testing.T) {
	runSample(t, "zzzzzzzzzz", "zzzzzz")
}
