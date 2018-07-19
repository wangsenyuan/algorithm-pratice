package main

import "testing"

func runSmall(t *testing.T, H, W, R int, rocks [][]int, expect int) {
	res := solveSmall(H, W, R, rocks)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", H, W, R, rocks, expect, res)
	}
}

func TestSample1(t *testing.T) {
	H, W, R := 1, 1, 0
	runSmall(t, H, W, R, nil, 1)
}

func TestSample2(t *testing.T) {
	H, W, R := 4, 4, 1
	rocks := [][]int{{2, 1}}
	runSmall(t, H, W, R, rocks, 2)
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
