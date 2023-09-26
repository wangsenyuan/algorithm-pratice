package main

import "testing"

func runSample(t *testing.T, a []int64, b []int64, h int, s int, expect int64) {
	res := solve(a, b, h, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int64{5, 10}
	b := []int64{4, 9}
	h, s := 16, 1
	var expect int64 = 20
	runSample(t, a, b, h, s, expect)
}

func TestSample2(t *testing.T) {
	a := []int64{10, 1}
	b := []int64{5000, 100000}
	h, s := 25, 9
	var expect int64 = 25
	runSample(t, a, b, h, s, expect)
}

func TestSample3(t *testing.T) {
	a := []int64{2, 15}
	b := []int64{2, 19}
	h, s := 23, 1
	var expect int64 = 152
	runSample(t, a, b, h, s, expect)
}

func TestSample4(t *testing.T) {
	a := []int64{3, 19}
	b := []int64{4, 29}
	h, s := 11, 2

	var expect int64 = 67
	runSample(t, a, b, h, s, expect)
}
