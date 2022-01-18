package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 6, 1
	expect := []int{1, 3, 5, 2, 4, 6}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 2
	expect := []int{1, 5, 2, 6, 3, 7, 4, 8}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 14, 452
	expect := []int{1, 10, 6, 2, 11, 7, 3, 12, 8, 4, 13, 9, 5, 14}
	runSample(t, n, k, expect)
}
