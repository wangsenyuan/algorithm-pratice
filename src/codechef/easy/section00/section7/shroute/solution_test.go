package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, A, B []int, expect []int) {
	res := solve(n, m, A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 1
	A := []int{1, 0, 0, 0, 0}
	B := []int{5}
	expect := []int{4}
	runSample(t, n, m, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 1
	A := []int{1, 0, 0, 0, 2}
	B := []int{4}
	expect := []int{1}
	runSample(t, n, m, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 2
	A := []int{2, 0, 0, 0, 1}
	B := []int{3, 1}
	expect := []int{-1, 0}
	runSample(t, n, m, A, B, expect)
}
