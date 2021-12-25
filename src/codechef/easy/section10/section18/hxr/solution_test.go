package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []uint64, rgs [][]int, K int, expect []uint64) {
	res := solve(n, A, rgs, K)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %d, expect %v, bug got %v", n, A, rgs, K, expect, res)
	}
}
