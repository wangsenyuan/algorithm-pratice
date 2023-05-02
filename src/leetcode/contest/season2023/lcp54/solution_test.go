package lcp54

import "testing"

func runSample(t *testing.T, costs []int, roads [][]int, expect int64) {
	res := minimumCost(costs, roads)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cost := []int{1, 2, 3, 4, 5, 6}
	roads := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {2, 3}, {1, 2}, {2, 4}, {2, 5},
	}
	var expect int64 = 6
	runSample(t, cost, roads, expect)
}

func TestSample2(t *testing.T) {
	cost := []int{3, 2, 1, 4}
	roads := [][]int{
		{0, 2}, {2, 3}, {3, 1},
	}
	var expect int64 = 2
	runSample(t, cost, roads, expect)
}

func TestSample3(t *testing.T) {
	cost := []int{9, 2, 3, 4, 5, 6, 7}
	roads := [][]int{
		{1, 2}, {1, 3}, {2, 3}, {3, 6}, {6, 0}, {0, 3}, {4, 2}, {2, 5}, {4, 5},
	}
	var expect int64 = 5
	runSample(t, cost, roads, expect)
}

func TestSample4(t *testing.T) {
	cost := []int{11, 5, 4, 11, 7, 11, 6, 7, 16, 7, 15, 6, 7, 10, 4, 14, 18, 6, 3, 18}
	roads := [][]int{
		{0, 1}, {0, 3}, {0, 5}, {0, 9}, {0, 14}, {0, 4}, {0, 16}, {0, 12}, {0, 8}, {0, 7}, {1, 2}, {1, 13}, {1, 7}, {1, 3}, {1, 16}, {1, 17}, {1, 10}, {1, 9}, {2, 8}, {2, 10}, {2, 9}, {2, 6}, {2, 13}, {2, 11}, {2, 16}, {2, 18}, {2, 3}, {3, 4}, {3, 7}, {3, 15}, {3, 16}, {3, 11}, {4, 6}, {4, 12}, {4, 19}, {4, 8}, {4, 5}, {4, 14}, {4, 7}, {4, 10}, {5, 13}, {5, 7}, {5, 19}, {5, 15}, {5, 16}, {6, 7}, {6, 15}, {6, 13}, {6, 14}, {7, 11}, {7, 16}, {7, 15}, {7, 14}, {7, 8}, {8, 17}, {8, 11}, {8, 13}, {8, 10}, {9, 11}, {9, 12}, {9, 19}, {10, 14}, {10, 11}, {10, 13}, {10, 15}, {10, 19}, {11, 14}, {11, 16}, {11, 17}, {12, 16}, {12, 18}, {12, 14}, {13, 15}, {14, 15}, {14, 16}, {14, 18}, {14, 17}, {15, 18}, {15, 16}, {16, 18}, {16, 17}, {16, 19}, {17, 19},
	}
	var expect int64 = 3
	runSample(t, cost, roads, expect)
}

func TestSample5(t *testing.T) {
	cost := []int{12, 10, 9, 2, 15, 15, 14, 14, 7, 5, 16, 12, 11}
	roads := [][]int{
		{0, 1}, {0, 2}, {0, 3}, {0, 12}, {1, 5}, {1, 11}, {1, 3}, {2, 12}, {3, 4}, {4, 9}, {5, 6}, {5, 11}, {6, 7}, {6, 10}, {6, 12}, {6, 8}, {7, 8}, {7, 9}, {8, 12}, {10, 11},
	}
	var expect int64 = 2
	runSample(t, cost, roads, expect)
}
