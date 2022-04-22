package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, x int) {
	res := solve(n, x)

	if !check(res, x) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(arr []int, avg int) bool {
	var sum int

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i > 0 && arr[i] == arr[i-1] {
			return false
		}
	}
	return sum == len(arr)*avg
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000, 100)
}
