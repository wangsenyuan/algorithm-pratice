package main

import "testing"

func runSample(t *testing.T, K int, A string, expect int) {
	res := solve(K, A)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 0
	A := "abcd"
	expect := 6
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := -1
	A := "abac"
	expect := 4
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	A := "aab"
	expect := 4
	runSample(t, k, A, expect)
}
