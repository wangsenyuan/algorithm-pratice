package main

import "testing"

func runSample(t *testing.T, P, K, L int, F []int, expect int64) {
	res := solve(P, K, L, F)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", P, K, L, F, expect, res)
	}
}

func TestSample1(t *testing.T) {
	P, K, L := 3, 2, 6
	F := []int{8, 2, 5, 2, 4, 9}
	expect := int64(47)
	runSample(t, P, K, L, F, expect)
}

func TestSample2(t *testing.T) {
	P, K, L := 3, 9, 26
	F := []int{1, 1, 1, 100, 100, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 11, 11, 11, 11, 1, 1, 1, 100}
	expect := int64(397)
	runSample(t, P, K, L, F, expect)
}
