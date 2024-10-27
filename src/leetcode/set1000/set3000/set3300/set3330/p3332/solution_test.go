package p3332

import "testing"

func runSample(t *testing.T, n int, k int, stayScore [][]int, travelScore [][]int, expect int) {
	res := maxScore(n, k, stayScore, travelScore)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	k := 1
	stay := [][]int{
		{2, 3},
	}
	travel := [][]int{
		{0, 2}, {1, 0},
	}
	expect := 3
	runSample(t, n, k, stay, travel, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 2
	stay := [][]int{
		{3, 4, 2},
		{2, 1, 2},
	}
	travel := [][]int{
		{0, 2, 1}, {2, 0, 4}, {3, 2, 0},
	}
	expect := 8
	runSample(t, n, k, stay, travel, expect)
}
