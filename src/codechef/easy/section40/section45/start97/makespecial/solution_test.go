package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int64, L []int64, R []int64, expect []int64) {
	res := solve(n, L, R)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n int64 = 3
	L := []int64{1, 3, 4}
	R := []int64{3, 6, 14}
	expect := []int64{1, 2, 5}
	runSample(t, n, L, R, expect)
}

func TestSample2(t *testing.T) {
	var n int64 = 6
	L := []int64{1, 3, 4}
	R := []int64{3, 6, 14}
	expect := []int64{0, 0, 2}
	runSample(t, n, L, R, expect)
}
