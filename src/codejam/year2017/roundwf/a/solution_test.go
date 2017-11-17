package main

import "testing"

func TestSample1(t *testing.T) {
	n := 4
	dices := [][]int{
		{4, 8, 15, 16, 23, 42},
		{8, 6, 7, 5, 30, 9},
		{1, 2, 3, 4, 55, 6},
		{2, 10, 18, 36, 54, 86},
	}
	res := solve(n, dices)

	if res != 4 {
		t.Errorf("sample should give answer 4, but %d", res)
	}
}

func TestSample2(t *testing.T) {
	n := 2
	dices := [][]int{
		{1, 2, 3, 4, 5, 6},
		{60, 50, 40, 30, 20, 10},
	}
	res := solve(n, dices)

	if res != 1 {
		t.Errorf("sample should give answer 1, but %d", res)
	}
}

func TestSample3(t *testing.T) {
	n := 3
	dices := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	res := solve(n, dices)

	if res != 3 {
		t.Errorf("sample should give answer 3, but %d", res)
	}
}
