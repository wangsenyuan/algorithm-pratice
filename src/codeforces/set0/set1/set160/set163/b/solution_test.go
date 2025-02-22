package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, k int, h int, W []int, V []int, expect []int) {
	res := solve(n, k, h, W, V)
	a := getMinTime(k, h, V, expect)
	b := getMinTime(k, h, V, res)

	if math.Abs(a-b) > 1e-7 {
		t.Fatalf("Sample result %v gives wrong min time %.8f, want %.8f", res, b, a)
	}
	for i := 0; i+1 < k; i++ {
		x := res[i] - 1
		y := res[i+1] - 1
		if W[x] > W[y] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func getMinTime(k int, h int, V []int, arr []int) float64 {
	var res float64
	for i := 0; i < k; i++ {
		id := arr[i] - 1
		res = math.Max(res, float64((i+1)*h)/float64(V[id]))
	}
	return res
}

func TestSample1(t *testing.T) {
	n, k, h := 5, 3, 2
	W := []int{1, 2, 3, 2, 1}
	V := []int{1, 2, 1, 2, 10}
	expect := []int{5, 2, 4}
	runSample(t, n, k, h, W, V, expect)
}

func TestSample2(t *testing.T) {
	n, k, h := 5, 3, 10
	W := []int{3, 4, 3, 2, 1}
	V := []int{5, 4, 3, 2, 1}
	expect := []int{4, 3, 1}
	runSample(t, n, k, h, W, V, expect)
}

func TestSample3(t *testing.T) {
	n, k, h := 50, 25, 3
	W := []int{1, 3, 1, 2, 1, 2, 1, 2, 3, 1, 1, 2, 2, 3, 3, 1, 1, 3,
		2, 3, 2, 1, 2, 1, 2, 1, 3, 2, 1, 3, 3,
		3, 2, 3, 1, 1, 1, 2, 3, 3, 1, 3, 1, 2, 2, 2, 3, 1, 3, 1}
	V := []int{1, 2, 2, 3, 3, 1, 3, 2, 2, 2, 2, 2, 2, 3, 3, 2, 1, 2,
		3, 2, 3, 2, 1, 3, 2, 1, 3, 3, 3, 3, 2, 3, 2, 2, 1, 3, 3, 3,
		2, 3, 1, 3, 2, 2, 2, 1, 3, 2, 2, 2}
	expect := []int{36, 37, 6, 23, 46, 8, 12, 13, 25, 33, 44, 45, 4, 19, 21, 28, 38, 14, 15, 27, 30, 32, 40, 42, 47}
	runSample(t, n, k, h, W, V, expect)
}
