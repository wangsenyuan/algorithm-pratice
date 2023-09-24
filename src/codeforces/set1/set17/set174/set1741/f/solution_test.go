package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, segments [][]int, expect []int) {
	res := solve(segments)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	segments := [][]int{
		{1, 2, 1},
		{3, 4, 1},
		{5, 6, 2},
	}
	expect := []int{3, 1, 1}
	runSample(t, segments, expect)
}

func TestSample2(t *testing.T) {
	segments := [][]int{
		{100000000, 200000000, 1},
		{900000000, 1000000000, 2},
	}
	expect := []int{700000000, 700000000}
	runSample(t, segments, expect)
}

func TestSample3(t *testing.T) {
	segments := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
		{4, 5, 4},
		{5, 6, 5},
	}
	expect := []int{0, 0, 0, 0, 0}
	runSample(t, segments, expect)
}

func TestSample4(t *testing.T) {
	segments := [][]int{
		{1, 5, 1},
		{4, 9, 2},
		{1, 2, 1},
		{8, 9, 2},
		{5, 7, 3},
	}
	expect := []int{0, 0, 2, 1, 0}
	runSample(t, segments, expect)
}

func TestSample5(t *testing.T) {
	segments := [][]int{
		{1, 100, 2},
		{10, 90, 1},
	}
	expect := []int{0, 0}
	runSample(t, segments, expect)
}

func TestSample6(t *testing.T) {
	segments := [][]int{
		{1, 1, 1},
		{10, 10, 2},
		{1000000000, 1000000000, 3},
	}
	expect := []int{9, 9, 999999990}
	runSample(t, segments, expect)
}

func TestSample7(t *testing.T) {
	segments := [][]int{
		{3, 4, 1},
		{2, 5, 1},
		{1, 6, 2},
	}
	expect := []int{0, 0, 0}
	runSample(t, segments, expect)
}

func TestSample8(t *testing.T) {
	segments := [][]int{
		{5, 6, 2},
		{11, 12, 3},
		{7, 8, 2},
		{3, 4, 2},
		{1, 2, 1},
		{9, 10, 2},
	}
	expect := []int{3, 1, 3, 1, 1, 1}
	runSample(t, segments, expect)
}
