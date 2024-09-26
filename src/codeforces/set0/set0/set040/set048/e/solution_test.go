package main

import "testing"

func runSample(t *testing.T, head int, tail int, R int, op1 [][]int, op2 [][]int, who string, expect int) {
	winner, cnt := solve(head, tail, R, op1, op2)

	if winner != who || expect != cnt {
		t.Fatalf("Sample expect %s %d, but got %s %d", winner, expect, who, cnt)
	}
}

func TestSample1(t *testing.T) {
	head, tail, R := 2, 2, 4
	op1 := [][]int{
		{1, 0},
		{0, 1},
	}
	op2 := [][]int{
		{0, 1},
		{0, 1},
		{0, 0},
	}
	winner := "Ivan"
	expect := 2
	runSample(t, head, tail, R, op1, op2, winner, expect)
}

func TestSample2(t *testing.T) {
	head, tail, R := 2, 2, 4
	op1 := [][]int{
		{0, 1},
	}
	op2 := [][]int{
		{1, 0},
	}
	winner := "Draw"
	expect := 0
	runSample(t, head, tail, R, op1, op2, winner, expect)
}

func TestSample3(t *testing.T) {
	head, tail, R := 2, 2, 5
	op1 := [][]int{
		{1, 1},
	}
	op2 := [][]int{
		{3, 0},
	}
	winner := "Zmey"
	expect := 2
	runSample(t, head, tail, R, op1, op2, winner, expect)
}

func TestSample4(t *testing.T) {
	head, tail, R := 37, 6, 50
	op1 := [][]int{
		{2, 3},
		{4, 1},
		{1, 0},
		{2, 5},
		{5, 1},
		{4, 4},
		{0, 4},
		{5, 4},
		{0, 3},
		{4, 0},
	}
	op2 := [][]int{
		{3, 3},
		{2, 2},
		{1, 2},
		{1, 3},
		{0, 1},
		{3, 2},
		{1, 1},
		{1, 5},
		{5, 0},
		{5, 3},
		{1, 1},
		{5, 0},
		{5, 1},
		{5, 1},
		{5, 2},
		{1, 3},
		{2, 4},
		{4, 2},
		{3, 3},
		{3, 3},
		{5, 0},
		{2, 2},
		{0, 1},
		{5, 5},
		{1, 1},
		{1, 2},
		{0, 3},
		{4, 5},
		{1, 4},
		{3, 0},
		{5, 2},
		{2, 0},
		{3, 4},
		{0, 2},
		{0, 0},
		{3, 1},
		{1, 2},
		{5, 0},
		{0, 4},
		{3, 0},
	}
	winner := "Ivan"
	expect := 8
	runSample(t, head, tail, R, op1, op2, winner, expect)
}

func TestSample5(t *testing.T) {
	head, tail, R := 10, 10, 30
	op1 := [][]int{
		{3, 2},
		{4, 1},
		{3, 3},
		{4, 2},
		{2, 5},
		{4, 5},
		{5, 7},
	}
	op2 := [][]int{
		{0, 2},
		{3, 1},
		{4, 3},
		{3, 4},
		{2, 5},
		{6, 2},
		{5, 6},
	}
	winner := "Zmey"
	expect := 11
	runSample(t, head, tail, R, op1, op2, winner, expect)
}

func TestSample6(t *testing.T) {
	head, tail, R := 0, 1, 200
	op1 := [][]int{
		{0, 2},
	}
	op2 := [][]int{
		{1, 0},
	}
	winner := "Zmey"
	expect := 599
	runSample(t, head, tail, R, op1, op2, winner, expect)
}
