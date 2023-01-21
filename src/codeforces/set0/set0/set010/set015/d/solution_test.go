package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, a int, b int, H [][]int, expect_pos [][]int, expect_ans []int64) {
	pos, ans := solve(n, m, a, b, H)

	if !reflect.DeepEqual(pos, expect_pos) || !reflect.DeepEqual(ans, expect_ans) {
		t.Errorf("Sample expect %v, %v, but got %v, %v", expect_pos, expect_ans, pos, ans)
	}
}

func TestSample1(t *testing.T) {
	n, m, a, b := 2, 2, 1, 2
	H := [][]int{
		{1, 2},
		{3, 5},
	}
	pos := [][]int{
		{1, 1},
		{2, 1},
	}
	ans := []int64{1, 2}

	runSample(t, n, m, a, b, H, pos, ans)
}

func TestSample2(t *testing.T) {
	n, m, a, b := 4, 4, 2, 2
	H := [][]int{
		{1, 5, 3, 4},
		{2, 7, 6, 1},
		{1, 1, 2, 2},
		{2, 2, 1, 2},
	}
	pos := [][]int{
		{3, 1},
		{3, 3},
		{1, 2},
	}
	ans := []int64{2, 3, 9}

	runSample(t, n, m, a, b, H, pos, ans)
}
