package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, seq []int, expect []int) {
	res := solve(n, seq)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, seq, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	seq := []int{1, 1}
	expect := []int{1, 2}
	runSample(t, n, seq, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	seq := []int{1, 2, 3, 3}
	expect := []int{1, 2, 3, 3}
	runSample(t, n, seq, expect)
}
