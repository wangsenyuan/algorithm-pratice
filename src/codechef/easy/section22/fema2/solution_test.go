package main

import "testing"

func runSample(t *testing.T, n int, K int, s string, expect int) {
	res := solve(n, K, s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 5
	s := "I::M"
	expect := 1
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 9, 10
	s := "MIM_XII:M"
	expect := 2
	runSample(t, n, k, s, expect)
}
