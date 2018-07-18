package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(S)

	if res != int64(expect) {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcd", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "dbca", 1)
}
func TestSample3(t *testing.T) {
	runSample(t, "dcba", 0)
}
