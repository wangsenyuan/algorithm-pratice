package main

import "testing"

func runSample(t *testing.T, n int, m int, prices [][]int, discounts [][]int, expect int64) {
	res := solve(n, m, prices, discounts)

	if res != expect {
		t.Errorf("sample %d %d %v %v, expect %d, but got %d", n, m, prices, discounts, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	prices := [][]int{{3, 4}, {1, 2}}
	discounts := [][]int{{1, 0}, {0, 1}}
	runSample(t, n, m, prices, discounts, 3)
}

func TestSample2(t *testing.T) {
	n, m := 2, 4
	prices := [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}}
	discounts := [][]int{{2, 3, 2, 1}, {1, 2, 1, 1}}
	runSample(t, n, m, prices, discounts, 2)
}
