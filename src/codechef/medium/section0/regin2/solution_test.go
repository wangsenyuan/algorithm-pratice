package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, B []int, expect []uint64) {
	res := solve(n, A, B)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{3, 1, 1, 5}
	B := []int{4, 5, 3, 4}
	expect := []uint64{5, 12, 21, 31}
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{4, 6, 4, 4, 3}
	B := []int{9, 6, 7, 9, 8}
	expect := []uint64{6, 17, 36, 61, 91}
	runSample(t, n, A, B, expect)
}
