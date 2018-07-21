package main

import "testing"

func runSample(t *testing.T, S string, expect int64) {
	res := solve([]byte(S))
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "11010", 9)
}

func TestSample2(t *testing.T) {
	runSample(t, "101001011", 18)
}
