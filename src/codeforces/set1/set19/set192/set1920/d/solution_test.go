package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ops [][]int, queries []int, expect []int) {
	res := solve(ops, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{1, 3},
		{2, 3},
	}
	queries := []int{1, 2, 3, 4, 5, 6, 14, 15, 16, 20}
	expect := []int{1, 2, 1, 2, 3, 1, 2, 3, 1, 3}
	runSample(t, ops, queries, expect)
}

func TestSample2(t *testing.T) {
	ops := [][]int{
		{1, 3},
		{1, 8},
		{2, 15},
		{1, 6},
		{1, 9},
		{1, 1},
		{2, 6},
		{1, 1},
		{2, 12},
		{2, 10},
	}
	queries := []int{32752, 25178, 3198, 3199, 2460, 2461, 31450, 33260, 9016, 4996}
	expect := []int{9, 8, 1, 3, 1, 3, 6, 3, 8, 8}
	runSample(t, ops, queries, expect)
}

func TestSample3(t *testing.T) {
	ops := [][]int{
		{1, 6},
		{1, 11},
		{2, 392130334},
		{1, 4},
		{2, 744811750},
		{1, 10},
		{1, 5},
		{2, 209373780},
		{2, 178928984},
		{1, 3},
		{2, 658326464},
		{2, 1000000000},
	}
	queries := []int{914576963034536490, 640707385283752918, 636773368365261971, 584126563607944922, 1000000000000000000}
	expect := []int{11, 11, 11, 10, 11}
	runSample(t, ops, queries, expect)
}

func TestSample4(t *testing.T) {
	ops := [][]int{
		{1, 1},
		{1, 2},
	}
	queries := []int{1, 2}
	expect := []int{1, 2}
	runSample(t, ops, queries, expect)
}