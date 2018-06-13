package main

import "testing"

func runSample(t *testing.T, n int, K int, W []int, C []int, expect int64) {
	res := solve(n, K, W, C)
	if res != expect {
		t.Errorf("sample %d %d %v %v, expect %d, but got %d", n, K, W, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 1, 3
	C := []int{2}
	W := []int{2}
	var expect int64 = 2
	runSample(t, n, k, W, C, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 4
	C := []int{2, 2, 3}
	W := []int{1, 2, 5}
	var expect int64 = 5
	runSample(t, n, k, W, C, expect)
}
