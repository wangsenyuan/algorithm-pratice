package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, s int64, expect []int) {
	res := solve(n, k, s)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	var s int64 = 15
	expect := []int{1, 1, 1, 1}
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 6
	var s int64 = 36
	expect := []int{0, 0, 1}
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 5
	var s int64 = 7
	runSample(t, n, k, s, nil)
}
