package main

import "testing"

func runSample(t *testing.T, a []int, expect int64) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 8, 12, 90}
	var expect int64 = 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{313923, 246114, 271842, 371982, 284858, 10674, 532090, 593483, 185123, 364245, 665161, 241644, 604914, 645577, 410849, 387586, 732231, 952593, 249651, 36908}
	var expect int64 = 6
	runSample(t, a, expect)
}
