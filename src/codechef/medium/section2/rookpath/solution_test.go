package main

import "testing"

func runSample(t *testing.T, n int, m int, P [][]int) {
	res := solve(n, m, P)

	vis := make([]bool, m)
	vis[res[0]-1] = true
	vis[res[1]-1] = true
	var row bool
	if sameRow(P[res[0]-1], P[res[1]-1]) {
		row = true
	}

	for i := 2; i < m; i++ {
		if vis[res[i]-1] {
			t.Fatalf("poition %v, already visited", P[res[i]-1])
		}
		vis[res[i]-1] = true
		tmp := sameRow(P[res[i-1]-1], P[res[i]-1])
		if tmp == row {
			t.Fatalf("should change direction, but not")
		}
		row = tmp
	}
}

func sameRow(a, b []int) bool {
	return a[0] == b[0]
}

func TestSample1(t *testing.T) {
	n, m := 2, 4
	P := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}
	runSample(t, n, m, P)
}
