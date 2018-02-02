package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, hs []int, ms []int, expect [][]int) {
	res := solve(n, hs, ms)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v, expect %v, but got %v", n, hs, ms, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	hs := []int{0, 2, 0}
	ms := []int{1, 2, 3}
	expect := [][]int{{0, 1}}
	runSample(t, n, hs, ms, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	hs := []int{2, 1}
	ms := []int{1, 2}
	expect := [][]int{{0, 0}, {2, INF}}
	runSample(t, n, hs, ms, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	hs := []int{1, 2, 3}
	ms := []int{1, 2, 3}
	var expect [][]int
	runSample(t, n, hs, ms, expect)
}
