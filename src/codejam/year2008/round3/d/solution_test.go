package main

import "testing"

func runSmall(t *testing.T, H, W, R int, rocks [][]int, expect int) {
	res := solveSmall(H, W, R, rocks)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", H, W, R, rocks, expect, res)
	}
}

func runSample(t *testing.T, H, W, R int, rocks [][]int, expect int) {
	rcks := make([][]int64, R)
	for i := 0; i < R; i++ {
		rcks[i] = []int64{
			int64(rocks[i][0]), int64(rocks[i][1]),
		}
	}
	res := solve(int64(H), int64(W), R, rcks)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", H, W, R, rocks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	H, W, R := 1, 1, 0
	// runSmall(t, H, W, R, nil, 1)
	runSample(t, H, W, R, nil, 1)
}

func TestSample2(t *testing.T) {
	H, W, R := 4, 4, 1
	rocks := [][]int{{2, 1}}
	runSample(t, H, W, R, rocks, 2)
}

func TestSample3(t *testing.T) {
	H, W, R := 3, 3, 0
	runSmall(t, H, W, R, nil, 0)
}

func TestSample4(t *testing.T) {
	H, W, R := 7, 10, 2
	rocks := [][]int{{1, 2}, {7, 1}}
	runSmall(t, H, W, R, rocks, 5)
}

func TestSample5(t *testing.T) {
	H, W, R := 4, 4, 1
	rocks := [][]int{{3, 2}}
	runSmall(t, H, W, R, rocks, 1)
}
