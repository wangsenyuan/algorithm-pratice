package main

import "testing"

func runSample(t *testing.T, ss []string, expect int) {
	res := solve(ss)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ss := []string{
		"wolverine",
		"wolverinexmen",
		"wolv",
		"perl",
		"pe",
	}
	expect := 5
	runSample(t, ss, expect)
}
