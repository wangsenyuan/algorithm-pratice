package p2906

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect [][]int) {
	res := constructProductMatrix(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := [][]int{
		{24, 12},
		{8, 6},
	}
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{12345},
		{2},
		{1},
	}
	expect := [][]int{
		{2},
		{0},
		{0},
	}
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{3, 2, 5},
		{6, 4, 3},
		{6, 3, 1},
	}

	expect := [][]int{
		{615, 7095, 7776},
		{6480, 9720, 615},
		{6480, 615, 1845},
	}
	runSample(t, grid, expect)
}
