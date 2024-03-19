package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, requests []int, expect [][]int) {
	res := solve(k, requests)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	requests := []int{1, 1}
	expect := [][]int{
		{1, 1, 1},
		{-1},
	}
	runSample(t, k, requests, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	requests := []int{1, 2, 3, 1}
	expect := [][]int{
		{2, 2, 2},
		{1, 1, 2},
		{3, 1, 3},
		{2, 1, 1},
	}
	runSample(t, k, requests, expect)
}
