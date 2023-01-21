package main

import "testing"

func runSample(t *testing.T, m int, A [][]int, expect int64) {
	res := solve(m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	A := [][]int{
		{0, 0, 1},
		{4, 3, 2},
	}
	var expect int64 = 5
	runSample(t, m, A, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	A := [][]int{
		{0, 4, 8, 12, 16},
		{2, 6, 10, 14, 18},
	}
	var expect int64 = 19
	runSample(t, m, A, expect)
}

func TestSample3(t *testing.T) {
	m := 4
	A := [][]int{
		{0, 10, 10, 10},
		{10, 10, 10, 10},
	}
	var expect int64 = 17
	runSample(t, m, A, expect)
}
