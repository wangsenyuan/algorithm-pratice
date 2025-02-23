package p3464

import "testing"

func runSample(t *testing.T, side int, points [][]int, k int, expect int) {
	res := maxDistance(side, points, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{0, 0},
		{0, 2},
		{2, 2},
		{2, 0},
	}
	side := 2
	k := 4
	expect := 2
	runSample(t, side, points, k, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{0, 0},
		{1, 2},
		{2, 0},
		{2, 2},
		{2, 1},
	}
	side := 2
	k := 4
	expect := 1
	runSample(t, side, points, k, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 2}, {2, 1},
	}
	side := 2
	k := 5
	expect := 1
	runSample(t, side, points, k, expect)
}

func TestSample4(t *testing.T) {
	points := [][]int{
		{8, 0}, {5, 9}, {2, 0}, {4, 9}, {0, 1},
	}
	side := 9
	k := 4
	expect := 3
	runSample(t, side, points, k, expect)
}

func TestSample5(t *testing.T) {
	points := [][]int{
		{0, 7}, {8, 0}, {16, 16}, {1, 0}, {16, 4},
	}
	side := 16
	k := 4
	expect := 12
	runSample(t, side, points, k, expect)
}

func TestSample6(t *testing.T) {
	points := [][]int{
		{4, 2}, {0, 1}, {3, 0}, {1, 4}, {4, 3}, {0, 3}, {0, 0}, {0, 2},
	}
	side := 4
	k := 4
	expect := 4
	runSample(t, side, points, k, expect)
}
