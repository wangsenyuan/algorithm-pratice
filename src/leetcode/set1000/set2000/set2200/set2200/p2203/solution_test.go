package p2203

import "testing"

func runSample(t *testing.T, n int, edges [][]int, src1 int, src2 int, dest int, expect int64) {
	res := minimumWeight(n, edges, src1, src2, dest)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{0, 2, 2}, {0, 5, 6}, {1, 0, 3}, {1, 4, 5}, {2, 1, 1}, {2, 3, 3}, {2, 3, 4}, {3, 4, 2}, {4, 5, 1},
	}
	src1 := 0
	src2 := 1
	dest := 5
	var expect int64 = 9
	runSample(t, n, edges, src1, src2, dest, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 1}, {2, 1, 1},
	}
	src1 := 0
	src2 := 1
	dest := 2
	var expect int64 = -1
	runSample(t, n, edges, src1, src2, dest, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	edges := [][]int{
		{4, 7, 24}, {1, 3, 30}, {4, 0, 31}, {1, 2, 31}, {1, 5, 18}, {1, 6, 19}, {4, 6, 25}, {5, 6, 32}, {0, 6, 50},
	}
	src1 := 4
	src2 := 1
	dest := 6
	var expect int64 = 44
	runSample(t, n, edges, src1, src2, dest, expect)
}
