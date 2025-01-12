package p3419

import "testing"

func runSample(t *testing.T, n int, edges [][]int, threshold int, expect int) {
	res := minMaxWeight(n, edges, threshold)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 0, 1}, {2, 0, 2}, {3, 0, 1}, {4, 3, 1}, {2, 1, 1},
	}
	threshold := 2
	expect := 1
	runSample(t, n, edges, threshold, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 1}, {0, 2, 2}, {0, 3, 1}, {0, 4, 1}, {1, 2, 1}, {1, 4, 1},
	}
	threshold := 1
	expect := -1
	runSample(t, n, edges, threshold, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 1}, {1, 3, 3}, {1, 4, 5}, {2, 3, 2}, {3, 4, 2}, {4, 0, 1},
	}
	threshold := 1
	expect := 2
	runSample(t, n, edges, threshold, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 1}, {1, 3, 3}, {1, 4, 5}, {2, 3, 2}, {4, 0, 1},
	}
	threshold := 1
	expect := -1
	runSample(t, n, edges, threshold, expect)
}
