package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q [][]int, expect []int) {
	res := solve(Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := [][]int{
		{1, 8},
		{1, 10},
		{1, 6},
		{3},
		{2},
		{1, 9},
		{2},
		{3},
	}
	expect := []int{2, 1, 3, 4}
	runSample(t, Q, expect)
}

func TestSample2(t *testing.T) {
	Q := [][]int{
		{1, 8},
		{1, 10},
		{1, 8},
		{3},
		{3},
		{3},
	}
	expect := []int{2, 1, 3}
	runSample(t, Q, expect)
}

func TestSample3(t *testing.T) {
	Q := [][]int{
		{1, 103913},
		{3},
		{1, 103913},
		{1, 103913},
		{3},
		{1, 103913},
		{1, 103913},
		{2},
	}
	expect := []int{1, 2, 3}
	runSample(t, Q, expect)
}
