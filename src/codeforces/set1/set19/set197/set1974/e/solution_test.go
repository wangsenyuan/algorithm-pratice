package main

import "testing"

func runSample(t *testing.T, x int, records [][]int, expect int) {
	res := solve(x, records)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 10
	records := [][]int{
		{1, 5},
	}
	expect := 0
	runSample(t, x, records, expect)
}

func TestSample2(t *testing.T) {
	x := 80
	records := [][]int{
		{0, 10},
		{200, 100},
	}
	expect := 10
	runSample(t, x, records, expect)
}

func TestSample3(t *testing.T) {
	x := 100
	records := [][]int{
		{70, 100},
		{100, 200},
		{150, 150},
	}
	expect := 200
	runSample(t, x, records, expect)
}

func TestSample4(t *testing.T) {
	x := 8
	records := [][]int{
		{3, 1},
		{5, 3},
		{3, 4},
		{1, 5},
		{5, 3},
	}
	expect := 15
	runSample(t, x, records, expect)
}

func TestSample5(t *testing.T) {
	x := 5
	records := [][]int{
		{1, 5},
		{2, 1},
	}
	expect := 1
	runSample(t, x, records, expect)
}

func TestSample6(t *testing.T) {
	x := 3
	records := [][]int{
		{2, 5},
		{2, 4},
		{4, 1},
		{5, 1},
		{3, 4},
	}
	expect := 9
	runSample(t, x, records, expect)
}

func TestSample7(t *testing.T) {
	x := 2
	records := [][]int{
		{2, 1},
		{1, 2},
		{3, 5},
		{3, 2},
		{3, 2},
	}
	expect := 9
	runSample(t, x, records, expect)
}
