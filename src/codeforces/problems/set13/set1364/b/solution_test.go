package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, expect []int) {
	res := solve(n, P)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := []int{3, 2, 1}
	expect := []int{3, 1}
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	P := []int{1, 3, 4, 2}
	expect := []int{1, 4, 2}
	runSample(t, n, P, expect)
}
