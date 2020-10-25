package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, E [][]int, expect [][]int) {
	_, res := solve(n, m, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := [][]int{{1, 3}}
	runSample(t, n, m, E, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 1
	E := [][]int{
		{1, 1},
	}
	expect := [][]int{}
	runSample(t, n, m, E, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 0
	E := [][]int{

	}
	expect := [][]int{{1, 2}, {1, 2}}
	runSample(t, n, m, E, expect)
}
