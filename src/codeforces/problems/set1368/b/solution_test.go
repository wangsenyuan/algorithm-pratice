package main

import "testing"

func runSample(t *testing.T, k int64, expect int) {
	res := solve(k)

	if len(res) != expect {
		t.Errorf("Sample expect %d-length string, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var k int64 = 421654016
	expect := len("ccccccccooooooooddddddddeeeeeeefffffffooooooorrrrrrrccccccceeeeeeesssssss")
	runSample(t, k, expect)
}

func TestSample2(t *testing.T) {
	var k int64 = 3
	expect := len("codeforcesss")
	runSample(t, k, expect)
}
