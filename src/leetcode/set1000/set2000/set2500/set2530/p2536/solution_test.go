package p2536

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, qs [][]int, expect [][]int) {
	res := rangeAddQueries(n, qs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	qs := [][]int{
		{1, 1, 2, 2},
		{0, 0, 1, 1},
	}
	expect := [][]int{
		{1, 1, 0},
		{1, 2, 1},
		{0, 1, 1},
	}
	runSample(t, n, qs, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	qs := [][]int{
		{0, 0, 1, 1},
	}
	expect := [][]int{
		{1, 1},
		{1, 1},
	}
	runSample(t, n, qs, expect)
}
