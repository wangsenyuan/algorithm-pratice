package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Errorf("Sample %s %d, expect %s, but got %s", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "helowrd", 0, "abcfgij")
}

func TestSample2(t *testing.T) {
	runSample(t, "background", 0, "efhijlmpqs")
}

func TestSample3(t *testing.T) {
	runSample(t, "abcdefghijklmnopqrstuvwxyz", 0, "NOPE")
}
