package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m, q int, A []int, ops [][]int, expect []int) {
	res := solve(n, m, q, A, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v, expect %v, but got %v", n, m, q, A, ops, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, q := 5, 2, 1
	A := []int{3, 2, 4, 1, 5}
	ops := [][]int{{2, 4}}
	expect := []int{340}
	runSample(t, n, m, q, A, ops, expect)
}
