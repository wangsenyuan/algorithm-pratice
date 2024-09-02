package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	cnt := 1
	tot := 1
	for _, x := range res {
		cnt += x
		tot += cnt
	}
	if tot != n {
		t.Fatalf("Sample result %v, get wrong final result %d, expecting %d", res, tot, n)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1)
}
