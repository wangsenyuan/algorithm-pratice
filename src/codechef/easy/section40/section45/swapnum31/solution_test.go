package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, k int, expect []int) {
	res := solve(A, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 1}
	k := 2
	expect := []int{1, 2, 3}
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 2, 4, 1}
	k := 3
	expect := []int{1, 2, 4, 3}
	runSample(t, A, k, expect)
}
