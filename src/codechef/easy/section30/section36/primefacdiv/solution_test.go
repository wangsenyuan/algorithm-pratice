package main

import "testing"

func runSample(t *testing.T, A, B int64, expect bool) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var A int64 = 21
	var B int64 = 63
	expect := true
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	var A int64 = 25
	var B int64 = 20
	expect := false
	runSample(t, A, B, expect)
}
