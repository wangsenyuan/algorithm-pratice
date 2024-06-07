package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, h []int, queries [][]int, expect []int) {
	res := solve(h, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := []int{1, 2, 3, 4}
	queries := [][]int{
		{1, 1},
		{1, 4},
		{4, 3},
		{4, 5},
	}
	expect := []int{4, 3, 3, 4}
	runSample(t, h, queries, expect)
}

func TestSample2(t *testing.T) {
	h := []int{1, 3, 2, 6}
	queries := [][]int{
		{3, 5},
		{2, 4},
	}
	expect := []int{4, 3}
	runSample(t, h, queries, expect)
}

func TestSample3(t *testing.T) {
	h := []int{1, 4, 3, 2, 7, 6, 5}
	queries := [][]int{
		{4, 4},
	}
	expect := []int{4}
	runSample(t, h, queries, expect)
}

func TestSample4(t *testing.T) {
	h := []int{22, 90, 86, 54, 23, 79, 1, 92, 52, 17, 34, 26, 45}
	queries := [][]int{
		{11, 47},
		{11, 95},
		{12, 32},
		{7, 60},
		{11, 56},
		{5, 65},
		{9, 41},
		{12, 30},
		{1, 15},
		{5, 65},
	}
	expect := []int{4, 5, 4, 4, 4, 5, 4, 4, 4, 5}
	runSample(t, h, queries, expect)
}

func TestSample5(t *testing.T) {
	h := []int{22, 90, 86, 54, 23, 79, 1, 92, 52, 17, 34, 26, 45}
	queries := [][]int{
		{12, 30},
	}
	expect := []int{4}
	runSample(t, h, queries, expect)
}
