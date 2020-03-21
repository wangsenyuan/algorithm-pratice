package main

import "testing"

func runSample(t *testing.T, ss []string, expect int) {
	res := solve(ss)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", ss, expect, res)
	}
}

func TestSample1(t *testing.T) {
	ss := []string{
		"2+2=10",
		"1+2=a",
		"3+5=8",
	}
	runSample(t, ss, 2)
}
