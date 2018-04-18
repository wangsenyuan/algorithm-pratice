package main

import "testing"

func runSample(t *testing.T, s string, n int, expect int64) {
	res := solve([]byte(s), n)

	if res != expect {
		t.Errorf("sample %s %d, expect %d, but got %d", s, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aba", 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, "baa", 3, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, "bba", 3, 0)
}
