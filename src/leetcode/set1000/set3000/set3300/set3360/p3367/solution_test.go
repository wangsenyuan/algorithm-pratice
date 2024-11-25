package p3367

import "testing"

func runSample(t *testing.T, edges [][]int, k int, expect int) {
	res := maximizeSumOfWeights(edges, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1, 4}, {0, 2, 2}, {2, 3, 12}, {2, 4, 6},
	}
	k := 2
	expect := 22
	runSample(t, edges, k, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1, 5}, {1, 2, 10}, {0, 3, 15}, {3, 4, 20}, {3, 5, 5}, {0, 6, 10},
	}
	k := 3
	expect := 65
	runSample(t, edges, k, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{0, 1, 3}, {1, 2, 20}, {1, 3, 10},
	}
	k := 2
	expect := 30
	runSample(t, edges, k, expect)
}
