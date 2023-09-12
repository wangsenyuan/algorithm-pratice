package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, w int, s string, Q [][]int, expect [][]int) {
	res := solve(w, s, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1003004"
	w := 4
	Q := [][]int{
		{1, 2, 1},
	}
	expect := [][]int{
		{2, 4},
	}
	runSample(t, w, s, Q, expect)
}

func TestSample2(t *testing.T) {
	s := "179572007"
	w := 4
	Q := [][]int{
		{2, 7, 3},
		{2, 7, 4},
	}
	expect := [][]int{
		{1, 5},
		{1, 2},
	}
	runSample(t, w, s, Q, expect)
}

func TestSample3(t *testing.T) {
	s := "111"
	w := 2
	Q := [][]int{
		{2, 2, 6},
	}
	expect := [][]int{
		{-1, -1},
	}
	runSample(t, w, s, Q, expect)
}

func TestSample4(t *testing.T) {
	s := "0000"
	w := 1
	Q := [][]int{
		{1, 4, 0},
		{1, 4, 1},
	}
	expect := [][]int{
		{1, 2},
		{-1, -1},
	}
	runSample(t, w, s, Q, expect)
}

func TestSample5(t *testing.T) {
	s := "484"
	w := 1
	Q := [][]int{
		{2, 2, 0},
		{2, 3, 7},
		{1, 2, 5},
		{3, 3, 8},
		{2, 2, 6},
	}
	expect := [][]int{
		{1, 3},
		{1, 3},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}
	runSample(t, w, s, Q, expect)
}
