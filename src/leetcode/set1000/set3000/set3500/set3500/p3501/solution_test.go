package p3501

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := maxActiveSectionsAfterTrade(s, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01"
	queries := [][]int{
		{0, 1},
	}
	expect := []int{1}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "0100"
	queries := [][]int{
		{0, 3}, {0, 2}, {1, 3}, {2, 3},
	}
	expect := []int{4, 3, 1, 1}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "1000100"
	queries := [][]int{
		{1, 5}, {0, 6}, {0, 4},
	}
	expect := []int{6, 7, 2}
	runSample(t, s, queries, expect)
}

func TestSample4(t *testing.T) {
	s := "01010"
	queries := [][]int{
		{0, 3}, {1, 4}, {1, 3},
	}
	expect := []int{4, 4, 2}
	runSample(t, s, queries, expect)
}
