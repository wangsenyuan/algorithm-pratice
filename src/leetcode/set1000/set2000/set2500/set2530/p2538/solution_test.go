package p2538

import (
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, price []int, expect int64) {
	res := maxOutput(n, edges, price)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}}
	price := []int{9, 8, 7, 6, 10, 5}
	var expect int64 = 24
	runSample(t, n, edges, price, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	price := []int{1, 1, 1}
	var expect int64 = 2
	runSample(t, n, edges, price, expect)
}
