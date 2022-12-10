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
	A := []int{0, 0, 2, 3, 0}
	expect := []int{1, 4, 3, 2, 5}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0, 1, 2, 4}
	runSample(t, A, nil)
}

func TestSample3(t *testing.T) {
	A := []int{-1, -1, -1, -1}
	expect := []int{1, 2, 3, 4}
	runSample(t, A, expect)
}
