package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, candies [][]int, expect []int) {
	res := solve(candies)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	candies := [][]int{
		{1, 0},
		{4, 1},
		{2, 0},
		{4, 1},
		{5, 1},
		{6, 1},
		{3, 0},
		{2, 0},
	}
	expect := []int{3, 3}
	runSample(t, candies, expect)
}

func TestSample2(t *testing.T) {
	candies := [][]int{
		{1, 1},
		{1, 1},
		{2, 1},
		{2, 1},
	}
	expect := []int{3, 3}
	runSample(t, candies, expect)
}

func TestSample3(t *testing.T) {
	candies := [][]int{
		{2, 0},
		{2, 0},
		{4, 1},
		{4, 1},
		{4, 1},
		{7, 0},
		{7, 1},
		{7, 0},
		{7, 1},
	}
	expect := []int{9, 5}
	runSample(t, candies, expect)
}
