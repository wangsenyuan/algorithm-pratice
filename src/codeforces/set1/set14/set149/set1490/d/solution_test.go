package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 5, 2, 1, 4}
	expect := []int{1, 0, 2, 3, 1}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 3, 1, 2}
	expect := []int{0, 1, 3, 2}
	runSample(t, A, expect)
}
