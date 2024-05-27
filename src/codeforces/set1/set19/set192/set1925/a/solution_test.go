package main

import "testing"

func runSample(t *testing.T, n, k int, expect int) {
	res := solve(n, k)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 1, 2
	expect := len("ab")
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 1
	expect := len("aa")
	runSample(t, n, k, expect)
}
func TestSample3(t *testing.T) {
	n, k := 2, 2
	expect := len("baab")
	runSample(t, n, k, expect)
}