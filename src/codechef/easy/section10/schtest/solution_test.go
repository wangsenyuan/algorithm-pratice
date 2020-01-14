package main

import "testing"

func runSample(t *testing.T, N int64, R int64, X []int64, Y []int64, expect int64) {
	res := solve(N, R, X, Y)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", N, R, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, R := int64(5), int64(3)
	Y := []int64{4}
	expect := int64(3)
	runSample(t, N, R, nil, Y, expect)
}

func TestSample2(t *testing.T) {
	N, R := int64(10), int64(2)
	X := []int64{3, 1, 7, 6}
	Y := []int64{4, 3, 1, 5, 9, 7}
	expect := int64(2)
	runSample(t, N, R, X, Y, expect)
}

func TestSample3(t *testing.T) {
	N, R := int64(10), int64(4)
	X := []int64{3, 1, 7, 6}
	Y := []int64{4, 3, 1, 5, 9, 7}
	expect := int64(3)
	runSample(t, N, R, X, Y, expect)
}
