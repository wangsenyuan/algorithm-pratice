package main

import "testing"

func runSample(t *testing.T, s string, k int, m int, expect int) {
	res := solve(len(s), s, k, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, m := 9, 2
	s := "5418"
	expect := 14
	runSample(t, s, k, m, expect)
}

func TestSample2(t *testing.T) {
	k, m := 63, 3
	s := "40514"
	expect := 88
	runSample(t, s, k, m, expect)
}
