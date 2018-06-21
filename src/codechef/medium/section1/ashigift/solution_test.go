package main

import "testing"

func runSample(t *testing.T, X int64, B int, dishes [][]int64, C int, clans [][]int64, expect uint64) {
	res := solve(X, B, dishes, C, clans)

	if res != expect {
		t.Errorf("Sample %d %d %v %d %v, expect %d, but got %d", X, B, dishes, C, clans, expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := int64(10)
	B := 2
	dishes := [][]int64{
		{1, 3},
		{8, 1},
	}
	C := 0
	expect := uint64(5)
	runSample(t, X, B, dishes, C, nil, expect)
}

func TestSample2(t *testing.T) {
	X := int64(10)
	B := 2
	dishes := [][]int64{
		{1, 3},
		{8, 5},
	}
	C := 0
	expect := uint64(9)
	runSample(t, X, B, dishes, C, nil, expect)
}

func TestSample3(t *testing.T) {
	X := int64(10)
	B := 2
	dishes := [][]int64{
		{1, 3},
		{8, 5},
	}
	C := 3
	clans := [][]int64{
		{1, 2, 1},
		{4, 3, 2},
		{9, 1, 1},
	}
	expect := uint64(6)
	runSample(t, X, B, dishes, C, clans, expect)
}
