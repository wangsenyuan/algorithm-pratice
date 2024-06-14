package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	n := len(a)
	ok, houses, move_to := solve(n, a)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v %v", expect, ok, houses, move_to)
	}

	if !expect {
		return
	}
	for i := 0; i < n; i++ {
		j := move_to[i] - 1
		d := distance(houses[i], houses[j])
		if d != a[i] {
			t.Fatalf("Sample result %v, not correct, as distance from %d -> %d is %d, not %d", houses, i, j, d, a[i])
		}
	}
}
func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	a := []int{0, 4, 2, 4}
	runSample(t, a, true)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 3, 1}
	runSample(t, a, true)
}

func TestSample3(t *testing.T) {
	a := []int{0, 0}
	runSample(t, a, true)
}
