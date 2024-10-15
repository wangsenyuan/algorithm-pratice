package main

import "testing"

func runSample(t *testing.T, monsters []int, heros [][]int, expect int) {
	res := solve(monsters, heros)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 11, 14, 1, 8}
	b := [][]int{
		{3, 2},
		{100, 1},
	}
	expect := 5
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{78, 22, 90, 12, 42}
	b := [][]int{
		{51, 1},
		{87, 3},
		{48, 2},
		{13, 1},
		{98, 5},
	}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3, 1}
	b := [][]int{
		{3, 3},
		{2, 4},
	}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 3, 1, 4}
	b := [][]int{
		{4, 3},
		{2, 4},
	}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 1, 3}
	b := [][]int{
		{3, 1},
		{2, 4},
	}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 2, 2}
	b := [][]int{
		{3, 2},
	}
	expect := 2
	runSample(t, a, b, expect)
}
