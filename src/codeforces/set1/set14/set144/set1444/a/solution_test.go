package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int64, Q []int64, expect []int64) {
	res := solve(P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int64{10, 12, 179}
	Q := []int64{4, 6, 822}
	expect := []int64{10, 4, 179}
	runSample(t, P, Q, expect)
}
