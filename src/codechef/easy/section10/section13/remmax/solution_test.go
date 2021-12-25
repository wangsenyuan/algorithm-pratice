package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "23", "19")
	runSample(t, "7", "7")
	runSample(t, "100", "99")
	runSample(t, "1123", "1099")
	runSample(t, "101", "101")

	runSample(t, "102", "102")
}
