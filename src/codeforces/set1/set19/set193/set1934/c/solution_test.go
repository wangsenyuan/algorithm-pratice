package main

import "testing"

func runSample(t *testing.T, first []int, second []int, n int, m int) {
	ask := func(x int, y int) int {
		return min(distance(first, []int{x, y}), distance(second, []int{x, y}))
	}

	x, y := solve(n, m, ask)

	if distance(first, []int{x, y}) != 0 && distance(second, []int{x, y}) != 0 {
		t.Fatalf("Sample result %d %d, not correct", x, y)
	}

}

func distance(a, b []int) int {
	dx := abs(a[0] - b[0])
	dy := abs(a[1] - b[1])
	return dx + dy
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	first := []int{2, 3}
	second := []int{3, 2}
	n, m := 4, 4
	runSample(t, first, second, n, m)
}

func TestSample2(t *testing.T) {
	first := []int{3, 4}
	second := []int{3, 2}
	n, m := 5, 5
	runSample(t, first, second, n, m)
}
