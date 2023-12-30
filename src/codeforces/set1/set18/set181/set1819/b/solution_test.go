package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, rects [][]int, expect [][]int) {
	res := solve(rects)

	sort.Slice(expect, func(i, j int) bool {
		return expect[i][0] < expect[j][0] || expect[i][0] == expect[j][0] && expect[i][1] < expect[j][1]
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0] || res[i][0] == res[j][0] && res[i][1] < res[j][1]
	})

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rects := [][]int{
		{1, 2},
		{3, 5},
		{1, 3},
	}
	expect := [][]int{{4, 5}}
	runSample(t, rects, expect)
}

func TestSample2(t *testing.T) {
	rects := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	expect := [][]int{{1, 3}, {3, 1}}
	runSample(t, rects, expect)
}

func TestSample3(t *testing.T) {
	rects := [][]int{
		{10, 10},
	}
	expect := [][]int{{10, 10}}
	runSample(t, rects, expect)
}

func TestSample4(t *testing.T) {
	rects := [][]int{
		{3, 2},
		{5, 5},
		{2, 2},
		{8, 7},
	}
	expect := [][]int{{13, 7}}
	runSample(t, rects, expect)
}
