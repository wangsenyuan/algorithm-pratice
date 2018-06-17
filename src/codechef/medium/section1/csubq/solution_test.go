package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q, L, R int, queries [][]int, expect []int) {
	res := solve(N, Q, L, R, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d %v, expect %v, but got %v", N, Q, L, R, queries, expect, res)
	}
}
