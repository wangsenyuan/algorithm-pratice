package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 5, 8, 13, 21, 34}
	expect := []int{1, 4, 10, 58, 578, 20098, 5236738, 24641495}

	for i, num := range nums {
		runSample(t, num, expect[i])
	}
}
