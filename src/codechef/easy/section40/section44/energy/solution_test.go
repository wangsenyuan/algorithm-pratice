package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, L int, B int, G int, boys_att [][]int, girls_att [][]int, boys_like [][]int, girls_like [][]int, expect []int) {
	res := solve(L, B, G, boys_att, girls_att, boys_like, girls_like)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B, G, L := 2, 3, 10
	boys_att := [][]int{
		{0, 10},
		{1, 6},
	}
	boys_like := [][]int{
		{0, 1},
		{0, 2, 1},
	}
	girls_att := [][]int{
		{4, 5},
		{3, 8},
		{2, 8},
	}
	girls_like := [][]int{
		{0, 1},
		{1},
		{0},
	}
	expect := []int{7, 2, 1}
	runSample(t, L, B, G, boys_att, girls_att, boys_like, girls_like, expect)
}
