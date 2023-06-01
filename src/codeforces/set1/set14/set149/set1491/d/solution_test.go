package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q [][]int, expect []bool) {
	res := solve(Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := [][]int{
		{1, 4},
		{3, 6},
		{1, 6},
		{6, 2},
		{5, 5},
	}
	expect := []bool{true, true, false, false, true}
	runSample(t, Q, expect)
}

func TestSample2(t *testing.T) {
	Q := [][]int{
		{3, 6},
	}
	expect := []bool{true}
	runSample(t, Q, expect)
}
