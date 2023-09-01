package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect [][]int) {
	res := solve(A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 5}
	B := []int{7, 11, 13}
	expect := [][]int{
		{5, 4, 2},
		{11, 10, 8},
	}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1000}
	B := []int{5000}
	expect := [][]int{
		{4000},
		{4000},
	}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := []int{1, 2, 3, 4}
	expect := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 20, 30, 40}
	B := []int{22, 33, 33, 55}
	expect := [][]int{
		{12, 2, 3, 15},
		{23, 13, 3, 15},
	}
	runSample(t, A, B, expect)
}
