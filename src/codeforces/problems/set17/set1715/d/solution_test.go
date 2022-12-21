package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{4, 1, 2},
	}
	expect := []int{0, 3, 2, 2}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 1, 1073741823},
	}
	expect := []int{1073741823, 0}
	runSample(t, n, E, expect)
}
