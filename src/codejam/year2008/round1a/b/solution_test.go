package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, flavors [][]int, expect []int) {
	res := solve(n, m, flavors)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, m, flavors, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 3
	flavors := [][]int{
		{1, 1},
		{1, 0, 2, 0},
		{5, 0},
	}
	expect := []int{1, 0, 0, 0, 0}
	runSample(t, n, m, flavors, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 2
	flavors := [][]int{
		{1, 0},
		{1, 1},
	}
	runSample(t, n, m, flavors, nil)
}
