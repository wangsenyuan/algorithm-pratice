package main

import "testing"

func runSample(t *testing.T, S string, expect int64) {
	res := solve([]byte(S))
	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ABACABA", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "ABCAB", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "ABCABC", 5)
}
