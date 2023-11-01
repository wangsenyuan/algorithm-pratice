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
	a := []int{1, 2}
	queries := [][]int{
		{1, 3},
	}
	expect := []int{3, 5}
	runSample(t, a, queries, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 1, 3, 2}
	queries := [][]int{
		{2, 6},
		{4, 6},
		{1, 1},
		{3, 11},
	}
	expect := []int{14, 16, 24, 24, 24}
	runSample(t, a, queries, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 7, 9, 4, 5, 2, 3, 6, 8}
	queries := [][]int{
		{3, 1},
		{2, 1},
		{9, 1},
	}
	expect := []int{57, 54, 36, 36}
	runSample(t, a, queries, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1}
	queries := [][]int{
		{1, 5},
		{4, 4},
		{3, 8},
	}
	expect := []int{6, 18, 27, 28}
	runSample(t, a, queries, expect)
}

func TestSample5(t *testing.T) {
	a := []int{7, 9, 4, 3, 12, 6, 5, 1, 8, 10, 2, 11}
	queries := [][]int{
		{3, 267878},
		{10, 107293},
	}
	expect := []int{108, 1607316, 1607316}
	runSample(t, a, queries, expect)
}
