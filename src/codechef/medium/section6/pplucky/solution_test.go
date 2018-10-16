package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int64) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, "4747", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, "4447477747", 20)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, "744744747", 17)
}
