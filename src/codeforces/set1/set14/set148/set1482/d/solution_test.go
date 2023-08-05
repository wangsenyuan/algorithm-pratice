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
	A := []int{5, 9, 2, 10, 15}
	expect := []int{2, 3}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 4, 2, 4, 2}
	expect := []int{2, 1}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	expect := []int{2, 1}
	runSample(t, A, expect)
}
