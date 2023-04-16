package p2646

import "testing"

func runSample(t *testing.T, n int, edges [][]int, price []int, trips [][]int, expect int) {
	res := minimumTotalPrice(n, edges, price, trips)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3},
	}
	price := []int{2, 2, 10, 6}
	trips := [][]int{
		{0, 3}, {2, 1}, {2, 3},
	}
	expect := 23
	runSample(t, n, edges, price, trips, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	edges := [][]int{
		{0, 1},
	}
	price := []int{2, 2}
	trips := [][]int{
		{0, 0},
	}
	expect := 1
	runSample(t, n, edges, price, trips, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	edges := [][]int{
		{2, 5}, {3, 4}, {4, 1}, {1, 7}, {6, 7}, {7, 0}, {0, 5}, {5, 8},
	}
	price := []int{4, 4, 6, 4, 2, 4, 2, 14, 8}
	trips := [][]int{
		{1, 5}, {2, 7}, {4, 3}, {1, 8}, {2, 8}, {4, 3}, {1, 5}, {1, 4}, {2, 1}, {6, 0},
		{0, 7}, {8, 6}, {4, 0}, {7, 5}, {7, 5}, {6, 0}, {5, 1}, {1, 1}, {7, 5}, {1, 7},
		{8, 7}, {2, 3}, {4, 1}, {3, 5}, {2, 5}, {3, 7}, {0, 1}, {5, 8}, {5, 3}, {5, 2},
	}
	expect := 429
	runSample(t, n, edges, price, trips, expect)
}
