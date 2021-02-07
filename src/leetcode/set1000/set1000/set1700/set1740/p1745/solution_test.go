package p1745

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, candiesCount []int, queries [][]int, expect []bool) {
	res := canEat(candiesCount, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	candies := []int{7, 4, 5, 3, 8}
	queries := [][]int{
		{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000},
	}
	expect := []bool{true, false, true}
	runSample(t, candies, queries, expect)
}

func TestSample2(t *testing.T) {
	candies := []int{5, 2, 6, 4, 1}
	queries := [][]int{
		{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1},
	}
	expect := []bool{false, true, true, false, false}
	runSample(t, candies, queries, expect)
}
