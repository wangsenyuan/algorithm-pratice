package p2768

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, n int, coords [][]int, expect []int64) {
	res := countBlackBlocks(m, n, coords)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 3, 3
	coords := [][]int{{0, 0}}
	expect := []int64{3, 1, 0, 0, 0}
	runSample(t, m, n, coords, expect)
}

func TestSample2(t *testing.T) {
	m, n := 3, 3
	coords := [][]int{{0, 0}, {1, 1}, {0, 2}}
	expect := []int64{0, 2, 2, 0, 0}
	runSample(t, m, n, coords, expect)
}

func TestSample3(t *testing.T) {
	m, n := 32, 32
	coords := [][]int{
		{17, 29}, {29, 16}, {19, 20}, {18, 9}, {16, 7}, {20, 25}, {22, 19}, {4, 9}, {14, 17}, {6, 23}, {2, 2}, {20, 1}, {8, 7}, {4, 7}, {14, 14}, {10, 10}, {1, 27}, {18, 23}, {6, 30}, {8, 18}, {26, 23}, {25, 8}, {5, 6}, {3, 4},
	}
	expect := []int64{866, 94, 1, 0, 0}
	runSample(t, m, n, coords, expect)
}
