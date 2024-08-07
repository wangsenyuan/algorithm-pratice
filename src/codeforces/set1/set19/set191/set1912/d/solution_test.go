package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, b int, n int, expect []int) {
	res := solve(b, n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 3, []int{2, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 11, []int{3, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 4, []int{1, 2})
}
