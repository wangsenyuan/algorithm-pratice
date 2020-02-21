package main

import "testing"

func runSample(t *testing.T, first, second []int, expect int64) {
	res := solve(first, second)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", first, second, expect, res)
	}
}

func TestSample1(t *testing.T) {
	first := []int{2, 2020}
	second := []int{2, 2020}
	runSample(t, first, second, 1)
}

func TestSample2(t *testing.T) {
	first := []int{10, 2009}
	second := []int{2, 2020}
	runSample(t, first, second, 3)
}
