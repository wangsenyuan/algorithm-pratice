package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []string) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2, 101},
		{1, 3, 100},
		{2, 3, 2},
		{2, 4, 2},
		{3, 4, 1},
	}
	expect := []string{
		"none",
		"any",
		"at least one",
		"at least one",
		"any",
	}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{1, 3, 2},
	}
	expect := []string{
		"any",
		"any",
		"none",
	}
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{1, 3, 1},
	}
	expect := []string{
		"at least one",
		"at least one",
		"at least one",
	}
	runSample(t, n, E, expect)
}
