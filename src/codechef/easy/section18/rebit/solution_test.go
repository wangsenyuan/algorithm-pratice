package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, L string, expect []int64) {
	res := solve(L)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", L, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(#&#)", []int64{436731905, 935854081, 811073537, 811073537})
}
