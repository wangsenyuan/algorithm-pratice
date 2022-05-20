package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S string, E [][]int, Q [][]int, expect []bool) {
	res := solve(n, S, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}

}

func TestSample1(t *testing.T) {
	n := 7
	S := "ydbxcbd"
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{2, 5},
		{6, 4},
		{7, 6},
	}
	Q := [][]int{
		{3, 5},
		{3, 6},
	}
	expect := []bool{false, true}

	runSample(t, n, S, E, Q, expect)
}
