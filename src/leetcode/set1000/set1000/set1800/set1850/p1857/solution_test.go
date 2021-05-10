package p1857

import "testing"

func runSample(t *testing.T, colors string, edges [][]int, expect int) {
	res := largestPathValue(colors, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	colors := "abaca"
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {3, 4},
	}
	expect := 3
	runSample(t, colors, edges, expect)
}

func TestSample2(t *testing.T) {
	colors := "a"
	edges := [][]int{
		{0, 0},
	}
	expect := 0
	runSample(t, colors, edges, expect)
}

func TestSample3(t *testing.T) {
	colors := "hhqhuqhqff"
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9},
	}
	expect := 3
	runSample(t, colors, edges, expect)
}
