package p2603

import "testing"

func runSample(t *testing.T, coins []int, edges [][]int, expect int) {
	res := collectTheCoins(coins, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	coins := []int{1, 0, 0, 0, 0, 1}
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5},
	}
	expect := 2
	runSample(t, coins, edges, expect)
}

func TestSample2(t *testing.T) {
	coins := []int{0, 0, 0, 1, 1, 0, 0, 1}
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {5, 6}, {5, 7},
	}
	expect := 2
	runSample(t, coins, edges, expect)
}

func TestSample3(t *testing.T) {
	coins := []int{1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1}
	edges := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}, {5, 6}, {6, 7}, {3, 8}, {6, 9}, {7, 10}, {10, 11}, {10, 12}, {7, 13}, {12, 14}, {13, 15}, {14, 16}, {15, 17}, {10, 18},
	}
	expect := 12
	runSample(t, coins, edges, expect)
}
