package main

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, n int, A []int, X []int, expect []int) {
	res := solve(n, A, X)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := []int{3, 1, 6, 7, 2, 5, 4}
	X := []int{1, 2, 3, 4, 5, 6, 7}
	expect := []int{0, 1, 1, 2, 1, 0, 0}
	runSample(t, n, A, X, expect)
}
