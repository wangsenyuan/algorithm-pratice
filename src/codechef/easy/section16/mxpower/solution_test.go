package main

import "testing"

func runSample(t *testing.T, N int, E [][]int64, expect int64) {
	res := solve(N, E)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 1
	E := [][]int64{{-5}}
	expect := int64(-5)
	runSample(t, N, E, expect)
}

func TestSample2(t *testing.T) {
	N := 2
	E := [][]int64{{-1, -1}, {1, 1}}
	expect := int64(1)
	runSample(t, N, E, expect)
}

func TestSample3(t *testing.T) {
	N := 3
	E := [][]int64{
		{-10, 5, -50},
		{26, -1, 10},
		{0, 6, 45},
	}
	expect := int64(46)
	runSample(t, N, E, expect)
}
