package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, kids [][]int, expect []int) {
	res := solve(kids)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	kids := [][]int{
		{3, 4},
		{5, 5},
		{4, 3},
	}
	expect := []int{1, 2}
	runSample(t, kids, expect)
}

func TestSample2(t *testing.T) {
	kids := [][]int{
		{7, 7},
		{4, 7},
		{5, 5},
		{1, 2},
		{3, 4},
		{2, 1},
	}
	expect := []int{1, 2, 1, 2}
	runSample(t, kids, expect)
}
