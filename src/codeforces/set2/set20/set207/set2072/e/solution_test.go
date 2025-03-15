package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, k int) {
	res := solve(k)
	if len(res) > 500 {
		t.Fatalf("Sample result %v, not correct, having too much points", res)
	}
	slices.SortFunc(res, func(a, b []int) int {
		if a[1] != b[1] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	var sum int
	for i := 0; i < len(res); {
		j := i
		for i < len(res) && res[i][1] == res[j][1] {
			i++
		}
		cnt := i - j
		sum += cnt * (cnt - 1) / 2
	}
	if sum != k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0)
}
func TestSample2(t *testing.T) {
	runSample(t, 2)
}
func TestSample3(t *testing.T) {
	runSample(t, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 1e5)
}
