package main

import "testing"

func runSample(t *testing.T, n int, k int64, expect string) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %s, but got %s", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4, "0101")
}
