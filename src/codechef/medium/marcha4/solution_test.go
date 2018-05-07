package main

import "testing"

func runSample(t *testing.T, n int, k int, expectBegin, expectEnd int) {
	begin, end := solve(n, k)
	if begin != expectBegin || end != expectEnd {
		t.Errorf("sample %d %d, expect %d %d, but got %d %d", n, k, expectBegin, expectEnd, begin, end)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	begin, end := 25, 56
	runSample(t, n, k, begin, end)
}

func TestSample2(t *testing.T) {
	n, k := 9, 3
	begin, end := 387, 489
	runSample(t, n, k, begin, end)
}
