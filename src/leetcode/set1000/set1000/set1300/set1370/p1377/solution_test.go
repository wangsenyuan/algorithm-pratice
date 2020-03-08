package p1377

import (
	"math"
	"testing"
)

func runSample(tc *testing.T, n int, edges [][]int, t int, target int, expect float64) {
	res := frogPosition(n, edges, t, target)

	if math.Abs(res-expect) > 1e-7 {
		tc.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(tc *testing.T) {
	n := 7
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	t := 2
	target := 4
	expect := 0.16666666666666666
	runSample(tc, n, edges, t, target, expect)
}

func TestSample2(tc *testing.T) {
	n := 7
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	t := 1
	target := 7
	expect := 0.3333333333333333
	runSample(tc, n, edges, t, target, expect)
}

func TestSample3(tc *testing.T) {
	n := 7
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	t := 20
	target := 6
	expect := 0.16666666666666666
	runSample(tc, n, edges, t, target, expect)
}
