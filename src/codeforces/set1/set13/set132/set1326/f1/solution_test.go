package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S []string, expect []int64) {
	res := solve(n, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []string{
		"011",
		"101",
		"110",
	}
	expect := []int64{0, 0, 0, 6}
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := []string{
		"0101",
		"1000",
		"0001",
		"1010",
	}
	expect := []int64{2, 2, 6, 2, 2, 6, 2, 2}
	runSample(t, n, S, expect)
}
