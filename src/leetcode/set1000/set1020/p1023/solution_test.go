package p1023

import "testing"

func runSample(t *testing.T, S string, N int, expect bool) {
	res := queryString(S, N)
	if res != expect {
		t.Errorf("Sample %s %d, expect %t, but got %t", S, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "0110", 3, true)
}
