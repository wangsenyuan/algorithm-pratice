package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{1, 0}
	expect := []int{0, 0, 1, 2, 0}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{-1, 1, -1}
	expect := []int{0, 1, 3, 2, 1, 0, 0}
	runSample(t, n, A, expect)
}
