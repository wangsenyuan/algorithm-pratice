package main

import "testing"

func runSample(t *testing.T, S string, n int, k int64, expect int) {
	res := solve(S, n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSampele1(t *testing.T) {
	S := "abac"
	n := 3
	var k int64 = 4
	expect := 1
	runSample(t, S, n, k, expect)
}

func TestSampele2(t *testing.T) {
	S := "abac"
	n := 3
	var k int64 = 9
	expect := 2
	runSample(t, S, n, k, expect)
}
