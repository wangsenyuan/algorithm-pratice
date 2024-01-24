package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, segs [][]int, queries []int, expect []int) {
	res := solve(segs, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{6, 17, 7, 14},
		{1, 12, 3, 8},
		{16, 24, 20, 22},
	}
	queries := []int{10, 2, 23, 15, 28, 18}
	expect := []int{14, 14, 23, 15, 28, 22}
	runSample(t, segs, queries, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{3, 14, 7, 10},
		{16, 24, 20, 22},
		{1, 16, 3, 14},
	}
	queries := []int{2, 4, 6, 8, 18, 23, 11, 13, 15}
	expect := []int{14, 14, 14, 14, 22, 23, 14, 14, 15}
	runSample(t, segs, queries, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 4, 2, 3},
		{3, 9, 6, 7},
	}
	queries := []int{4, 8, 1}
	expect := []int{7, 8, 7}
	runSample(t, segs, queries, expect)
}

func TestSample4(t *testing.T) {
	segs := [][]int{
		{18, 24, 18, 24},
		{1, 8, 2, 4},
		{11, 16, 14, 14},
		{26, 32, 28, 30},
		{5, 10, 6, 8},
	}
	queries := []int{15, 14, 13, 27, 22, 17, 31, 1, 7}
	expect := []int{15, 14, 14, 30, 24, 17, 31, 4, 8}
	runSample(t, segs, queries, expect)
}

func TestSample5(t *testing.T) {
	segs := [][]int{
		{9, 22, 14, 20},
		{11, 26, 13, 24},
		{21, 33, 22, 23},
		{21, 33, 25, 32},
		{1, 6, 3, 4},
		{18, 29, 20, 21},
	}
	queries := []int{11, 23, 16, 5, 8, 33, 2, 21}
	expect := []int{32, 32, 32, 5, 8, 33, 4, 32}
	runSample(t, segs, queries, expect)
}
