package main

import "testing"

func runSample(t *testing.T, k int, prices []int, expect int) {
	res := solve(k, prices)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	prices := []int{3, 2, 7, 1, 4, 8}
	expect := 8
	runSample(t, k, prices, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	prices := []int{2, 1, 5, 6}
	expect := 4
	runSample(t, k, prices, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	prices := []int{1, 1}
	expect := 2
	runSample(t, k, prices, expect)
}

func TestSample4(t *testing.T) {
	k := 10
	prices := []int{2, 0, 0, 5, 9, 4}
	expect := 1
	runSample(t, k, prices, expect)
}

func TestSample5(t *testing.T) {
	k := 5
	prices := []int{5, 3, 8, 6, 3, 2}
	expect := 5
	runSample(t, k, prices, expect)
}
