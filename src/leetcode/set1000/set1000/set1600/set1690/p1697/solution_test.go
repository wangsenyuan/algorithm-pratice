package p1697

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edgeList [][]int, queries [][]int, expect []bool) {
	res := distanceLimitedPathsExist(n, edgeList, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16},
	}
	queries := [][]int{
		{0, 1, 2}, {0, 2, 5},
	}
	expect := []bool{false, true}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13},
	}
	queries := [][]int{
		{0, 4, 14}, {1, 4, 13},
	}
	expect := []bool{true, false}
	runSample(t, n, edges, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 34
	edges := [][]int{
		{17, 26, 57}, {4, 15, 100}, {31, 18, 23}, {6, 18, 32}, {12, 26, 19}, {8, 9, 14}, {18, 33, 99}, {4, 30, 60}, {27, 5, 52}, {5, 12, 31}, {30, 16, 63}, {31, 25, 15}, {32, 5, 89}, {27, 13, 31}, {11, 5, 30}, {18, 30, 5}, {10, 31, 21}, {30, 16, 100}, {1, 5, 15}, {24, 11, 61}, {2, 10, 97}, {20, 32, 12}, {24, 17, 1}, {33, 5, 27}, {11, 6, 71}, {16, 31, 93}, {16, 32, 59}, {12, 31, 28}, {22, 13, 79}, {19, 24, 39}, {28, 17, 36}, {3, 9, 82}, {19, 10, 88}, {9, 23, 89}, {17, 2, 70}, {25, 2, 66}, {30, 8, 74}, {26, 0, 77}, {25, 23, 86}, {7, 4, 48}, {22, 11, 21}, {20, 24, 46}, {30, 20, 41}, {16, 20, 48}, {28, 0, 77},
	}
	queries := [][]int{
		{14, 25, 32}, {16, 7, 80}, {16, 2, 2}, {2, 24, 3}, {17, 4, 30}, {30, 6, 56}, {0, 4, 89}, {5, 15, 89}, {10, 18, 40}, {8, 1, 37}, {19, 2, 13}, {1, 28, 66}, {7, 16, 24}, {13, 33, 67}, {32, 31, 26}, {12, 29, 88}, {8, 24, 72}, {30, 16, 87}, {6, 26, 14}, {28, 25, 100}, {12, 13, 54}, {11, 21, 59}, {26, 8, 60}, {26, 19, 27}, {20, 26, 80}, {30, 11, 21}, {15, 7, 49}, {26, 28, 81}, {4, 32, 22}, {21, 26, 32}, {22, 25, 46}, {9, 6, 13}, {21, 26, 9}, {14, 29, 33}, {5, 3, 13}, {23, 27, 84}, {7, 6, 27}, {28, 25, 50}, {2, 10, 78}, {33, 14, 62}, {12, 19, 76}, {29, 26, 26}, {6, 8, 9}, {7, 18, 56}, {26, 17, 91},
	}
	expect := []bool{false, true, false, false, false, true, true, false, true, false, false, true, false, true, false, false, false, true, false, true, true, false, false, false, true, false, false, true, false, false, true, false, false, false, false, false, false, true, true, false, true, false, false, false, true}
	runSample(t, n, edges, queries, expect)
}
