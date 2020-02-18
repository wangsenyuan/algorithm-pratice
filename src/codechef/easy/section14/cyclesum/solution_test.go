package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int64, expect []int64) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int64{-5, 4, 1, 2}
	expect := []int64{7, 7, 4, 5}
	runSample(t, n, A, expect)
}
