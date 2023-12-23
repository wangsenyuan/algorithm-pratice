package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, A []int, Q [][]int, expect []int) {
	res := solve(n, P, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 5, 2, 3, 1, 1}
	P := []int{1, 2, 3, 3, 2}
	Q := [][]int{
		{4, 5},
		{6, 6},
	}
	expect := []int{33, 27}
	runSample(t, n, P, A, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 14
	A := []int{3, 2, 5, 3, 1, 4, 2, 2, 2, 5, 5, 5, 2, 4}
	P := []int{1, 2, 3, 1, 1, 4, 7, 3, 3, 1, 5, 3, 8}
	Q := [][]int{
		{4, 4},
		{4, 10},
		{13, 10},
		{3, 12},
		{13, 9},
		{3, 12},
		{9, 10},
		{11, 5},
	}
	expect := []int{47, 53, 48, 36, 42, 36, 48, 14}
	runSample(t, n, P, A, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 15
	A := []int{13046, 83844, 14823, 64255, 15301, 90234, 84972, 93547, 88028, 11665, 54415, 13159, 83950, 951, 42336}
	P := []int{1, 1, 2, 3, 1, 5, 3, 1, 6, 5, 8, 6, 11, 11}
	Q := [][]int{
		{14, 15},
		{14, 15},
		{7, 11},
		{7, 12},
		{6, 6},
		{4, 4},
		{10, 13},
		{1, 1},
		{2, 3},
		{3, 9},
		{12, 11},
		{6, 9},
		{1, 1},
		{15, 14},
		{6, 6},
	}
	expect := []int{3625293807, 3625293807, 5247791426, 2939428640, 8312372872, 11328719477, 9291649622, 170198116, 1413017728, 1475037160, 2537329077, 8113316668, 170198116, 3625293807, 8312372872}
	runSample(t, n, P, A, Q, expect)
}
