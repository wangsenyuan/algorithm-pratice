package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, Q int, F []int, K []int, expect []int) {
	res := solve(N, Q, F, K)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, F, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 3, 4
	F := []int{0, 1, 2}
	K := []int{7, 2, 5, 1000000000}
	expect := []int{3, 1, 0, 0}
	runSample(t, N, Q, F, K, expect)
}
