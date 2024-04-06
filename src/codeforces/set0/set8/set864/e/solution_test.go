package main

import "testing"

func runSample(t *testing.T, items [][]int, expect int) {
	res, ord := solve(len(items), items)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var time int
	var sum int
	for _, i := range ord {
		i--
		if time+items[i][0] > items[i][1] {
			t.Fatalf("Sample result %v, not correct", ord)
		}
		time += items[i][0]
		sum += items[i][2]
	}

	if sum != expect {
		t.Fatalf("Sample result %v, not correct", ord)
	}
}

func TestSample1(t *testing.T) {
	items := [][]int{
		{3, 7, 4},
		{2, 6, 5},
		{3, 7, 6},
	}
	expect := 11

	runSample(t, items, expect)
}

func TestSample2(t *testing.T) {
	items := [][]int{
		{5, 6, 1},
		{3, 3, 5},
	}
	expect := 1

	runSample(t, items, expect)
}

func TestSample3(t *testing.T) {
	items := [][]int{
		{12, 44, 17},
		{10, 12, 11},
		{16, 46, 5},
		{17, 55, 5},
		{6, 60, 2},
	}
	expect := 35

	runSample(t, items, expect)
}
