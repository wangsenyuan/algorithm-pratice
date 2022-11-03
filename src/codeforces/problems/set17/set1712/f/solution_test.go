package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, X []int, expect []int) {
	res := solve(n, P, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	P := []int{1, 2, 2}
	X := []int{1, 2, 3, 4}
	expect := []int{1, 2, 2, 2}
	runSample(t, n, P, X, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	P := []int{1, 2, 3, 4, 2, 1}
	X := []int{2, 1, 3, 7, 5, 6, 4}
	expect := []int{3, 3, 4, 5, 5, 5, 4}
	runSample(t, n, P, X, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	P := []int{1, 2}
	X := []int{1}
	expect := []int{1}
	runSample(t, n, P, X, expect)
}
