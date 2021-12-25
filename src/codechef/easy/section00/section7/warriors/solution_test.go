package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q int, P []int, X []int, expect []int) {
	res := solve(N, Q, P, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got%v", N, Q, P, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 3, 4
	P := []int{1, 2, 1}
	X := []int{10, 2, 3, 1}
	expect := []int{3, 2, 3, 0}
	runSample(t, N, Q, P, X, expect)
}
