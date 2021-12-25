package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, A []int, expect []int) {
	res := solve(n, k, A)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, []int{1, 2}, []int{2, 3, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, []int{1}, nil)
}
