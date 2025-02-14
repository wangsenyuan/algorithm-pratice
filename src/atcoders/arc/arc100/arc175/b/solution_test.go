package main

import "testing"

func runSample(t *testing.T, s string, A int, B int, expect int) {
	res := solve(s, A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B := 3, 2
	s := ")))(()"
	expect := 5
	runSample(t, s, A, B, expect)
}

func TestSample2(t *testing.T) {
	A, B := 2622, 26092458
	s := "))()((((()()(("
	expect := 52187538
	runSample(t, s, A, B, expect)
}
