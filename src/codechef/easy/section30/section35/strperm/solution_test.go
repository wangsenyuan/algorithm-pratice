package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, conditions [][]int, expect []int) {
	res := solve(n, conditions)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	conditions := [][]int{
		{2, 1},
		{1, 2},
	}
	expect := []int{2, 1, 3}
	runSample(t, n, conditions, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	conditions := [][]int{
		{1, 2},
		{4, 1},
		{3, 2},
	}

	runSample(t, n, conditions, nil)
}

func TestSample3(t *testing.T) {
	n := 4
	conditions := [][]int{
		{2, 3},
	}

	expect := []int{1, 2, 3, 4}

	runSample(t, n, conditions, expect)
}
