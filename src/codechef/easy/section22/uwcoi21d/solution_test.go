package main

import "testing"

func runSample(t *testing.T, m int, deliveries [][]int, expect int64) {
	res := solve(m, len(deliveries), deliveries)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 4
	deliveries := [][]int{
		{2, 3, 4, 1},
		{1, 4, 5, 2},
		{3, 3, 5, 3},
	}
	var expect int64 = 16
	runSample(t, m, deliveries, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	deliveries := [][]int{
		{1, 2, 10, 2},
		{2, 1, 9, 1},
		{4, 1, 7, 1},
	}
	var expect int64 = 23
	runSample(t, m, deliveries, expect)
}
