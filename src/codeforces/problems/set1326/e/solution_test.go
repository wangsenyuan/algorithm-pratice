package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, Q []int, expect []int) {
	res := solve(n, P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, P, Q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := []int{3, 2, 1}
	Q := []int{1, 2, 3}

	expect := []int{3, 2, 1}

	runSample(t, n, P, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	P := []int{2, 3, 6, 1, 5, 4}
	Q := []int{5, 2, 1, 4, 6, 3}

	expect := []int{6, 5, 5, 5, 4, 1}

	runSample(t, n, P, Q, expect)
}
