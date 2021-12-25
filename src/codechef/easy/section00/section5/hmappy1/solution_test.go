package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, A []int, S string, expect []int) {
	res := solve(n, k, A, S)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %s, expect %v, but got %v", n, k, A, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 1, 0, 1, 1}
	k := 3
	S := "?!?!?"
	expect := []int{2, 3, 3}
	runSample(t, n, k, A, S, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 1, 0, 1, 1}
	k := 4
	S := "?!?!?"
	expect := []int{2, 3, 4}
	runSample(t, n, k, A, S, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	A := []int{1, 1, 1, 0, 1, 1}
	k := 5
	S := "!?!?"
	expect := []int{4, 5}
	runSample(t, n, k, A, S, expect)
}
