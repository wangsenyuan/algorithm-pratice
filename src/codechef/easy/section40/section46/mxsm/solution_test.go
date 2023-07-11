package main

import "testing"

func runSample(t *testing.T, L [][]int, expect_sum int64, expect_cnt int) {
	res_sum, res_cnt := solve(L)

	if expect_sum != res_sum || expect_cnt != res_cnt {
		t.Fatalf("Sample expect %d, %d, but got %d %d", expect_sum, expect_cnt, res_sum, res_cnt)
	}
}

func TestSample1(t *testing.T) {
	L := [][]int{
		{3, 4},
		{4, 5},
	}
	var sum int64 = 63
	var cnt = 1
	runSample(t, L, sum, cnt)
}

func TestSample2(t *testing.T) {
	L := [][]int{
		{4, -5, 4},
		{-2, 3, -2},
	}
	var sum int64 = 12
	var cnt = 2
	runSample(t, L, sum, cnt)
}
