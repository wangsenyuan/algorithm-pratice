package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, c []int, p []int, deps [][]int, expect []int) {
	res := solve(c, p, deps)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{30, 8, 3, 5, 10}
	p := []int{3}
	deps := [][]int{
		{2, 4, 5},
		{},
		{},
		{3, 5},
		{},
	}
	expect := []int{23, 8, 0, 5, 10}
	runSample(t, c, p, deps, expect)
}

func TestSample2(t *testing.T) {
	c := []int{5, 143, 3}
	p := []int{1, 3}
	deps := [][]int{
		{2},
		{},
		{1, 2},
	}
	expect := []int{0, 143, 0}
	runSample(t, c, p, deps, expect)
}

func TestSample3(t *testing.T) {
	c := []int{5, 4, 1, 3, 4}
	p := []int{2}
	deps := [][]int{
		{4, 5},
		{3, 5, 4},
		{1, 4},
		{5},
		{},
	}
	expect := []int{5, 0, 1, 3, 4}
	runSample(t, c, p, deps, expect)
}
