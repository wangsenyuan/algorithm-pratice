package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S []string, Q []string, expect []int64) {
	res := solve(S, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []string{
		"abc",
		"abd",
	}
	Q := []string{
		"3 1",
		"3 2",
		"1 1 2 5",
		"3 1",
		"3 2",
	}
	expect := []int64{0, 0, 5, 5}
	runSample(t, S, Q, expect)
}

func TestSample2(t *testing.T) {
	S := []string{
		"abc",
	}
	Q := []string{
		"2 1 2 de",
		"2 2 3 fg",
		"1 1 2 1",
		"1 2 3 5",
		"3 1",
		"3 2",
		"3 3",
	}
	expect := []int64{1, 6, 6}
	runSample(t, S, Q, expect)
}
