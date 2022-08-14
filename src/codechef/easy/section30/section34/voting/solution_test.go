package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, D [][]int, expect_dist int, expect_rank []int) {
	res_dist, res_rank := solve(D)

	if expect_dist != res_dist || !reflect.DeepEqual(expect_rank, res_rank) {
		t.Errorf("Sample expect %d %v, but got %d, %v", expect_dist, expect_rank, res_dist, res_rank)
	}
}

func TestSample1(t *testing.T) {
	D := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	expect_dist := 0
	expect_rank := []int{1, 2, 3}
	runSample(t, D, expect_dist, expect_rank)
}

func TestSample2(t *testing.T) {
	D := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	expect_dist := 4
	expect_rank := []int{2, 1, 3}
	runSample(t, D, expect_dist, expect_rank)
}
