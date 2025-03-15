package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 1, 5
	a := []int{5}
	runSample(t, n, k, a)
}

func TestSample2(t *testing.T) {
	n, k := 3, 16
	a := []int{16, 0, 16}
	runSample(t, n, k, a)
}

func TestSample3(t *testing.T) {
	n, k := 9, 1
	a := []int{1, 0, 0, 0, 0, 0, 0, 0, 1}
	runSample(t, n, k, a)
}
