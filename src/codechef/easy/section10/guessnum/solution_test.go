package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, M int64, expect []int64) {
	res := solve(A, M)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", A, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 35, []int64{10})
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 50, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 65, []int64{13, 15, 16})
}
