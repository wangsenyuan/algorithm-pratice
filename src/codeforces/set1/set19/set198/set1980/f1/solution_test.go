package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, fountains [][]int, plots int, expect []int) {
	alice, res := solve(n, m, fountains)

	if !reflect.DeepEqual(res, expect) || alice != plots {
		t.Fatalf("Sample expect %d, %v, but got %d %v", plots, expect, alice, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	fountains := [][]int{
		{1, 1},
		{1, 2},
		{2, 2},
	}
	plots := 1
	expect := []int{1, 0, 1}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 5
	fountains := [][]int{
		{1, 2},
		{2, 2},
		{3, 4},
		{4, 3},
	}
	plots := 11
	expect := []int{0, 1, 0, 1}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 5
	fountains := [][]int{
		{1, 2},
		{1, 5},
		{1, 1},
		{2, 2},
		{2, 4},
		{2, 5},
		{1, 4},
		{2, 3},
		{1, 3},
	}
	plots := 1
	expect := []int{0, 0, 1, 1, 0, 0, 0, 0, 0}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample4(t *testing.T) {
	n, m := 6, 4
	fountains := [][]int{
		{6, 2},
		{1, 3},
		{1, 4},
		{1, 2},
	}
	plots := 6
	expect := []int{1, 0, 0, 0}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample5(t *testing.T) {
	n, m := 3, 4
	fountains := [][]int{
		{2, 1},
		{3, 2},
		{1, 4},
		{1, 3},
		{2, 4},
	}
	plots := 1
	expect := []int{1, 1, 0, 0, 0}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample6(t *testing.T) {
	n, m := 4, 4
	fountains := [][]int{
		{3, 4},
		{2, 4},
		{1, 1},
		{2, 3},
	}
	plots := 9
	expect := []int{1, 0, 1, 1}
	runSample(t, n, m, fountains, plots, expect)
}

func TestSample7(t *testing.T) {
	n, m := 6, 5
	fountains := [][]int{
		{2, 4},
		{4, 2},
		{2, 1},
		{1, 2},
		{5, 2},
		{1, 4},
		{4, 3},
		{3, 3},
		{6, 3},
		{2, 2},
	}
	plots := 5
	expect := []int{0, 0, 1, 0, 1, 0, 0, 0, 1, 0}
	runSample(t, n, m, fountains, plots, expect)
}
