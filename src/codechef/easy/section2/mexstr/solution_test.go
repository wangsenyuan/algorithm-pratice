package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(len(s), s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1001011"
	expect := "1100"
	runSample(t, s, expect)
}
