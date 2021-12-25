package main

import "testing"

func TestSample(t *testing.T) {
	n, m := 6, 3
	ss := []int{0, 1, 3, 4, 7, 10}
	a, b := solve(n, ss, m)
	if a != 3 {
		t.Errorf("the sample should have stops equals to 3, but get %d", a)
	}
	if b != 2 {
		t.Errorf("the sample should have 2 ways to have the minmum stops, but get %d", b)
	}
}
