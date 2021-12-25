package main

import "testing"

func runSample(t *testing.T, n int, ingredients []int, expect int64) {
	res := solve(n, ingredients)
	if res != expect {
		t.Errorf("sample: %d %v, should give %d, but got %d", n, ingredients, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	ingredients := []int{1, 2}
	expect := int64(8)
	runSample(t, n, ingredients, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	ingredients := []int{1, 2, 3}
	expect := int64(24)
	runSample(t, n, ingredients, expect)
}
