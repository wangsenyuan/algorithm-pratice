package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, ops [][]int, expect []bool) {
	res := solve(n, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	ops := [][]int{
		{1, 1},
		{1, 2},
		{2},
		{1, 8},
		{1, 9},
	}
	expect := []bool{true, true, true, true}
	runSample(t, n, ops, expect)
}

func TestSample2(t *testing.T) {
	n := 20
	ops := [][]int{
		{1, 3},
		{2},
		{1, 7},
		{1, 12},
	}
	expect := []bool{true, false, true}
	runSample(t, n, ops, expect)
}

func TestSample3(t *testing.T) {
	n := 16
	ops := [][]int{
		{1, 6}, {1, 6}, {1, 10}, {1, 9}, {1, 1}, {1, 9}, {1, 7}, {1, 3}, {1, 2}, {1, 10},
	}
	expect := []bool{true,
		false,
		true,
		true,
		true,
		false,
		true,
		false,
		true,
		true}
	runSample(t, n, ops, expect)
}
