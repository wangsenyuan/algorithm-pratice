package main

import "testing"

func runSample(t *testing.T, ss []string, expect int64) {
	res := solve(len(ss), ss)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", ss, expect, res)
	}
}

func TestSample1(t *testing.T) {
	ss := []string{"a", "ab", "abb"}
	runSample(t, ss, 2)
}
