package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, roads [][]int, expect []int) {
	res := solve(a, roads)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 4, 1, 3}
	roads := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 4},
		{4, 5},
		{5, 2},
	}
	expect := []int{5, 12}
	runSample(t, a, roads, expect)
}

func TestSample2(t *testing.T) {
	a := []int{999999999, 999999999, 999999999, 999999999, 1000000000, 999999999, 1000000000}
	roads := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{4, 5},
		{4, 6},
		{6, 7},
	}
	expect := []int{6, 5999999995}
	runSample(t, a, roads, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3, 5, 7, 3, 4, 1, 4, 3, 4, 2, 2, 5, 1}
	roads := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 1},
		{4, 4},
		{4, 5},
		{5, 6},
		{5, 6},
		{5, 12},
		{6, 7},
		{6, 8},
		{7, 5},
		{7, 7},
		{7, 9},
		{8, 4},
		{9, 11},
		{10, 9},
		{11, 10},
		{11, 10},
		{12, 13},
		{13, 14},
		{14, 12},
	}
	expect := []int{11, 37}
	runSample(t, a, roads, expect)
}

func TestSample4(t *testing.T) {
	a := []int{56222230, 648830822, 625661495, 405731832, 907493368, 709324506, 581264082, 916197089, 879420933, 206411164, 267473421, 953927943, 528302118, 271708706, 795094046, 754816745, 90376225, 436181609, 715057080, 99423178}
	roads := [][]int{
		{2, 7},
		{17, 12},
		{13, 2},
		{13, 4},
		{11, 6},
		{6, 5},
		{17, 11},
		{20, 17},
		{20, 9},
		{18, 14},
		{8, 10},
		{5, 8},
		{15, 13},
		{17, 15},
		{8, 14},
		{10, 19},
		{18, 2},
		{6, 13},
		{8, 20},
		{13, 1},
		{20, 20},
	}
	expect := []int{10, 5543778855}
	runSample(t, a, roads, expect)
}
