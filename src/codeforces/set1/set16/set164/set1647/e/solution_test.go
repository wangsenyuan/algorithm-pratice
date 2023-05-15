package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, A []int, expect []int) {
	res := solve(P, A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{5, 5, 3, 3, 1}
	A := []int{1, 8, 2, 9, 4}
	expect := []int{1, 3, 2, 5, 4}
	runSample(t, P, A, expect)
}

func TestSample2(t *testing.T) {
	P := []int{1, 3, 2, 5, 2}
	A := []int{3, 2, 5, 4, 1}
	expect := []int{3, 2, 5, 4, 1}
	runSample(t, P, A, expect)
}
