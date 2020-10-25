package p5156

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m [][]int, expect [][]int) {
	res := matrixRankTransform(m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := [][]int{
		{1, 2},
		{2, 3},
	}
	runSample(t, m, expect)
}

func TestSample2(t *testing.T) {
	m := [][]int{
		{7, 7},
		{7, 7},
	}
	expect := [][]int{
		{1, 1},
		{1, 1},
	}
	runSample(t, m, expect)
}

func TestSample3(t *testing.T) {
	m := [][]int{
		{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19},
	}
	expect := [][]int{
		{4, 2, 3}, {1, 3, 4}, {5, 1, 6}, {1, 3, 4},
	}
	runSample(t, m, expect)
}

func TestSample4(t *testing.T) {
	m := [][]int{
		{7, 3, 6}, {1, 4, 5}, {9, 8, 2},
	}
	expect := [][]int{
		{5, 1, 4}, {1, 2, 3}, {6, 3, 1},
	}
	runSample(t, m, expect)
}

func TestSample5(t *testing.T) {
	m := [][]int{
		{-37, -50, -3, 44},
		{-37, 46, 13, -32},
		{47, -42, -3, -40},
		{-17, -22, -39, 24},
	}
	expect := [][]int{
		{2, 1, 4, 6},
		{2, 6, 5, 4},
		{5, 2, 4, 3},
		{4, 3, 1, 5},
	}
	runSample(t, m, expect)
}
