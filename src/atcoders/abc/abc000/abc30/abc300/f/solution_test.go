package main

import "testing"

func runSample(t *testing.T, s string, m int, k int, expect int64) {
	res := solve(m, k, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "oxoxooxoxoxooxoxooxxxoxxxooxox"
	m, k := 1000000000, 9982443530
	var expect int64 = 19964887064
	runSample(t, s, m, k, expect)
}
