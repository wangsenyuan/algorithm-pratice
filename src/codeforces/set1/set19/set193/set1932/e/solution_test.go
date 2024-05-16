package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "42"
	expect := "46"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "12345"
	expect := "13715"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "100"
	// 100 => 99 = 3
	// 90 => 89 => 2
	// ...
	// 10 => 09 => 2
	// 2 + 9 + 100
	expect := "111"
	runSample(t, s, expect)
}
