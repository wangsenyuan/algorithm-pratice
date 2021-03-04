package main

import "testing"

func runSample(t *testing.T, n int, A []int64, k int64, expect bool) {
	res := solve(n, A, k)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var k int64 = 0
	A := []int64{2, 3, 7}
	expect := true
	runSample(t, n, A, k, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	var k int64 = -1
	A := []int64{31415926, 27182818}
	expect := false
	runSample(t, n, A, k, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	var k int64 = 1000000000000000000
	A := []int64{1, 1000000000000000000}
	expect := true
	runSample(t, n, A, k, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	var k int64 = -1000000000000000000
	A := []int64{-1000000000000000000, 123}
	expect := true
	runSample(t, n, A, k, expect)
}

func TestSample5(t *testing.T) {
	n := 6
	var k int64 = 80
	A := []int64{-5, -20, 13, -14, -2, -11}
	expect := false
	runSample(t, n, A, k, expect)
}
