package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rows [][]int, Q []int, expect []int) {
	res := solve(rows, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rows := [][]int{
		{1, 5},
		{10, 11},
		{8, 8},
	}
	Q := []int{12, 5}
	expect := []int{2, 3}
	runSample(t, rows, Q, expect)
}
