package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, S, K, m, M int, expect []int) {
	res := solve(N, S, K, m, M)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d %d, expect %v, but got %v", N, S, K, m, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 6, 1, 1, 5, []int{1, 1, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 4, 2, 1, 3, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 10, 3, 1, 5, []int{1, 1, 1, 1, 2, 4})
}
