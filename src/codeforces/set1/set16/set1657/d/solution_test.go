package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, C int, units [][]int64, monsters [][]int64, expect []int) {
	res := solve(C, units, monsters)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := 10
	units := [][]int64{
		{3, 4, 6},
		{5, 5, 5},
		{10, 3, 4},
	}
	monsters := [][]int64{
		{8, 3},
		{5, 4},
		{10, 15},
	}
	expect := []int{5, 3, -1}
	runSample(t, C, units, monsters, expect)
}

func TestSample2(t *testing.T) {
	C := 15
	units := [][]int64{
		{14, 10, 3},
		{9, 2, 2},
		{10, 4, 3},
		{7, 3, 5},
		{4, 3, 1},
	}
	monsters := [][]int64{
		{11, 2},
		{1, 1},
		{4, 7},
		{2, 1},
		{1, 14},
		{3, 3},
	}
	expect := []int{14, 4, 14, 4, 7, 7}
	runSample(t, C, units, monsters, expect)
}
