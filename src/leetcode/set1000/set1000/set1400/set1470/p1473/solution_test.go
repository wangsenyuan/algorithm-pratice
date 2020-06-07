package p1473

import "testing"

func runSample(t *testing.T, houses []int, cost [][]int, m int, n int, target int, expect int) {
	res := minCost(houses, cost, m, n, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	houses := []int{0, 0, 0, 0, 0}
	cost := [][]int{
		{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1},
	}
	m := 5
	n := 2
	target := 3
	expect := 9
	runSample(t, houses, cost, m, n, target, expect)
}

func TestSample2(t *testing.T) {
	houses := []int{0, 2, 1, 2, 0}
	cost := [][]int{
		{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1},
	}
	m := 5
	n := 2
	target := 3
	expect := 11
	runSample(t, houses, cost, m, n, target, expect)
}

func TestSample3(t *testing.T) {
	houses := []int{0, 0, 0, 0, 0}
	cost := [][]int{
		{1, 10}, {10, 1}, {1, 10}, {10, 1}, {1, 10},
	}
	m := 5
	n := 2
	target := 5
	expect := 5
	runSample(t, houses, cost, m, n, target, expect)
}
