package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, K, Q int, F []int, X []int, expect []bool) {
	res := solve(N, K, Q, F, X)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v, expect %v, but got %v", N, K, Q, F, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, Q := 4, 2, 3
	F := []int{2, 4}
	X := []int{2, 8, 5}
	expect := []bool{true, true, false}
	runSample(t, N, K, Q, F, X, expect)
}
