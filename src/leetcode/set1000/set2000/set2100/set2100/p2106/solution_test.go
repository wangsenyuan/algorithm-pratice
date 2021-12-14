package p2106

import "testing"

func runSample(t *testing.T, fruits [][]int, startPos int, k int, expect int) {
	res := maxTotalFruits(fruits, startPos, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	fruits := [][]int{
		{2, 8}, {6, 3}, {8, 6},
	}
	startPos := 5
	k := 4
	expect := 9
	runSample(t, fruits, startPos, k, expect)
}

func TestSample2(t *testing.T) {
	fruits := [][]int{
		{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9},
	}
	startPos := 5
	k := 4
	expect := 14
	runSample(t, fruits, startPos, k, expect)
}

func TestSample3(t *testing.T) {
	fruits := [][]int{
		{0, 3}, {6, 4}, {8, 5},
	}
	startPos := 3
	k := 2
	expect := 0
	runSample(t, fruits, startPos, k, expect)
}

func TestSample4(t *testing.T) {
	fruits := [][]int{
		{0, 10000},
	}
	startPos := 20000
	k := 20000
	expect := 10000
	runSample(t, fruits, startPos, k, expect)
}

func TestSample5(t *testing.T) {
	fruits := [][]int{
		{0, 7}, {7, 4}, {9, 10}, {12, 6}, {14, 8}, {16, 5}, {17, 8}, {19, 4}, {20, 1}, {21, 3}, {24, 3}, {25, 3}, {26, 1}, {28, 10}, {30, 9}, {31, 6}, {32, 1}, {37, 5}, {40, 9},
	}
	startPos := 21
	k := 30
	expect := 71
	runSample(t, fruits, startPos, k, expect)
}
