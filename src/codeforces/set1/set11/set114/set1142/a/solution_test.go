package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, a int, b int, expect []int) {
	res := solve(n, k, a, b)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 3
	a, b := 1, 1
	expect := []int{1, 6}
	runSample(t, n, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	a, b := 0, 0
	expect := []int{1, 3}
	runSample(t, n, k, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, k := 1, 10
	a, b := 5, 3
	expect := []int{5, 5}
	runSample(t, n, k, a, b, expect)
}
