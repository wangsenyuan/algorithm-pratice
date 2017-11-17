package main

import "testing"

func TestSample1(t *testing.T) {
	n, m := 2, 2
	edges := [][]int{
		{1, 2},
		{2, 1},
	}
	_, found := solve(n, m, edges)
	if found {
		t.Errorf("this sample should not have solution, but got one")
	}
}

func TestSample2(t *testing.T) {
	n, m := 3, 2
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	ans, found := solve(n, m, edges)
	if !found {
		t.Errorf("this sample should have solution, but got none")
	} else if ans != 2 {
		t.Errorf("this sample should give answer 2, but got %d", ans)
	}
}
