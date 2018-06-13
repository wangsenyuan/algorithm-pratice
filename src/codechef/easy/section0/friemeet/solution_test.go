package main

import "testing"

func TestSample1(t *testing.T) {
	n, m := 6, 2
	edges := [][]int{
		[]int{1, 3, 1},
		[]int{2, 3, 2},
		[]int{3, 4, 3},
		[]int{4, 5, 4},
		[]int{4, 6, 5},
	}
	special := []bool{true, false, false, false, true, false}
	a, b := solve(n, m, edges, special)

	if a != 4 && b != 1 {
		t.Errorf("this case should give 4 1, but got %d %d", a, b)
	}
}

func TestSample2(t *testing.T) {
	n, m := 6, 6
	edges := [][]int{
		[]int{1, 3, 1},
		[]int{2, 3, 2},
		[]int{3, 4, 3},
		[]int{4, 5, 4},
		[]int{4, 6, 5},
	}
	special := []bool{true, true, true, true, true, true}
	a, b := solve(n, m, edges, special)

	if a != 29 && b != 6 {
		t.Errorf("this case should give 29 6, but got %d %d", a, b)
	}
}
