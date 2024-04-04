package main

import "testing"

func runSample(t *testing.T, x []int, k int, expect []int) {
	res := solve(x, k)

	a := check(x, expect)
	b := check(x, res)

	if a != b {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func check(x []int, res []int) int {
	var dist int
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			dist += abs(x[res[i]-1] - x[res[j]-1])
		}
	}
	return dist
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	x := []int{1, 100, 101}
	k := 2
	expect := []int{2, 3}
	runSample(t, x, k, expect)
}
