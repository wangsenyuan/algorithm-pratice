package main

import "testing"

func runSample(t *testing.T, n int, es []int64, from int, end int, expect int64) {
	res := solve(n, es, from, end)
	if res != expect {
		t.Errorf("sample %d %v %d %d, should give %d, but got %d", n, es, from, end, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	es := []int64{1, 2, 1, 1}
	from, end := 1, 3
	expect := int64(2)
	runSample(t, n, es, from, end, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	es := []int64{-5, 100, 100, 100, 2}
	from, end := 1, 5
	expect := int64(-8)
	runSample(t, n, es, from, end, expect)
}
