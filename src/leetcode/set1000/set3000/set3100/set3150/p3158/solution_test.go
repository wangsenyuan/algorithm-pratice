package p3158

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries [][]int, expect []bool) {
	res := getResults(queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := [][]int{
		{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2},
	}
	expect := []bool{false, true, true}
	runSample(t, queries, expect)
}

func TestSample2(t *testing.T) {
	queries := [][]int{
		{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6},
	}
	expect := []bool{true, true, false}
	runSample(t, queries, expect)
}

func TestSample3(t *testing.T) {
	queries := [][]int{
		{2, 1, 1},
	}
	expect := []bool{true}
	runSample(t, queries, expect)
}

func TestSample4(t *testing.T) {
	queries := [][]int{
		{1, 18}, {1, 19}, {2, 26, 15}, {1, 5}, {1, 26}, {2, 2, 23}, {1, 9}, {2, 9, 20}, {2, 21, 11}, {2, 17, 20},
	}
	expect := []bool{true, false, false, false, false}
	runSample(t, queries, expect)
}
