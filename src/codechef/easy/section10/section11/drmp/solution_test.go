package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, M int64, expect []int64) {
	res := solve(M)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d, expect %v, but got %v", M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, []int64{4, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, 6, []int64{7, 8, 9, 10, 12})
}
