package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, q int, X []int, Y []int) {
	res := solve(n, A, q, X)

	if !reflect.DeepEqual(Y, res) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, A, q, X, Y, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{2, 2, 3}
	q := 3
	X := []int{1, 3, 4}
	Y := []int{2, 3, 2}
	runSample(t, n, A, q, X, Y)
}
