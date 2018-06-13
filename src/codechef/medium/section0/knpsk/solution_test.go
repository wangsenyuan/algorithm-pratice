package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, W []int, C []int, expect []int64) {
	res := solve(n, W, C)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v, expect %v, but got %v", n, W, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	W := []int{1, 2, 2, 2, 2}
	C := []int{1, 2, 3, 4, 5}
	expect := []int64{1, 5, 6, 9, 10, 12, 13, 14, 15}
	runSample(t, n, W, C, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	W := []int{2, 2, 2, 2, 2}
	C := []int{1, 2, 3, 4, 5}
	expect := []int64{0, 5, 5, 9, 9, 12, 12, 14, 14, 15}
	runSample(t, n, W, C, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	W := []int{1, 1, 1, 1, 1}
	C := []int{1, 2, 3, 4, 5}
	expect := []int64{5, 9, 12, 14, 15}
	runSample(t, n, W, C, expect)
}
