package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, W int, H int, cuts [][]int, expect []int) {
	res := solve(W, H, cuts)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W, H := 2, 2
	cuts := [][]int{
		{1, 0, 1, 2},
		{0, 1, 1, 1},
	}
	expect := []int{1, 1, 2}
	runSample(t, W, H, cuts, expect)
}

func TestSample2(t *testing.T) {
	W, H := 2, 2
	cuts := [][]int{
		{1, 0, 1, 2},
		{0, 1, 1, 1},
		{1, 1, 2, 1},
	}
	expect := []int{1, 1, 1, 1}
	runSample(t, W, H, cuts, expect)
}

func TestSample3(t *testing.T) {
	W, H := 2, 4
	cuts := [][]int{
		{0, 1, 2, 1},
		{0, 3, 2, 3},
	}
	expect := []int{2, 2, 4}
	runSample(t, W, H, cuts, expect)
}
