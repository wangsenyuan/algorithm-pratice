package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, A []int, expect []int) {
	solve(n, k, A)

	if !reflect.DeepEqual(A, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, A)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	A := []int{3, 1, 4}
	expect := []int{2, 1, 5}
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 10
	A := []int{1, 0}
	expect := []int{0, 1}
	runSample(t, n, k, A, expect)
}
