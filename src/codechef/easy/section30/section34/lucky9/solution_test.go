package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 只有74符合答案,
	s := "47"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	// 只有74符合答案,
	s := "4477"
	expect := 4
	runSample(t, s, expect)
}
