package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, b []string) {
	res := solve(b)
	n := len(b)

	for i := 0; i < n; i++ {
		sort.Ints(res[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := checkSubSet(res[i], res[j])
			if b[i][j] == '1' != tmp {
				t.Fatalf("Sample result %v not correct", res)
			}
		}
	}
}

func checkSubSet(a []int, b []int) bool {
	for i, j := 0, 0; i < len(a); i++ {
		for j < len(b) && b[j] < a[i] {
			j++
		}
		if j == len(b) {
			return false
		}
		j++
	}
	return len(a) < len(b)
}

func TestSample1(t *testing.T) {
	b := []string{
		"0001",
		"1001",
		"0001",
		"0000",
	}
	runSample(t, b)
}
