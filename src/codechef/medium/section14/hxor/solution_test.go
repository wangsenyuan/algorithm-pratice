package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, A []int, expect []int) {

	solve(n, k, A)

	if !reflect.DeepEqual(A, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, A)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	A := []int{2, 2, 3}
	expect := []int{0, 0, 3}
	runSample(t, n, k, A, expect)
}
