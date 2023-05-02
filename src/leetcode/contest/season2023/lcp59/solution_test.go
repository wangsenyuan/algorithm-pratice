package lcp59

import "testing"

func runSample(t *testing.T, num int, wood [][]int, expect int64) {
	res := buildBridge(num, wood)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 10
	wood := [][]int{
		{1, 2}, {4, 7}, {8, 9},
	}
	var expect int64 = 3
	runSample(t, num, wood, expect)
}

func TestSample2(t *testing.T) {
	num := 10
	wood := [][]int{
		{1, 5}, {1, 1}, {10, 10}, {6, 7}, {7, 8},
	}
	var expect int64 = 10
	runSample(t, num, wood, expect)
}

func TestSample3(t *testing.T) {
	num := 10
	wood := [][]int{
		{1, 2}, {2, 4},
	}
	var expect int64 = 0
	runSample(t, num, wood, expect)
}
