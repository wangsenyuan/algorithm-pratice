package main

import "testing"

func runSample(t *testing.T, n, k int, S string, expect int) {
	res := solve(n, k, []byte(S))

	if res != expect {
		t.Errorf("Sample %d %d %s, expect %d, but got %d", n, k, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	S := "01001"
	expect := 1
	runSample(t, n, k, S, expect)
}
