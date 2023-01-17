package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int64, k int64, qs []int64, expect string) {
	res := solve(n, k, qs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n, k int64 = 3, 1
	qs := []int64{1, 2, 3}
	runSample(t, n, k, qs, "..X")
}

func TestSample2(t *testing.T) {
	var n, k int64 = 6, 3
	qs := []int64{1, 2, 3, 4, 5, 6}
	runSample(t, n, k, qs, ".X.X.X")
}

func TestSample3(t *testing.T) {
	var n, k int64 = 5, 2
	qs := []int64{1, 2, 3, 4, 5}
	runSample(t, n, k, qs, "...XX")
}
