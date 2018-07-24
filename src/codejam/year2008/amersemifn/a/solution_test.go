package main

import "testing"

func runSample(t *testing.T, N int, mixtures [][]string, expect int) {
	res := solve(N, mixtures)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, mixtures, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 3
	mixtures := [][]string{
		{"SOUP", "3", "STOCK", "salt", "water"},
		{"STOCK", "2", "chicken", "VEGETABLES"},
		{"VEGETABLES", "2", "celery", "onions"},
	}
	expect := 2
	runSample(t, N, mixtures, expect)
}

func TestSample2(t *testing.T) {
	N := 5
	mixtures := [][]string{
		{"MILKSHAKE", "4", "milk", "icecream", "FLAVOR", "FRUIT"},
		{"FRUIT", "2", "banana", "berries"},
		{"FLAVOR", "2", "SPICES", "CHOCOLATE"},
		{"SPICES", "2", "nutmeg", "cinnamon"},
		{"CHOCOLATE", "2", "cocoa", "syrup"},
	}
	expect := 3
	runSample(t, N, mixtures, expect)
}
