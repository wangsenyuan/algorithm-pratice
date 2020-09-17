package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, B []int, expect []int) {
	res := solve(n, B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	B := []int{0, 1, 1, -2, 1}
	expect := []int{0, 1, 2, 0, 3}
	runSample(t, n, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	B := []int{1000, 999999000, -1000000000}
	expect := []int{1000, 1000000000, 0}
	runSample(t, n, B, expect)
}
