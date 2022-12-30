package main

import "testing"

func runSample(t *testing.T, n int, k int64, expect int64) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var k int64 = 2
	var expect int64 = 2
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	var k int64 = 4
	var expect int64 = 5
	runSample(t, n, k, expect)
}


func TestSample3(t *testing.T) {
	n := 3
	var k int64 = 7
	var expect int64 = 12
	runSample(t, n, k, expect)
}
