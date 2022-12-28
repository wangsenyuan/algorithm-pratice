package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 8, 1
	expect := []int{1, 1, 2, 2, 3, 4, 5, 6}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 2
	expect := []int{0, 1, 0, 1, 1, 1, 1, 2, 2, 2}
	runSample(t, n, k, expect)
}
