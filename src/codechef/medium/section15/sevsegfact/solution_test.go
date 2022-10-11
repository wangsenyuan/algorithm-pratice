package main

import "testing"

func runSample(t *testing.T, num string, cost []int, expect int64) {
	res := solve(num, cost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "10"
	cost := []int{1, 31, 21, 5, 3, 12, 9}
	var expect int64 = 48
	runSample(t, num, cost, expect)
}

func TestSample2(t *testing.T) {
	num := "69697"
	cost := []int{1, 2, 3, 4, 5, 6, 7}
	var expect int64 = 5
	runSample(t, num, cost, expect)
}

func TestSample3(t *testing.T) {
	num := "3964083074239146524786"
	cost := []int{1, 41, 31, 21, 11, 1, 41}
	var expect int64 = 72
	runSample(t, num, cost, expect)
}
