package main

import "testing"

func runSample(t *testing.T, n int, F int, H []int, expect int64) {
	res := solve(n, F, H)
	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, F, H, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, F := 3, 2
	H := []int{2, 2, 2}
	runSample(t, n, F, H, 3)
}

func TestSample2(t *testing.T) {
	n, F := 3, 2
	H := []int{2, 1, 4}
	runSample(t, n, F, H, 3)
}

func TestSample3(t *testing.T) {
	n, F := 5, 2
	H := []int{1, 2, 2, 4, 4}
	runSample(t, n, F, H, 10)
}
