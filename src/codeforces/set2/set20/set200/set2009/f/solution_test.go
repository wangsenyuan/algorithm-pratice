package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []int) {
	res := solve(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	queries := [][]int{
		{1, 9},
		{3, 5},
		{8, 8},
	}
	expect := []int{18, 8, 1}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 8, 3, 2, 4}
	queries := [][]int{
		{1, 14},
		{3, 7},
		{7, 10},
		{2, 11},
		{1, 25},
	}
	expect := []int{55, 20, 13, 41, 105}
	runSample(t, a, queries, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6}
	queries := [][]int{
		{1, 1},
	}
	expect := []int{6}
	runSample(t, a, queries, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 1, 6, 10, 4}
	queries := [][]int{
		{3, 21},
		{6, 17},
		{2, 2},
		{1, 5},
		{1, 14},
		{9, 15},
		{12, 13},
	}
	expect := []int{96, 62, 1, 24, 71, 31, 14}
	runSample(t, a, queries, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, 9, 10, 10, 1}
	queries := [][]int{
		{20, 25},
		{3, 11},
		{20, 22},
	}
	expect := []int{44, 65, 15}
	runSample(t, a, queries, expect)
}
