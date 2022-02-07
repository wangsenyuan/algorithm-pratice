package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, a int, b []int) {
	c, d := solve(n, E)

	if a != c || !reflect.DeepEqual(b, d) {
		t.Errorf("Sample expect %d %v, but got %d %v", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
		{5, 8},
		{6, 9},
		{6, 10},
	}
	a := 3
	b := []int{8, 9, 10}
	runSample(t, n, E, a, b)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{4, 6},
		{5, 7},
	}
	a := 3
	b := []int{7}
	runSample(t, n, E, a, b)
}
