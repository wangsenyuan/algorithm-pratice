package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 11
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
		{10, 11},
	}
	expect := []int{1, 12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789, 345678826}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 12
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{2, 7},
		{2, 8},
		{2, 9},
		{2, 10},
		{3, 11},
		{11, 12},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{1, 8},
		{1, 9},
		{1, 10},
	}
	expect := []int{1, 12, 13, 14, 15, 16, 17, 18, 19, 1210, 121011}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 12
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
		{10, 11},
		{11, 12},
		{1, 3},
		{1, 4},
		{1, 10},
	}
	expect := []int{1, 12, 13, 134, 1345, 13456, 1498, 149, 14, 1410, 141011}
	runSample(t, n, edges, expect)
}
