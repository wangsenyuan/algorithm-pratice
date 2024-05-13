package main

import "testing"

func runSample(t *testing.T, n int, l int, r int, expect int) {
	res := solve(n, l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 9, 99, 711426616)
}
