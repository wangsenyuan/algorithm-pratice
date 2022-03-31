package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, S [][]int, expect [][]int) {
	res := solve(S)

	sort.Slice(expect, func(i, j int) bool {
		return expect[i][2] < expect[j][2]
	})

	sort.Slice(res, func(i, j int) bool {
		return res[i][2] < res[j][2]
	})

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := [][]int{{1, 1}}
	expect := [][]int{{1, 1, 1}}
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := [][]int{
		{1, 3},
		{2, 3},
		{2, 2},
	}
	expect := [][]int{
		{1, 3, 1},
		{2, 2, 2},
		{2, 3, 3},
	}
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := [][]int{
		{1, 1},
		{3, 5},
		{4, 4},
		{3, 6},
		{4, 5},
		{1, 6},
	}
	expect := [][]int{
		{1, 1, 1},
		{3, 5, 3},
		{4, 4, 4},
		{3, 6, 6},
		{4, 5, 5},
		{1, 6, 2},
	}
	runSample(t, S, expect)
}
