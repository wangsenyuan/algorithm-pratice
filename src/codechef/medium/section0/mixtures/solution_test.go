package main

import "testing"

func runSample(t *testing.T, n int, mixtures []int, expect int64) {
	res := solve(n, mixtures)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, mixtures, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	mixtures := []int{18, 19}
	runSample(t, n, mixtures, 342)
}

func TestSample2(t *testing.T) {
	n := 3
	mixtures := []int{40, 60, 20}
	runSample(t, n, mixtures, 2400)
}
