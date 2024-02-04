package p3024

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := numberOfPairs(points)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{
		{1, 1}, {2, 2}, {3, 3},
	}
	expect := 0
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{
		{6, 2}, {4, 4}, {2, 6},
	}
	expect := 2
	runSample(t, points, expect)
}

func TestSample3(t *testing.T) {
	points := [][]int{
		{3, 1}, {1, 3}, {1, 1},
	}
	expect := 2
	runSample(t, points, expect)
}

func TestSample4(t *testing.T) {
	points := [][]int{
		{0, 3}, {2, 4}, {0, 6},
	}
	expect := 2
	runSample(t, points, expect)
}
