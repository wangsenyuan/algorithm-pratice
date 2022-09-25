package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, M int, X int, expect []int) {
	res := solve(A, M, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5000, 5001, 5002}
	M := 8000
	X := 2
	expect := []int{2, 3}
	runSample(t, A, M, X, expect)
}

func TestSample2(t *testing.T) {
	A := []int{7999, 7998, 8000}
	M := 5000
	X := 2
	expect := []int{1, 2, 3}
	runSample(t, A, M, X, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5999, 5998, 6000, 6001}
	M := 6000
	X := 3
	expect := []int{1, 3, 4}
	runSample(t, A, M, X, expect)
}
