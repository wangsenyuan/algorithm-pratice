package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, V []int, X []int, expect []int) {
	res := solve(V, X)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	V := []int{10, 10, 5}
	X := []int{5, 7, 2}
	expect := []int{5, 12, 4}
	runSample(t, V, X, expect)
}

func TestSample2(t *testing.T) {
	V := []int{30, 25, 20, 15, 10}
	X := []int{9, 10, 12, 4, 13}
	expect := []int{9, 20, 35, 11, 25}
	runSample(t, V, X, expect)
}
