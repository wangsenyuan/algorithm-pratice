package main

import "testing"

func runSample(t *testing.T, N, K, M int, trips [][]int, expect bool) {
	res := solve(N, K, M, trips)

	if expect && len(res) == 0 {
		t.Errorf("Sample %d %d %d %v, expect result but got nil", N, K, M, trips)
	} else if !expect && len(res) > 0 {
		t.Errorf("Sample %d %d %d %v, expect nil but got %v", N, K, M, trips, res)
	} else if len(res) > 0 && !checkResult(N, trips, res) {
		t.Errorf("Sample %d %d %d %v, result %v, not correct", N, K, M, trips, res)
	} else {
		t.Logf("Sample %d %d %d %v, got result %v", N, K, M, trips, res)
	}
}

func checkResult(N int, trips [][]int, res []int) bool {
	for _, trip := range trips {
		if !checkTrips(N, trip, res) {
			return false
		}
	}
	return true
}

func checkTrips(N int, trip []int, res []int) bool {
	x, y, z := trip[0], trip[1], trip[2]

	if res[z-1] != x {
		return false
	}

	cnt := 1

	for i := 0; i < z-1; i++ {
		if res[i] == x {
			cnt++
		}
	}

	return cnt == y
}

func TestSample1(t *testing.T) {
	N, K, M := 3, 3, 2
	trips := [][]int{
		{1, 1, 2},
		{2, 2, 3},
	}
	runSample(t, N, K, M, trips, true)
}

func TestSample2(t *testing.T) {
	N, K, M := 3, 1, 1
	trips := [][]int{
		{1, 1, 1},
	}
	runSample(t, N, K, M, trips, true)
}

func TestSample3(t *testing.T) {
	N, K, M := 3, 1, 1
	trips := [][]int{
		{1, 1, 2},
	}
	runSample(t, N, K, M, trips, false)
}

func TestSample4(t *testing.T) {
	N, K, M := 5, 2, 2
	trips := [][]int{
		{1, 2, 4},
		{1, 4, 5},
	}
	runSample(t, N, K, M, trips, false)
}

func TestSample5(t *testing.T) {
	N, K, M := 6, 2, 2
	trips := [][]int{
		{1, 2, 4},
		{2, 2, 5},
	}
	runSample(t, N, K, M, trips, false)
}

func TestSample6(t *testing.T) {
	N, K, M := 20, 3, 3
	trips := [][]int{
		{1, 3, 5},
		{2, 3, 10},
		{3, 3, 15},
	}
	runSample(t, N, K, M, trips, true)
}
