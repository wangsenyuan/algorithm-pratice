package main

import "testing"

func runSample(t *testing.T, boxes [][]int) {
	res := solve(boxes)
	if len(res) != (len(boxes)+1)/2 {
		t.Fatalf("Sample result %v not correct", res)
	}
	sum := make([]int, 2)
	for i := 0; i < len(boxes); i++ {
		sum[0] += boxes[i][0]
		sum[1] += boxes[i][1]
	}

	got := make([]int, 2)
	for _, i := range res {
		got[0] += boxes[i-1][0]
		got[1] += boxes[i-1][1]
	}

	if got[0]*2 < sum[0] || got[1]*2 < sum[1] {
		t.Fatalf("Sample result %v not correct", res)
	}
}

func TestSample1(t *testing.T) {
	boxes := [][]int{
		{10, 15},
		{5, 7},
		{20, 18},
	}
	runSample(t, boxes)
}

func TestSample2(t *testing.T) {
	boxes := [][]int{
		{0, 0},
	}
	runSample(t, boxes)
}
