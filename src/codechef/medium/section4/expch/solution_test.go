package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "((())", 866666673)
}

func TestPow(t *testing.T) {
	if pow(3, 3) != 27 {
		t.Errorf("pow not correct")
	}
}
