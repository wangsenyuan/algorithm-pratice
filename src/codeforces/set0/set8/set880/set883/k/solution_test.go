package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, segs [][]int, expect int, arr []int) {
	res, hi := solve(segs)

	if res != expect || !reflect.DeepEqual(arr, hi) {
		t.Fatalf("Sample expect %d, %v, but got %d %v", expect, arr, res, hi)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{4, 5},
		{4, 5},
		{4, 10},
	}
	res := 16
	arr := []int{9, 9, 10}
	runSample(t, segs, res, arr)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{1, 100},
		{100, 1},
		{1, 100},
		{100, 1},
	}
	res := 202
	arr := []int{101, 101, 101, 101}
	runSample(t, segs, res, arr)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 1},
		{100, 100},
		{1, 1},
	}
	res := -1
	runSample(t, segs, res, nil)
}
