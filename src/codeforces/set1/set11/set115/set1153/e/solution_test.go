package main

import "testing"

func runSample(t *testing.T, n int, snakes [][]int) {
	var num int
	ask := func(r1 int, c1 int, r2 int, c2 int) int {
		num++
		if num > 2019 {
			t.Fatalf("Sample asked too much")
		}
		var cnt int
		for i := 1; i < len(snakes); i++ {
			x1, y1 := snakes[i-1][0], snakes[i-1][1]
			x2, y2 := snakes[i][0], snakes[i][1]
			// x1 == x2 or y1 == y2
			if x1 == x2 {
				// 水平方向
				if x1 >= r1 && x1 <= r2 {
					y1, y2 = min(y1, y2), max(y1, y2)
					if y2 == c1 || y1 == c2 {
						cnt++
					}
				}
			} else {
				// 垂直方向
				if y1 >= c1 && y1 <= c2 {
					x1, x2 = min(x1, x2), max(x1, x2)
					if x2 == r1 || x1 == r2 {
						cnt++
					}
				}
			}
		}
		return cnt
	}

	res := solve(n, ask)

	rh, rc := res[0], res[1]
	th, tc := res[2], res[3]

	head := snakes[0]
	tail := snakes[len(snakes)-1]

	if rh == head[0] && rc == head[1] && th == tail[0] && tc == tail[1] {
		return
	}

	if rh == tail[0] && rc == tail[1] && th == head[0] && tc == head[1] {
		return
	}

	t.Fatalf("Sample got wrong result %v", res)
}

func TestSample1(t *testing.T) {
	n := 3
	snakes := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{3, 2},
		{3, 3},
		{2, 3},
		{2, 2},
		{1, 2},
		{1, 3},
	}
	runSample(t, n, snakes)
}

func TestSample2(t *testing.T) {
	n := 2
	snakes := [][]int{
		{1, 1},
		{2, 1},
	}
	runSample(t, n, snakes)
}

func TestSample3(t *testing.T) {
	n := 3
	snakes := [][]int{
		{2, 1},
		{2, 2},
		{2, 3},
	}
	runSample(t, n, snakes)
}
