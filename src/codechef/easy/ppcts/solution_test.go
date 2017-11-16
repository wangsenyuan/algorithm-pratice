package main

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	n, m := 2, 3
	dcs := [][]int{
	 {1, 2}, {0, 1},
	}

	res := solve(n, m, dcs, "RDL")

	expect := []int{4, 6, 6}

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("sample should give %v, but got %v", expect, res)
	}
}
