package main

import "testing"

func runSample(t *testing.T, x int, y int, d int, vectors [][]int, expect string) {
	res := solve(x, y, d, vectors)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y, d := 0, 0, 3
	vectors := [][]int{
		{1, 1},
		{1, 2},
	}
	expect := "Anton"

	runSample(t, x, y, d, vectors, expect)
}

func TestSample2(t *testing.T) {
	x, y, d := 0, 0, 4
	vectors := [][]int{
		{1, 1},
		{1, 2},
	}
	expect := "Dasha"

	runSample(t, x, y, d, vectors, expect)
}
