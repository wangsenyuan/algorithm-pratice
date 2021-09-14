package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, X []int, expect []int) {
	res := solve(n, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	X := []int{1, 2, 3, 2, 4}
	expect := []int{1, 2, 3, 5, 4}
	runSample(t, n, X, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	X := []int{4, 1, 3, 2}
	expect := []int{4, 1, 3, 2}
	runSample(t, n, X, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	X := []int{1, 1, 1}
	expect := []int{1, 2, 3}
	runSample(t, n, X, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	X := []int{2, 5, 1, 5, 2}
	expect := []int{2, 5, 1, 6, 3}
	runSample(t, n, X, expect)
}
