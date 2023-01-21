package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, H []int, a int, b [][]int) {
	c, d := solve(n, k, H)

	if a != c || !reflect.DeepEqual(b, d) {
		t.Errorf("Sample expect %d %v, but got %d %v", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	H := []int{14, 12, 10}
	a := 2
	b := [][]int{
		{1, 2},
		{2, 3},
	}
	runSample(t, n, k, H, a, b)
}

func TestSample2(t *testing.T) {
	n, k := 4, 5
	H := []int{8, 19, 10, 13}
	a := 2
	b := [][]int{
		{3, 4},
	}
	runSample(t, n, k, H, a, b)
}
