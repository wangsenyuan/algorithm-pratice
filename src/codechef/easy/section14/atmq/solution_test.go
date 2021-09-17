package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{2, 3, 4, 1, 5}
	expect := []int{1, 4, 2, 5, 3}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 2, 4, 2}
	expect := []int{1, 3, 2, 4}
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	A := []int{5, 1, 3, 7, 6, 9, 8, 9, 2, 10}
	expect := []int{1, 10, 2, 3, 8, 4, 7, 5, 9, 6}
	runSample(t, n, A, expect)
}
