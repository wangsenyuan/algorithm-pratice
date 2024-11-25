package p3363

import "testing"

func runSample(t *testing.T, fruits [][]int, expect int) {
	res := maxCollectedFruits(fruits)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	fruits := [][]int{
		{1, 2, 3, 4},
		{5, 6, 8, 7},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expect := 100
	runSample(t, fruits, expect)
}

func TestSample2(t *testing.T) {
	fruits := [][]int{
		{1, 1},
		{1, 1},
	}
	expect := 4
	runSample(t, fruits, expect)
}

