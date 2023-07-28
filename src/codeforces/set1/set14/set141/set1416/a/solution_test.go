package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	expect := []int{-1, -1, 3, 2, 1}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 4, 4, 4, 2}
	expect := []int{-1, 4, 4, 4, 2}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 1, 5, 3, 1}
	expect := []int{-1, -1, 1, 1, 1, 1}
	runSample(t, A, expect)
}
