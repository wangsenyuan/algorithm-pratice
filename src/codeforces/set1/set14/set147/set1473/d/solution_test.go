package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, Q [][]int, expect []int) {
	res := solve(s, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "-+--+--+"
	Q := [][]int{
		{1, 8},
		{2, 8},
		{2, 5},
		{1, 1},
	}
	expect := []int{1, 2, 4, 4}
	runSample(t, s, Q, expect)
}
